// Package peasypdf provides a Go client for the PeasyPDF API.
//
// PeasyPDF offers PDF tools including merge, split, compress, and convert.
// This client requires no authentication and has zero external dependencies.
//
// Usage:
//
//	client := peasypdf.New()
//	tools, err := client.ListTools(ctx)
//	term, err := client.GetGlossaryTerm(ctx, "pdf-a")
package peasypdf

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// DefaultBaseURL is the default PeasyPDF API base URL.
const DefaultBaseURL = "https://peasypdf.com"

// DefaultTimeout is the default HTTP client timeout.
const DefaultTimeout = 30 * time.Second

// Client is the PeasyPDF API client.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// Option configures the Client.
type Option func(*Client)

// WithBaseURL overrides the default base URL.
func WithBaseURL(u string) Option {
	return func(c *Client) { c.baseURL = strings.TrimRight(u, "/") }
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.httpClient.Timeout = d }
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// New creates a new PeasyPDF API client.
func New(opts ...Option) *Client {
	c := &Client{
		baseURL:    DefaultBaseURL,
		httpClient: &http.Client{Timeout: DefaultTimeout},
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

func (c *Client) doRequest(ctx context.Context, path string, params url.Values) ([]byte, error) {
	u := c.baseURL + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: create request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: http request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: read response: %w", err)
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, &NotFoundError{Resource: "resource", Identifier: path}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, &PeasyError{StatusCode: resp.StatusCode, Message: string(body)}
	}
	return body, nil
}

func decodePaginated[T any](data []byte, method string) (*PaginatedResponse[T], error) {
	var pr PaginatedResponse[T]
	if err := json.Unmarshal(data, &pr); err != nil {
		return nil, fmt.Errorf("peasypdf: decode %s: %w", method, err)
	}
	return &pr, nil
}

func buildListParams(opts ListOptions) url.Values {
	p := url.Values{}
	if opts.Page > 0 {
		p.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.Limit > 0 {
		p.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.Category != "" {
		p.Set("category", opts.Category)
	}
	if opts.Search != "" {
		p.Set("search", opts.Search)
	}
	return p
}

func buildPath(base, slug string) string {
	return fmt.Sprintf("%s%s/", base, url.PathEscape(slug))
}

func applyListOpts(opts []ListOptions) ListOptions {
	if len(opts) > 0 {
		return opts[0]
	}
	return ListOptions{}
}

// --- Tools ---

// ListTools returns a paginated list of tools.
func (c *Client) ListTools(ctx context.Context, opts ...ListOptions) (*PaginatedResponse[Tool], error) {
	o := applyListOpts(opts)
	data, err := c.doRequest(ctx, "/api/v1/tools/", buildListParams(o))
	if err != nil {
		return nil, fmt.Errorf("peasypdf: list tools: %w", err)
	}
	return decodePaginated[Tool](data, "tools")
}

// GetTool returns a single tool by slug.
func (c *Client) GetTool(ctx context.Context, slug string) (*Tool, error) {
	data, err := c.doRequest(ctx, buildPath("/api/v1/tools/", slug), nil)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: get tool: %w", err)
	}
	var t Tool
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, fmt.Errorf("peasypdf: decode tool: %w", err)
	}
	return &t, nil
}

// --- Categories ---

// ListCategories returns a paginated list of categories.
func (c *Client) ListCategories(ctx context.Context, opts ...ListOptions) (*PaginatedResponse[Category], error) {
	o := applyListOpts(opts)
	data, err := c.doRequest(ctx, "/api/v1/categories/", buildListParams(o))
	if err != nil {
		return nil, fmt.Errorf("peasypdf: list categories: %w", err)
	}
	return decodePaginated[Category](data, "categories")
}

// --- Formats ---

// ListFormats returns a paginated list of file formats.
func (c *Client) ListFormats(ctx context.Context, opts ...ListOptions) (*PaginatedResponse[Format], error) {
	o := applyListOpts(opts)
	data, err := c.doRequest(ctx, "/api/v1/formats/", buildListParams(o))
	if err != nil {
		return nil, fmt.Errorf("peasypdf: list formats: %w", err)
	}
	return decodePaginated[Format](data, "formats")
}

// GetFormat returns a single format by slug.
func (c *Client) GetFormat(ctx context.Context, slug string) (*Format, error) {
	data, err := c.doRequest(ctx, buildPath("/api/v1/formats/", slug), nil)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: get format: %w", err)
	}
	var f Format
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("peasypdf: decode format: %w", err)
	}
	return &f, nil
}

// --- Conversions ---

// ListConversions returns a paginated list of conversions.
func (c *Client) ListConversions(ctx context.Context, opts ...ListConversionsOptions) (*PaginatedResponse[Conversion], error) {
	var o ListConversionsOptions
	if len(opts) > 0 {
		o = opts[0]
	}
	p := url.Values{}
	if o.Page > 0 {
		p.Set("page", strconv.Itoa(o.Page))
	}
	if o.Limit > 0 {
		p.Set("limit", strconv.Itoa(o.Limit))
	}
	if o.Source != "" {
		p.Set("source", o.Source)
	}
	if o.Target != "" {
		p.Set("target", o.Target)
	}
	data, err := c.doRequest(ctx, "/api/v1/conversions/", p)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: list conversions: %w", err)
	}
	return decodePaginated[Conversion](data, "conversions")
}

// --- Glossary ---

// ListGlossary returns a paginated list of glossary terms.
func (c *Client) ListGlossary(ctx context.Context, opts ...ListOptions) (*PaginatedResponse[GlossaryTerm], error) {
	o := applyListOpts(opts)
	data, err := c.doRequest(ctx, "/api/v1/glossary/", buildListParams(o))
	if err != nil {
		return nil, fmt.Errorf("peasypdf: list glossary: %w", err)
	}
	return decodePaginated[GlossaryTerm](data, "glossary")
}

// GetGlossaryTerm returns a single glossary term by slug.
func (c *Client) GetGlossaryTerm(ctx context.Context, slug string) (*GlossaryTerm, error) {
	data, err := c.doRequest(ctx, buildPath("/api/v1/glossary/", slug), nil)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: get glossary term: %w", err)
	}
	var t GlossaryTerm
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, fmt.Errorf("peasypdf: decode glossary term: %w", err)
	}
	return &t, nil
}

// --- Guides ---

// ListGuides returns a paginated list of guides.
func (c *Client) ListGuides(ctx context.Context, opts ...ListGuidesOptions) (*PaginatedResponse[Guide], error) {
	var o ListGuidesOptions
	if len(opts) > 0 {
		o = opts[0]
	}
	p := url.Values{}
	if o.Page > 0 {
		p.Set("page", strconv.Itoa(o.Page))
	}
	if o.Limit > 0 {
		p.Set("limit", strconv.Itoa(o.Limit))
	}
	if o.Category != "" {
		p.Set("category", o.Category)
	}
	if o.AudienceLevel != "" {
		p.Set("audience_level", o.AudienceLevel)
	}
	if o.Search != "" {
		p.Set("search", o.Search)
	}
	data, err := c.doRequest(ctx, "/api/v1/guides/", p)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: list guides: %w", err)
	}
	return decodePaginated[Guide](data, "guides")
}

// GetGuide returns a single guide by slug.
func (c *Client) GetGuide(ctx context.Context, slug string) (*Guide, error) {
	data, err := c.doRequest(ctx, buildPath("/api/v1/guides/", slug), nil)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: get guide: %w", err)
	}
	var g Guide
	if err := json.Unmarshal(data, &g); err != nil {
		return nil, fmt.Errorf("peasypdf: decode guide: %w", err)
	}
	return &g, nil
}

// --- Use Cases ---

// ListUseCases returns a paginated list of use cases.
func (c *Client) ListUseCases(ctx context.Context, opts ...ListOptions) (*PaginatedResponse[UseCase], error) {
	o := applyListOpts(opts)
	p := url.Values{}
	if o.Page > 0 {
		p.Set("page", strconv.Itoa(o.Page))
	}
	if o.Limit > 0 {
		p.Set("limit", strconv.Itoa(o.Limit))
	}
	if o.Category != "" {
		p.Set("industry", o.Category)
	}
	if o.Search != "" {
		p.Set("search", o.Search)
	}
	data, err := c.doRequest(ctx, "/api/v1/use-cases/", p)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: list use cases: %w", err)
	}
	return decodePaginated[UseCase](data, "use cases")
}

// --- Search ---

// Search performs a unified search across tools, formats, and glossary.
func (c *Client) Search(ctx context.Context, query string, opts ...SearchOptions) (*SearchResult, error) {
	p := url.Values{"q": {query}}
	if len(opts) > 0 && opts[0].Limit > 0 {
		p.Set("limit", strconv.Itoa(opts[0].Limit))
	}
	data, err := c.doRequest(ctx, "/api/v1/search/", p)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: search: %w", err)
	}
	var sr SearchResult
	if err := json.Unmarshal(data, &sr); err != nil {
		return nil, fmt.Errorf("peasypdf: decode search: %w", err)
	}
	return &sr, nil
}

// --- Sites ---

// ListSites returns all Peasy Tools sites.
func (c *Client) ListSites(ctx context.Context) (*PaginatedResponse[Site], error) {
	data, err := c.doRequest(ctx, "/api/v1/sites/", nil)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: list sites: %w", err)
	}
	return decodePaginated[Site](data, "sites")
}

// --- OpenAPI ---

// OpenAPISpec returns the raw OpenAPI spec as JSON.
func (c *Client) OpenAPISpec(ctx context.Context) (json.RawMessage, error) {
	data, err := c.doRequest(ctx, "/api/openapi.json", nil)
	if err != nil {
		return nil, fmt.Errorf("peasypdf: openapi spec: %w", err)
	}
	return json.RawMessage(data), nil
}
