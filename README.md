# Simple Concurrent Web Scraper

A command-line tool that scrapes all links from a webpage and saves them to a file using Go's concurrency features.

## Features

- Fetch web pages using Go's net/http package
- Parse HTML with the goquery library
- Implement concurrency with goroutines and channels
- Save extracted links to a file

## Installation

1. Clone the repository or download the source code:
   ```bash
   git clone https://github.com/Minn21/Simple-Concurrent-Web-Scraper.git
   cd Simple-Concurrent-Web-Scraper
   ```

2. Initialize the Go module (if not already done):
   ```bash
   go mod init Simple-Concurrent-Web-Scraper
   ```

3. Install the required dependencies:
   ```bash
   go get github.com/PuerkitoBio/goquery
   ```

## Usage

Run the scraper with a URL as an argument:

```bash
go run main.go https://example.com
```

The program will:
1. Fetch the HTML content from the specified URL
2. Parse the HTML to extract all links (`<a href="...">` tags)
3. Save all links to a file named `links.txt` in the current directory
4. Display a confirmation message when complete

## Code Structure

The main components of the scraper are:

- **HTTP Request**: Uses Go's `net/http` package to fetch web pages
- **HTML Parsing**: Uses the `goquery` library to parse HTML and extract links
- **Concurrency**: Implements goroutines and channels for concurrent processing
- **File I/O**: Writes extracted links to a text file

## How Concurrency Works in This Project

1. The main goroutine creates a channel for links and a WaitGroup to synchronize operations
2. A separate goroutine fetches and parses the HTML, sending links to the channel
3. When the parsing is complete, the channel is closed
4. The main goroutine reads from the channel and writes links to the file

## Example Output

After running the scraper on a website, the `links.txt` file will contain a list of all links found on the page, one per line:

```
https://example.com/about
https://example.com/contact
/products
/services
https://twitter.com/example
...
```
