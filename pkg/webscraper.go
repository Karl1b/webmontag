package webmontag

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/karl1b/webmontag/pkg/internal/database"
)

type subPage struct {
	FullUrl        string `json:"fullUrl"`
	HasBeenScanned bool   `json:"hasBeenScanned"`
}

type externalLink struct {
	Url        string `json:"url"`
	Statuscode int    `json:"statusCode"`
}

type pageDetails struct {
	PageDomain      string         `json:"pageDomain"`
	ScannedSubpages []subPage      `json:"scannedSubpages"`
	ExternalLinks   []externalLink `json:"externalLinks"`
	SearchDeep      int            `json:"searchDeep"`
	MaxSearchDeep   int            `json:"maxSearchDeep"`
}

func StartScraper(rawURL string, queries *database.Queries) {

	parsed, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println(err)
		return // if the urlstring is not valid it will return here.
	}

	page := pageDetails{
		PageDomain:    parsed.Host,
		SearchDeep:    0,
		MaxSearchDeep: Options.maxSearchDeep,
	}

	// Options for the chromedp instance
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath("/opt/google/chrome/chrome"),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
	)

	scanOnePage(rawURL, &page, &opts) // Initial call of the recursive function

	checkExternalLinks(&page, &opts)

	// Marshall to json. This is for presentation only

	page.ScannedSubpages = []subPage{} // Set this unintersing to empty

	pageJSON, err := json.MarshalIndent(page, "", "  ")
	if err != nil {
		fmt.Println("Error converting page details to JSON:", err)
		return
	}

	// Construct filename:  URL hostname + timestamp
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL for filename:", err)
		return
	}
	filename := parsedURL.Hostname() + "_" + time.Now().Format("20060102_150405")
	filename = strings.ReplaceAll(filename, ".", "_") + ".json"
	// Write JSON data to file
	err = os.WriteFile(filename, pageJSON, 0644) // 0644 sets typical file permissions
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("JSON data saved to:", filename)

	// Print the JSON string.
	fmt.Println(string(pageJSON))
	// Saves the results to the db. Final Step only if this runs with db flag.
	if queries != nil {
		dbpageID, err := queries.UpsertPage(context.Background(), database.UpsertPageParams{
			Domain:        sql.NullString{String: rawURL, Valid: true},
			Searchdeep:    sql.NullInt32{Int32: int32(page.SearchDeep), Valid: true},
			Maxsearchdeep: sql.NullInt32{Int32: int32(page.MaxSearchDeep), Valid: true},
		})
		if err != nil {
			fmt.Print("UpsertPage Error")
			fmt.Println(err)
		}
		for _, externalLink := range page.ExternalLinks {

			externalLinkID, err := queries.InsertOrGetLink(context.Background(), database.InsertOrGetLinkParams{
				Domain:     sql.NullString{String: externalLink.Url, Valid: true},
				Statuscode: sql.NullInt32{Int32: int32(externalLink.Statuscode), Valid: true},
			})
			if err != nil {
				fmt.Print("InsertLink Error")
				fmt.Println(err)
				continue
			}

			err = queries.InsertPageLinkAssociation(context.Background(), database.InsertPageLinkAssociationParams{
				PagesID: sql.NullInt32{Int32: dbpageID, Valid: true},
				LinksID: sql.NullInt32{Int32: externalLinkID, Valid: true}})
			if err != nil {
				fmt.Print("InsertPageLinkAssociation Error")
				fmt.Println(err)
				continue
			}
		}
	}
}

// this is the recursive scrape function
func scanOnePage(currentPage string, page *pageDetails, opts *[]func(*chromedp.ExecAllocator)) {

	isImageURL := strings.HasSuffix(currentPage, ".gif") || strings.HasSuffix(currentPage, ".png") || strings.HasSuffix(currentPage, ".jpg")
	if isImageURL {
		return
	}

	if page.SearchDeep >= page.MaxSearchDeep {
		return // max search deep reached.
	}
	page.SearchDeep = page.SearchDeep + 1

	searchDeep := fmt.Sprint(page.SearchDeep)

	fmt.Println("I currently scrape: " + currentPage + "at search deep: " + searchDeep)

	ectx, ecancel := chromedp.NewExecAllocator(context.Background(), *opts...)
	defer ecancel()
	ctxone, cancel := chromedp.NewContext(ectx) // Create a new CDP context
	defer cancel()
	ctxWithTimeout, tcancel := context.WithTimeout(ctxone, time.Duration(Options.breaktimeS)*time.Second)
	defer tcancel()

	if err := chromedp.Run(ctxWithTimeout,
		chromedp.Navigate(currentPage),
		//chromedp.Sleep(0.25*time.Second), or even sleep for a random value to trick webservers?
	); err != nil {
		fmt.Println("Failed to navigate to domain:", currentPage, err)
		return
	}

	getLinks(ctxWithTimeout, page)

	for index := range page.ScannedSubpages {
		if !page.ScannedSubpages[index].HasBeenScanned {
			page.ScannedSubpages[index].HasBeenScanned = true
			scanOnePage(page.ScannedSubpages[index].FullUrl, page, opts)
		}
	}
}

func getLinks(ctx context.Context, page *pageDetails) error {

	var links []string

	err := chromedp.Run(ctx,
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a[href]')).map(a => a.href)`, &links),
	)
	if err != nil {
		fmt.Println("Failed to get links:", err)
		return err
	}

	for _, link := range links {

		if !isValidURL(link) {
			continue // if the link is not valid do nothing
		}
		if strings.Contains(link, page.PageDomain) {
			if !containsSubPage(page.ScannedSubpages, link) {
				page.ScannedSubpages = append(page.ScannedSubpages, subPage{
					FullUrl: link, HasBeenScanned: false,
				})

			}

		} else {
			if !containsExternalLink(page.ExternalLinks, link) {
				page.ExternalLinks = append(page.ExternalLinks, externalLink{Url: link, Statuscode: 0})
			}
		}
	}
	return nil
}

func containsSubPage(slice []subPage, value string) bool {
	for _, v := range slice {
		if v.FullUrl == value {
			return true
		}
	}
	return false
}

func containsExternalLink(slice []externalLink, value string) bool {
	for _, v := range slice {
		if v.Url == value {
			return true
		}
	}
	return false
}

func isValidURL(link string) bool {

	parsedURL, err := url.Parse(link)
	if err != nil {
		return false
	}

	return parsedURL.Scheme != "" && parsedURL.Host != ""
}

func ConcurrentScraper(urls []string, queries *database.Queries) {

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, Options.maxGoRoutines) // Create a semaphore with a capacity of 5

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			semaphore <- struct{}{}
			if queries != nil {
				StartScraper(url, queries)
			} else {
				StartScraper(url, nil)
			}
			<-semaphore
		}(url)
	}
	wg.Wait()
}

func checkExternalLinks(page *pageDetails, opts *[]func(*chromedp.ExecAllocator)) {

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, Options.maxGoRoutines) // Create a semaphore with a capacity of 5

	for i, _ := range page.ExternalLinks {

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}

			checkExternalLink(&page.ExternalLinks[i], page, opts)

			<-semaphore
		}(i)

	}
	wg.Wait()

}

func checkExternalLink(extLink *externalLink, page *pageDetails, opts *[]func(*chromedp.ExecAllocator)) {
	url := extLink.Url

	ectx, ecancel := chromedp.NewExecAllocator(context.Background(), *opts...)
	defer ecancel()
	ctxone, cancel := chromedp.NewContext(ectx) // Create a new CDP context
	defer cancel()
	ctxWithTimeout, tcancel := context.WithTimeout(ctxone, time.Duration(Options.breaktimeS)*time.Second)
	defer tcancel()

	if err := chromedp.Run(ctxWithTimeout,
		chromedp.Navigate(url),
		//chromedp.Sleep(0.25*time.Second), or even sleep for a random value to trick webservers?
	); err != nil {
		fmt.Println("Failed to navigate to domain:", url, err)
		return
	}

	statuscode, err := getHTTPStatusCode(url)
	if err != nil {
		statuscode = 999
	}

	extLink.Statuscode = statuscode

}

func getHTTPStatusCode(url string) (int, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
