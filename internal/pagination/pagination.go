package pagination

import (
	"math"

	"github.com/jinzhu/gorm"
)

func Paginate(db *gorm.DB, value interface{}, pageSize, page int32, orderBy string) (int, func(db *gorm.DB) *gorm.DB) {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))

	return totalPages, func(db *gorm.DB) *gorm.DB {
		return db.Offset(page).Limit(pageSize).Order(orderBy)
	}
}
