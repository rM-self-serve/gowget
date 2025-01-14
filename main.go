package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	outputFile = flag.String("O", "", "Write documents to FILE instead of default")
	version    = flag.Bool("version", false, "Display version information")
	help       = flag.Bool("help", false, "Display this help message")
)

func printUsage() {
	fmt.Printf("Usage: %s [OPTIONS] URL\n", os.Args[0])
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Println("\nExample:")
	fmt.Printf("  %s -O output.txt https://example.com/file.txt\n", os.Args[0])
}

func downloadFile(url, outputPath string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("User-Agent", "gowget")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error downloading file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned status code %d", resp.StatusCode)
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing to output file: %v", err)
	}

	return nil
}

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("gowget version\n")
		return
	}

	if *help || flag.NArg() == 0 {
		printUsage()
		return
	}

	url := flag.Arg(0)
	output := *outputFile

	if output == "" {
		output = filepath.Base(url)
		if output == "" || output == "." {
			output = "index.html"
		}
	}

	fmt.Printf("Downloading to %s...\n", output)
	err := downloadFile(url, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Download completed successfully!")
}
