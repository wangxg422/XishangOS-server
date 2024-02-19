package response

type PaginationRes struct {
	PageNo int `json:"pageNo"` // 当前页码
	Total  int `json:"total"`  // 总数据量
	List   any `json:"list"`
}
