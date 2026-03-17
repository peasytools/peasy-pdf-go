# peasy-pdf-go

[![Go Reference](https://pkg.go.dev/badge/github.com/peasytools/peasy-pdf-go.svg)](https://pkg.go.dev/github.com/peasytools/peasy-pdf-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/peasytools/peasy-pdf-go)](https://goreportcard.com/report/github.com/peasytools/peasy-pdf-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [PeasyPDF](https://peasypdf.com) API — PDF merge, split, rotate, and compress. Zero dependencies beyond the Go standard library.

Built from [PeasyPDF](https://peasypdf.com), a comprehensive PDF toolkit offering free online tools for merging, splitting, rotating, compressing, and converting PDF documents with detailed format guides and glossary.

> **Try the interactive tools at [peasypdf.com](https://peasypdf.com)** — [PDF Tools](https://peasypdf.com/), [PDF Glossary](https://peasypdf.com/glossary/), [PDF Guides](https://peasypdf.com/guides/)

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

## Learn More

- **Tools**: [PDF Merge](https://peasypdf.com/tools/pdf-merge/) · [PDF Split](https://peasypdf.com/tools/pdf-split/) · [PDF Compress](https://peasypdf.com/tools/pdf-compress/) · [All Tools](https://peasypdf.com/)
- **Guides**: [PDF Metadata Guide](https://peasypdf.com/guides/pdf-metadata/) · [All Guides](https://peasypdf.com/guides/)
- **Glossary**: [PDF/A](https://peasypdf.com/glossary/pdf-a/) · [Linearization](https://peasypdf.com/glossary/linearization/) · [All Terms](https://peasypdf.com/glossary/)
- **Formats**: [PDF](https://peasypdf.com/formats/pdf/) · [PDF/A](https://peasypdf.com/formats/pdf-a/) · [All Formats](https://peasypdf.com/formats/)
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
