package helpers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ConvertStringToUInt64(data string) uint64 {
	number, _ := strconv.ParseUint(data, 10, 64)
	return number
}
func ConvertStringToInt(data string) int {
	number, _ := strconv.Atoi(data)
	return number
}
func GetPaginationValues(ctx *fiber.Ctx) (int, int) {
	pageQuery := ctx.Query("page")
	limitQuery := ctx.Query("limit")
	page := ConvertStringToInt(pageQuery)
	limit := ConvertStringToInt(limitQuery)
	if page <= 0 {
		page = 1
	}	
	if limit <= 0 {
		limit = 1
	}
	offset := (page - 1) * limit
	return offset, limit
}
