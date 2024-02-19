package request

type Pagination struct {
	PageNo   int `json:"pageNo" form:"pageNo"`     // 当前页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}
