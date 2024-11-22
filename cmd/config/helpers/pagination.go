package helpers

import (
	pagingPb "go-grpc/pb/pagination"
	"math"

	"gorm.io/gorm"
)

func Pagination(sql *gorm.DB, page int64, pagination *pagingPb.Pagination) (int64, int64){
	var total int64
	var limit int64 = 1
	var offset int64 

	
	sql.Count(&total)

	if page == 1 {
		offset = 0
	}else {
		offset = (page - 1) * limit	
	}
	pagination.Total = uint64(total)
	pagination.PerPage = uint64(limit)
	pagination.CurrentPage = uint64(page)
	pagination.LastPage = uint64(math.Ceil(float64(total) / float64(limit)))


	return offset, limit
}
