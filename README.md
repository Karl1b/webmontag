# Webmontag Golang Demo

This repository contains a high-performance web scraper and web server written in Go, designed to offer a basic introduction to the Go language. Its primary functionality includes scraping all external links from specified web pages and supporting concurrent scraping of multiple pages.

**Important Note**: While this tool demonstrates powerful web scraping capabilities, it is crucial to use it responsibly. Web scraping may be subject to legal restrictions in your jurisdiction, and it's important to respect the terms of service of any website you interact with. This software is intended for educational purposes only and should not be deployed in production environments.

## Requirements

### Mandatory
- **Go Environment**: Ensure you have Go installed on your system. You can download it from [the official Go website](https://golang.org/dl/).

### Dependencies
1. **Chromedp**: A faster, simpler way to drive browsers supporting the Chrome DevTools Protocol.
   - Documentation: [https://pkg.go.dev/github.com/chromedp/chromedp](https://pkg.go.dev/github.com/chromedp/chromedp)
   - GitHub Repository: [https://github.com/chromedp/chromedp](https://github.com/chromedp/chromedp)

2. **PostgreSQL (Optional)**: If you wish to use a database for storing scrape results, install PostgreSQL and set up a database, user, and necessary permissions.

## Installation

To set up the web scraper/web server:

1. Clone the repository and navigate to the project directory.
2. Copy the example environment file and modify it according to your needs, especially the secret key:
```bash
cp working.env .env
go build
sudo vim ./pkg/up.sh 
./pkg/up.sh #migrate db
```

## Useage

1. `./webmontag` This command displays available commands and their usage.
2. `./webmontag scrape "https://example.com"` Scrapes a single url, logs json
3. `./webmontag ultrascrape ["https://example.com","https://example2.com"]` scrapes many urls concurrently.
4. `./webmontag startserver db'` Starts a webserver with a db connection.

### Server Interaction

With the server up and running, it interacts with HTTP requests as follows:

- **POST for Concurrent Page Scraping**:
    To scrape multiple pages concurrently for external links, send a POST request to `http://127.0.0.1:3333/dbscrapepages` with an Authorization header containing your key (`key yourkey`) and a JSON payload listing the URLs:
    ```json
    {"urls":["https://example.com","https://example.de"]}
    ```

- **GET Scraping Results**:
    Access all stored scraping results in the database by sending a GET request to `http://127.0.0.1:3333/dbseepages` with an Authorization header (`key yourkey`). The server will return a nicely formatted JSON response with the data.

## Disclaimer

Please note that this repository serves as a part of a presentation and is primarily intended for educational and demonstrative purposes. It is not recommended for production environments.


