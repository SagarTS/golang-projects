package helper

import (
	"go-pagination/initializers"
	"math"
)

type PaginationData struct{
	NextPage int
	PreviousPage int
	CurrentPage int
	TotalPages int
	TwoAfter int
	TwoBelow int
	ThreeAfter int
	Offset int
	BaseURL string
}

func GetPaginationData(page int, perPage int, model interface{}, baseURL string) PaginationData{
	// Calculate total page
	var totalRows int64
	initializers.DB.Model(model).Count(&totalRows)
	totalPages := math.Ceil(float64(totalRows/int64(perPage)))

	offset := (page-1) * perPage


	return PaginationData{
		NextPage: page + 1,
		PreviousPage: page - 1,
		CurrentPage: page,
		TotalPages: int(totalPages),
		TwoAfter: page + 2,
		TwoBelow: page -2,
		ThreeAfter: page + 3,
		Offset: offset,
		BaseURL: baseURL,
	}
}