package response

type LoginRes struct {
	Username    string      `json:"username"`
	UserId      int64       `json:"userId,string"`
	DeptId      int64       `json:"deptId,string"`
	MenuList    []*UserMenu `json:"menuList"`
	Permissions []string    `json:"permissions"`
	Token       string      `json:"token"`
}
