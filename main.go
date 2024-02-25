package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	webmontag "github.com/karl1b/webmontag/pkg"
)

var rootCmd = &cobra.Command{
	Use:   "webmontag",
	Short: "webmontag",
	Long:  `webmontag made by Karl Breuer.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("################################")
		fmt.Println("Welcome to Webmontag Webscraper.")
		fmt.Println("################################")
		fmt.Println("Available commands:")
		fmt.Println("scrape 'url'")
		fmt.Println("ultrascrape 'url1' 'url2'")
		fmt.Println("startserver")
		fmt.Println("startserver db")
	},
}

var startWebscraper = &cobra.Command{
	Use:   "scrape",
	Short: "Starts the Scraper with one page.",
	Long:  `Starts the Scraper with one page.`,
	Run: func(cmd *cobra.Command, args []string) {
		webmontag.StartScraper(args[0], nil)
	},
}

var ultraScraper = &cobra.Command{
	Use:   "ultrascrape",
	Short: "Starts the Scraper with an Array of pages.",
	Long:  `Starts the Scraper with an Array of pages.`,
	Run: func(cmd *cobra.Command, args []string) {
		webmontag.ConcurrentScraper(args, nil)
	},
}

var startserver = &cobra.Command{
	Use:   "startserver",
	Short: "Starts the Webserver use db flag to enable database",
	Long:  `Starts the Webserver use db flag to enable database`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			webmontag.Startwebserver(args[0])

		} else {
			webmontag.Startwebserver("")
		}

	},
}

func init() {
	rootCmd.AddCommand(startWebscraper)
	rootCmd.AddCommand(startserver)
	rootCmd.AddCommand(ultraScraper)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
