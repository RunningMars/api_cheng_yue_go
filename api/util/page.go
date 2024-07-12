package util

type Page[T interface{}] struct {
	Records *[]T  `json:"records"`
	Total   int64 `json:"total"`
	Size    int   `json:"size"`
	Current int   `json:"current"`
}

func NewPage[T interface{}](records *[]T, current, size int, total int64) *Page[T] {

	return &Page[T]{
		Records: records,
		Total:   total,
		Size:    size,
		Current: current,
	}
}
