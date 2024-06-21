package common

type PageRequest struct {
	PageNum  int `json:"pageNum" validate:"required"`
	PageSize int `json:"pageSize" validate:"required"`
}
