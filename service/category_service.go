package service

import (
	"github.com/adibhauzan/azahri_mart_be/model/web/request"
	"github.com/adibhauzan/azahri_mart_be/model/web/response"
)

type CategoryService interface {
	Create(request request.CreateCategoryRequest) response.CreateCategoryResponse
}
