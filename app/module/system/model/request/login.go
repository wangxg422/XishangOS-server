package request

type SysLoginReq struct {
	Username   string `json:"username" form:"username" v:"username@required|#用户名不能为空"`
	Password   string `json:"password" form:"password" v:"password@required|#密码不能为空"`
	VerifyCode string `json:"verifyCode" form:"verifyCode"`
	VerifyKey  string `json:"verifyKey" form:"verifyKey"`
}

type SysUserCreateUpdateReq struct {
	Id       int64   `json:"id"`
	UserName string  `json:"userName" v:"required#用户名不能为空"`
	Password string  `json:"password" v:"required#密码不能为空"`
	DeptId   int64   `json:"deptId" v:"required#用户部门不能为空"` //所属部门
	Email    string  `json:"email" v:"email#邮箱格式错误"`       //邮箱
	NickName string  `json:"nickName" v:"required#用户昵称不能为空"`
	Mobile   string  `json:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	PostIds  []int64 `json:"postIds"`
	Remark   string  `json:"remark"`
	RoleIds  []int64 `json:"roleIds"`
	Sex      int8    `json:"sex"`
	Status   int8    `json:"status"`
	IsAdmin  int     `json:"isAdmin"`
	Address  string  `json:"address"`
}
