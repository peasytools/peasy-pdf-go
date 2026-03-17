# peasy-pdf-go

[![Go Reference](https://pkg.go.dev/badge/github.com/peasytools/peasy-pdf-go.svg)](https://pkg.go.dev/github.com/peasytools/peasy-pdf-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/peasytools/peasy-pdf-go)](https://goreportcard.com/report/github.com/peasytools/peasy-pdf-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [PeasyPDF](https://peasypdf.com) API — merge, split, rotate, and compress PDF files. Built with `net/http`, `encoding/json`, and zero external dependencies.

Built from [PeasyPDF](https://peasypdf.com), a comprehensive PDF toolkit offering free online tools for merging, splitting, rotating, compressing, and converting PDF documents. The site includes detailed guides on PDF optimization, accessibility best practices, and format conversion, plus a glossary covering terms from linearization to OCR to PDF/A archival compliance.

> **Try the interactive tools at [peasypdf.com](https://peasypdf.com)** — [Merge PDF](https://peasypdf.com/pdf/merge-pdf/), [Split PDF](https://peasypdf.com/pdf/split-pdf/), [Compress PDF](https://peasypdf.com/pdf/compress-pdf/), [Rotate PDF](https://peasypdf.com/pdf/rotate-pdf/), and more.

<p align="center">
  <img src="demo.gif" alt="peasy-pdf-go demo — PDF merge, split, and compress tools in Go terminal" width="800">
</p>

## Table of Contents

- [Install](#install)
- [Quick Start](#quick-start)
- [What You Can Do](#what-you-can-do)
  - [PDF Document Operations](#pdf-document-operations)
  - [Browse Reference Content](#browse-reference-content)
  - [Search and Discovery](#search-and-discovery)
- [API Client](#api-client)
  - [Available Methods](#available-methods)
- [Learn More About PDF Tools](#learn-more-about-pdf-tools)
- [Also Available](#also-available)
- [Peasy Developer Tools](#peasy-developer-tools)
- [License](#license)

## Install

```bash
go get github.com/peasytools/peasy-pdf-go
```

Requires Go 1.21+.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	peasypdf "github.com/peasytools/peasy-pdf-go"
)

func main() {
	client := peasypdf.New()
	ctx := context.Background()

	// List available PDF tools
	tools, err := client.ListTools(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range tools.Results {
		fmt.Printf("%s: %s\n", t.Name, t.Description)
	}
}
```

## What You Can Do

### PDF Document Operations

The Portable Document Format (PDF) was created by Adobe in 1993 and became an open ISO standard (ISO 32000) in 2008. Today PDF is the most widely used format for document exchange, supporting text, images, forms, digital signatures, and embedded multimedia. PeasyPDF provides tools for every common PDF workflow — from merging invoices into a single file to compressing scanned documents for email delivery.

| Operation | Slug | Description |
|-----------|------|-------------|
| Merge PDF | `pdf-merge` | Combine multiple PDF documents into one file |
| Split PDF | `pdf-split` | Extract specific pages or split into individual files |
| Compress PDF | `pdf-compress` | Reduce file size by optimizing images and removing metadata |
| Rotate PDF | `pdf-rotate` | Rotate pages by 90, 180, or 270 degrees |
| PDF to PNG | `pdf-to-png` | Convert PDF pages to high-resolution PNG images |

```go
// Retrieve the PDF merge tool and inspect its capabilities
client := peasypdf.New()
ctx := context.Background()

// Get detailed information about the merge tool
tool, err := client.GetTool(ctx, "pdf-merge")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Tool: %s\n", tool.Name)         // PDF merge tool name
fmt.Printf("Description: %s\n", tool.Description) // How merging works

// List all available PDF tools with pagination
tools, err := client.ListTools(ctx, &peasypdf.ListOptions{Page: 1, Limit: 20})
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Total PDF tools available: %d\n", tools.Count)
```

Learn more: [Merge PDF Tool](https://peasypdf.com/pdf/merge-pdf/) · [How to Merge PDF Files](https://peasypdf.com/guides/how-to-merge-pdf-files/) · [PDF Compression Guide](https://peasypdf.com/guides/pdf-compression-guide/)

### Browse Reference Content

PeasyPDF includes a comprehensive glossary of document format terminology and in-depth guides for common workflows. The glossary covers foundational concepts like PDF linearization (web-optimized PDFs that load page-by-page), OCR (optical character recognition for scanned documents), DPI (dots per inch for print-quality output), and PDF/A (the ISO 19005 archival standard used by governments and libraries worldwide).

| Term | Description |
|------|-------------|
| [PDF](https://peasypdf.com/glossary/pdf/) | Portable Document Format — ISO 32000 open standard |
| [PDF/A](https://peasypdf.com/glossary/pdfa/) | Archival PDF subset (ISO 19005) for long-term preservation |
| [DPI](https://peasypdf.com/glossary/dpi/) | Dots per inch — resolution metric for print and rasterization |
| [OCR](https://peasypdf.com/glossary/ocr/) | Optical character recognition for searchable scanned PDFs |
| [Rasterization](https://peasypdf.com/glossary/rasterization/) | Converting vector PDF content to pixel-based images |

```go
// Browse the PDF glossary for document format terminology
glossary, err := client.ListGlossary(ctx, &peasypdf.ListOptions{
	Search: str("linearization"), // Search for web-optimized PDF concepts
})
if err != nil {
	log.Fatal(err)
}
for _, term := range glossary.Results {
	fmt.Printf("%s: %s\n", term.Term, term.Definition)
}

// Read a specific guide on PDF accessibility best practices
guide, err := client.GetGuide(ctx, "accessible-pdf-best-practices")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Guide: %s (Level: %s)\n", guide.Title, guide.AudienceLevel)
```

Learn more: [PDF Glossary](https://peasypdf.com/glossary/) · [Accessible PDF Best Practices](https://peasypdf.com/guides/accessible-pdf-best-practices/) · [How to Convert PDF to Images](https://peasypdf.com/guides/how-to-convert-pdf-to-images/)

### Search and Discovery

The API supports full-text search across all content types — tools, glossary terms, guides, use cases, and format documentation. Search results are grouped by content type, making it easy to find exactly what you need for a specific PDF workflow.

```go
// Search across all PDF content — tools, glossary, guides, and formats
results, err := client.Search(ctx, "compress pdf", nil)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Found %d tools, %d glossary terms, %d guides\n",
	len(results.Results.Tools),
	len(results.Results.Glossary),
	len(results.Results.Guides),
)

// Discover format conversion paths — what can PDF convert to?
conversions, err := client.ListConversions(ctx, &peasypdf.ListConversionsOptions{
	Source: str("pdf"), // Find all formats PDF can be converted to
})
if err != nil {
	log.Fatal(err)
}
for _, c := range conversions.Results {
	fmt.Printf("%s → %s\n", c.SourceFormat, c.TargetFormat)
}
```

Learn more: [REST API Docs](https://peasypdf.com/developers/) · [All PDF Tools](https://peasypdf.com/)

## API Client

The client wraps the [PeasyPDF REST API](https://peasypdf.com/developers/) with typed Go structs and zero external dependencies.

```go
client := peasypdf.New()
// Or with a custom base URL:
// client := peasypdf.New(peasypdf.WithBaseURL("https://custom.example.com"))
ctx := context.Background()

// List tools with pagination
tools, _ := client.ListTools(ctx, &peasypdf.ListOptions{Page: 1, Limit: 10})

// Get a specific tool by slug
tool, _ := client.GetTool(ctx, "pdf-merge")
fmt.Println(tool.Name, tool.Description)

// Search across all content
results, _ := client.Search(ctx, "compress pdf", nil)
fmt.Printf("Found %d tools\n", len(results.Results.Tools))

// Browse the glossary
glossary, _ := client.ListGlossary(ctx, &peasypdf.ListOptions{Search: str("linearization")})
for _, term := range glossary.Results {
	fmt.Printf("%s: %s\n", term.Term, term.Definition)
}

// Discover guides
guides, _ := client.ListGuides(ctx, &peasypdf.ListGuidesOptions{Category: str("optimization")})
for _, g := range guides.Results {
	fmt.Printf("%s (%s)\n", g.Title, g.AudienceLevel)
}

// List file format conversions
conversions, _ := client.ListConversions(ctx, &peasypdf.ListConversionsOptions{Source: str("pdf")})

// Get format details
format, _ := client.GetFormat(ctx, "pdf-a")
fmt.Printf("%s (%s): %s\n", format.Name, format.Extension, format.MimeType)
```

Helper for optional string parameters:

```go
func str(s string) *string { return &s }
```

### Available Methods

| Method | Description |
|--------|-------------|
| `ListTools(ctx, opts)` | List tools (paginated, filterable) |
| `GetTool(ctx, slug)` | Get tool by slug |
| `ListCategories(ctx, opts)` | List tool categories |
| `ListFormats(ctx, opts)` | List file formats |
| `GetFormat(ctx, slug)` | Get format by slug |
| `ListConversions(ctx, opts)` | List format conversions |
| `ListGlossary(ctx, opts)` | List glossary terms |
| `GetGlossaryTerm(ctx, slug)` | Get glossary term |
| `ListGuides(ctx, opts)` | List guides |
| `GetGuide(ctx, slug)` | Get guide by slug |
| `ListUseCases(ctx, opts)` | List use cases |
| `Search(ctx, query, limit)` | Search across all content |
| `ListSites(ctx)` | List Peasy sites |
| `OpenAPISpec(ctx)` | Get OpenAPI specification |

Full API documentation at [peasypdf.com/developers/](https://peasypdf.com/developers/).
OpenAPI 3.1.0 spec: [peasypdf.com/api/openapi.json](https://peasypdf.com/api/openapi.json).

## Learn More About PDF Tools

- **Tools**: [Merge PDF](https://peasypdf.com/pdf/merge-pdf/) · [Split PDF](https://peasypdf.com/pdf/split-pdf/) · [Compress PDF](https://peasypdf.com/pdf/compress-pdf/) · [Rotate PDF](https://peasypdf.com/pdf/rotate-pdf/) · [PDF to PNG](https://peasypdf.com/pdf/pdf-to-png/) · [All Tools](https://peasypdf.com/)
- **Guides**: [How to Merge PDF Files](https://peasypdf.com/guides/how-to-merge-pdf-files/) · [PDF Compression Guide](https://peasypdf.com/guides/pdf-compression-guide/) · [Accessible PDF Best Practices](https://peasypdf.com/guides/accessible-pdf-best-practices/) · [How to Convert PDF to Images](https://peasypdf.com/guides/how-to-convert-pdf-to-images/) · [All Guides](https://peasypdf.com/guides/)
- **Glossary**: [PDF](https://peasypdf.com/glossary/pdf/) · [PDF/A](https://peasypdf.com/glossary/pdfa/) · [DPI](https://peasypdf.com/glossary/dpi/) · [OCR](https://peasypdf.com/glossary/ocr/) · [Rasterization](https://peasypdf.com/glossary/rasterization/) · [All Terms](https://peasypdf.com/glossary/)
- **Formats**: [All Formats](https://peasypdf.com/formats/)
- **API**: [REST API Docs](https://peasypdf.com/developers/) · [OpenAPI Spec](https://peasypdf.com/api/openapi.json)

## Also Available

| Language | Package | Install |
|----------|---------|---------|
| **Python** | [peasy-pdf](https://pypi.org/project/peasy-pdf/) | `pip install "peasy-pdf[all]"` |
| **TypeScript** | [peasy-pdf](https://www.npmjs.com/package/peasy-pdf) | `npm install peasy-pdf` |
| **Rust** | [peasy-pdf](https://crates.io/crates/peasy-pdf) | `cargo add peasy-pdf` |
| **Ruby** | [peasy-pdf](https://rubygems.org/gems/peasy-pdf) | `gem install peasy-pdf` |

## Peasy Developer Tools

Part of the [Peasy Tools](https://peasytools.com) open-source developer ecosystem.

| Package | PyPI | npm | Go | Description |
|---------|------|-----|----|-------------|
| **peasy-pdf** | [PyPI](https://pypi.org/project/peasy-pdf/) | [npm](https://www.npmjs.com/package/peasy-pdf) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-pdf-go) | **PDF merge, split, rotate, compress — [peasypdf.com](https://peasypdf.com)** |
| peasy-image | [PyPI](https://pypi.org/project/peasy-image/) | [npm](https://www.npmjs.com/package/peasy-image) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-image-go) | Image resize, crop, convert, compress — [peasyimage.com](https://peasyimage.com) |
| peasy-audio | [PyPI](https://pypi.org/project/peasy-audio/) | [npm](https://www.npmjs.com/package/peasy-audio) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-audio-go) | Audio trim, merge, convert, normalize — [peasyaudio.com](https://peasyaudio.com) |
| peasy-video | [PyPI](https://pypi.org/project/peasy-video/) | [npm](https://www.npmjs.com/package/peasy-video) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-video-go) | Video trim, resize, thumbnails, GIF — [peasyvideo.com](https://peasyvideo.com) |
| peasy-css | [PyPI](https://pypi.org/project/peasy-css/) | [npm](https://www.npmjs.com/package/peasy-css) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-css-go) | CSS minify, format, analyze — [peasycss.com](https://peasycss.com) |
| peasy-compress | [PyPI](https://pypi.org/project/peasy-compress/) | [npm](https://www.npmjs.com/package/peasy-compress) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-compress-go) | ZIP, TAR, gzip compression — [peasytools.com](https://peasytools.com) |
| peasy-document | [PyPI](https://pypi.org/project/peasy-document/) | [npm](https://www.npmjs.com/package/peasy-document) | [Go](https://pkg.go.dev/github.com/peasytools/peasy-document-go) | Markdown, HTML, CSV, JSON conversion — [peasyformats.com](https://peasyformats.com) |
| peasytext | [PyPI](https://pypi.org/project/peasytext/) | [npm](https://www.npmjs.com/package/peasytext) | [Go](https://pkg.go.dev/github.com/peasytools/peasytext-go) | Text case conversion, slugify, word count — [peasytext.com](https://peasytext.com) |

## License

MIT
