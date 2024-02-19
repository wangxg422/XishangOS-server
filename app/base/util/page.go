package util

import (
	"github.com/wangxg422/XishangOS-backend/common/constant"
)

func PageLimitOffset(pageNo int, pageSize int) (int, int) {
	limit := constant.DefaultPageLimit

	if pageSize <= 0 {
		limit = constant.DefaultPageLimit
	}
	if pageNo <= 0 {
		pageNo = 1
	}

	return limit, (pageNo - 1) * limit
}
