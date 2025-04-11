package main

import (
	"bufio"
	"fmt"
	urlparser "net/url"
	"os"
	"strings"
)

func showHelp() {
	fmt.Println("\nUsage: cat <urls_file_path> | pathslapper <path-to-append>")
	fmt.Println("Example: ")
	fmt.Println("      cat urls.txt | pathslapper /api/v1/admin\n")
}

func removeURLQueryParams(url string) string {
	parsedURL, _ := urlparser.Parse(url) // ignore error
	parsedURL.RawQuery = ""
	parsedURL.Fragment = ""
	return parsedURL.String()
}

func main() {
	if len(os.Args) != 2 {
		showHelp()
		return
	}
	path := os.Args[1]

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parsedURL := removeURLQueryParams(scanner.Text())
		newURL := strings.TrimRight(parsedURL, "/") + path
		fmt.Println(newURL)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
