package response

type UserInfo struct {
	Username    string      `json:"username"`
	UserId      int64       `json:"userId,string"`
	DeptId      int64       `json:"deptId,string"`
	MenuList    []*UserMenu `json:"menuList"`
	Permissions []string    `json:"permissions"`
}

type LoginRes struct {
	Token    string    `json:"token"`
	UserInfo *UserInfo `json:"userInfo"`
}
