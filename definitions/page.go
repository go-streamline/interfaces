package definitions

type PaginationRequest struct {
	Page    int `json:"page" query:"page"`
	PerPage int `json:"per_page" query:"per_page"`
}

type PaginatedData[T any] struct {
	Data       []T
	TotalCount int
}

func (r *PaginationRequest) Offset() int {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.PerPage < 1 {
		r.PerPage = 10
	}
	return (r.Page - 1) * r.PerPage
}

func (r *PaginationRequest) Limit() int {
	if r.PerPage < 1 {
		r.PerPage = 10
	}
	return r.PerPage
}
