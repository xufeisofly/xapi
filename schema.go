package xapi

// PaginationSchema pagination
type PaginationSchema struct {
	IsEnd   bool  `json:"is_end"`
	IsFirst bool  `json:"is_first"`
	Total   int64 `json:"total"`
	Offset  int64 `json:"offset"`
	Limit   int64 `json:"limit"`
}

// PageSchema page
type PageSchema struct {
	Data       interface{}       `json:"data"`
	Pagination *PaginationSchema `json:"pagination"`
}
