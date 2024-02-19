package request

type SysDeptListReq struct {
	Keyword string `json:"keyword" form:"keyword"`
	Status  int8   `json:"status" form:"status"`
}

type SysDeptCreateReq struct {
	DeptName string `json:"deptName" v:"required#部门名称不能为空"`
	DeptCode string `json:"deptCode" v:"required#部门编码不能为空"`
	ParentId int64  `json:"parentId,string"`
	Leader   string `json:"leader" v:"email#邮箱格式错误"`
	Phone    string `json:"Phone" v:"required|phone#手机号不能为空|手机号格式错误"`
	Email    string `json:"email"`
	Status   int8   `json:"status"`
	Address  string `json:"address"`
	Remark   string `json:"remark"`
}

type SysDeptUpdateReq struct {
	Id       int64  `json:"id,string"`
	DeptName string `json:"deptName" v:"required#部门名称不能为空"`
	DeptCode string `json:"deptCode" v:"required#部门编码不能为空"`
	ParentId int64  `json:"parentId,string"`
	Leader   string `json:"leader" v:"email#邮箱格式错误"`
	Phone    string `json:"Phone" v:"required|phone#手机号不能为空|手机号格式错误"`
	Email    string `json:"email"`
	Status   int8   `json:"status"`
	Address  string `json:"address"`
	Remark   string `json:"remark"`
}
