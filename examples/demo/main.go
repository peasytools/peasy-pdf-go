// Demo script for peasy-pdf-go — PeasyPDF API client.
// Demonstrates listing tools, fetching a specific tool, and searching.
package main

import (
	"context"
	"fmt"
	"log"

	peasy "github.com/peasytools/peasy-pdf-go"
)

func main() {
	ctx := context.Background()
	client := peasy.New()

	// List available PDF tools
	fmt.Println("=== PDF Tools ===")
	tools, err := client.ListTools(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, tool := range tools.Results {
		fmt.Printf("  %s: %s\n", tool.Name, tool.Description)
	}
	fmt.Printf("  Total: %d tools\n", tools.Count)

	// Get a specific tool
	fmt.Println("\n=== Merge PDF Tool ===")
	tool, err := client.GetTool(ctx, "merge-pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  Name: %s\n", tool.Name)
	fmt.Printf("  Description: %s\n", tool.Description)
	fmt.Printf("  Category: %s\n", tool.Category)

	// List formats
	fmt.Println("\n=== Formats ===")
	formats, err := client.ListFormats(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range formats.Results {
		fmt.Printf("  %s (%s): %s\n", f.Name, f.Extension, f.MimeType)
	}

	// Search across content
	fmt.Println("\n=== Search: 'compress' ===")
	results, err := client.Search(ctx, "compress")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  Found %d tools, %d formats, %d glossary terms\n",
		len(results.Results.Tools),
		len(results.Results.Formats),
		len(results.Results.Glossary))

	// List glossary terms
	fmt.Println("\n=== Glossary ===")
	glossary, err := client.ListGlossary(ctx, peasy.ListOptions{Limit: 5})
	if err != nil {
		log.Fatal(err)
	}
	for _, term := range glossary.Results {
		fmt.Printf("  %s: %s\n", term.Term, truncate(term.Definition, 80))
	}
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
