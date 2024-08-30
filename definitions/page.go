package definitions

// PaginationRequest represents a request for paginated data.
// It includes the page number and the number of items per page.
type PaginationRequest struct {
	Page    int `json:"page" query:"page"`        // Page number of the paginated data.
	PerPage int `json:"perPage" query:"per_page"` // Number of items per page.
}

// PaginatedData represents a paginated response.
// It includes the data and the total count of items.
type PaginatedData[T any] struct {
	Data       []T `json:"data"`       // Data is the slice of items for the current page.
	TotalCount int `json:"totalCount"` // TotalCount is the total number of items available.
}

// Offset calculates the offset for the paginated data.
// It ensures that the page number and items per page are valid.
// Returns the offset as an integer.
func (r *PaginationRequest) Offset() int {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PerPage < 1 {
		r.PerPage = 10
	}
	return (r.Page - 1) * r.PerPage
}

// Limit calculates the limit for the paginated data.
// It ensures that the items per page are valid.
// Returns the limit as an integer.
func (r *PaginationRequest) Limit() int {
	if r.PerPage < 1 {
		r.PerPage = 10
	}
	return r.PerPage
}
