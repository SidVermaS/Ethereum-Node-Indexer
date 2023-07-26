package helpers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ConvertStringToFloat(data string) (float64, error) {
	number, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
func ConvertStringToUInt(data string) (uint, error) {
	number, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	return uint(number), err
}
func ConvertStringToUInt64(data string) (uint64, error) {
	number, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(number), nil
}
func ConvertStringToInt(data string) (int, error) {
	number, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	return number, nil
}
func GetPaginationValues(ctx *fiber.Ctx) (int, int) {
	pageQuery := ctx.Query("page")
	limitQuery := ctx.Query("limit")
	page, _ := ConvertStringToInt(pageQuery)
	limit, _ := ConvertStringToInt(limitQuery)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 1
	}
	offset := (page - 1) * limit
	return offset, limit
}
