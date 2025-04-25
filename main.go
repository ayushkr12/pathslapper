package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// RemoveURLQueryParams removes the query parameters and fragment from a URL.
func RemoveURLQueryParams(urlStr string) string {
	parsedURL, _ := url.Parse(urlStr) // ignore error
	parsedURL.RawQuery = ""
	parsedURL.Fragment = ""
	return parsedURL.String()
}

// AppendPath appends a path to a URL, ensuring no double slashes.
func AppendPath(urlStr, path string) string {
	// Ensure the path starts with a slash
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	// Remove query parameters and fragments from the URL
	parsedURL := RemoveURLQueryParams(urlStr)

	// Append the path to the URL, ensuring no double slashes
	return strings.TrimRight(parsedURL, "/") + path
}

func showHelp() {
	fmt.Print("\nUsage: cat <url> [<path>] | pathslapper <path> [<url>]\n\n")
	fmt.Println("Example (append urls to a path): ")
	fmt.Println("      cat urls.txt | pathslapper /api/v1/admin")
	fmt.Println("")
	fmt.Println("Example (append paths to a url): ")
	fmt.Println("      cat paths.txt | pathslapper http://example.com")
}

// handleURLFile processes each URL from the input and appends the specified path
func handleURLFile(scanner *bufio.Scanner, path string) {
	for scanner.Scan() {
		parsedURL := RemoveURLQueryParams(scanner.Text()) // Using RemoveURLQueryParams function
		newURL := AppendPath(parsedURL, path)             // Using AppendPath function
		fmt.Println(newURL)
	}
}

// handlePathFile processes each path and appends it to the given URL
func handlePathFile(scanner *bufio.Scanner, url string) {
	for scanner.Scan() {
		path := scanner.Text()
		newURL := AppendPath(url, path) // Using AppendPath function
		fmt.Println(newURL)
	}
}

func main() {
	if len(os.Args) != 2 {
		showHelp()
		return
	}

	// Read the first argument (path or URL)
	input := os.Args[1]

	// Create the scanner object outside the conditional blocks
	scanner := bufio.NewScanner(os.Stdin)

	// Check if input is a URL or path
	if strings.HasPrefix(input, "http") {
		// Case 1: A URL is provided, use pathslapper with URL
		url := input
		handlePathFile(scanner, url)
	} else {
		// Case 2: Path is provided, use it to append paths to URLs
		if !strings.HasPrefix(input, "/") {
			input = "/" + input
		}
		handleURLFile(scanner, input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
