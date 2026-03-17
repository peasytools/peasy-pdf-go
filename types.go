package peasypdf

// PaginatedResponse represents a paginated API response.
type PaginatedResponse[T any] struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []T     `json:"results"`
}

// Tool represents a PDF tool.
type Tool struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	URL         string `json:"url"`
}

// Category represents a tool category.
type Category struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ToolCount   int    `json:"tool_count"`
}

// Format represents a file format.
type Format struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Extension   string `json:"extension"`
	MimeType    string `json:"mime_type"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

// Conversion represents a format conversion.
type Conversion struct {
	Source      string `json:"source"`
	Target      string `json:"target"`
	Description string `json:"description"`
	ToolSlug    string `json:"tool_slug"`
}

// GlossaryTerm represents a glossary entry.
type GlossaryTerm struct {
	Slug       string `json:"slug"`
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Category   string `json:"category"`
}

// Guide represents a how-to guide.
type Guide struct {
	Slug          string `json:"slug"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	AudienceLevel string `json:"audience_level"`
	WordCount     int    `json:"word_count"`
}

// UseCase represents an industry use case.
type UseCase struct {
	Slug     string `json:"slug"`
	Title    string `json:"title"`
	Industry string `json:"industry"`
}

// Site represents a Peasy Tools site.
type Site struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	URL    string `json:"url"`
}

// SearchResult represents unified search results.
type SearchResult struct {
	Query   string           `json:"query"`
	Results SearchCategories `json:"results"`
}

// SearchCategories groups search results by type.
type SearchCategories struct {
	Tools    []Tool         `json:"tools"`
	Formats  []Format       `json:"formats"`
	Glossary []GlossaryTerm `json:"glossary"`
}

// ListOptions are common parameters for paginated list endpoints.
type ListOptions struct {
	Page     int
	Limit    int
	Category string
	Search   string
}

// ListGuidesOptions adds audience level filtering.
type ListGuidesOptions struct {
	Page          int
	Limit         int
	Category      string
	AudienceLevel string
	Search        string
}

// ListConversionsOptions adds source/target filtering.
type ListConversionsOptions struct {
	Page   int
	Limit  int
	Source string
	Target string
}

// SearchOptions for the search endpoint.
type SearchOptions struct {
	Limit int
}
