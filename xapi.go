package xapi

import (
	"encoding/json"
	"net/http"
)

// GetBodyJSON get body json
func GetBodyJSON(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&v)
	return err
}

// GetQuery get query
func GetQuery(r *http.Request) map[string][]string {
	return r.URL.Query()
}

// RenderJSON render json
func RenderJSON(w http.ResponseWriter, v interface{}) {
	_ = json.NewEncoder(w).Encode(v)
}

// Pagination pagination
func Pagination(total, offset, limit int64) *PaginationSchema {
	return &PaginationSchema{
		IsEnd:   offset+limit >= total,
		IsFirst: offset == 0,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}
}

// GetPage get page
func GetPage(data interface{}, total, offset, limit int64) *PageSchema {
	return &PageSchema{
		Data:       data,
		Pagination: Pagination(total, offset, limit),
	}
}
