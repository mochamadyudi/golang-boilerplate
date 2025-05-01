package utils

import (
	"strconv"

	"core.yuyuid.id/internal/database"
	"github.com/gofiber/fiber/v2"
)

func Paginate(c *fiber.Ctx) Pagination {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 1 {
		limit = 1
	}

	offset := (page - 1) * limit

	return Pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}

type RowsPagiationInterface[T any] struct {
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func RowsPagination[T any](
	model *T,
	pagination Pagination,
	Result []T) error {
	db := database.DB

	var total int64

	if err := db.Model(&model).Count(&total).Error; err != nil {
		return err
	}

	if err := db.Model(&model).Limit(pagination.Limit).Offset(pagination.Offset).Find(&Result).Error; err != nil {
		return err
	}

	totalPages := int64(0)
	if total > 0 {
		totalPages = total / int64(pagination.Limit)
		if total%int64(pagination.Limit) > 0 {
			totalPages++
		}
	}
	pagination.Total = int(total)

	pagination.Maxpage = int(totalPages)

	return nil
}
