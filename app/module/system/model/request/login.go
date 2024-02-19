package request

import "github.com/wangxg422/XishangOS-backend/app/base/model/request"

type SysLoginReq struct {
	Username   string `json:"username" form:"username" v:"username@required|#用户名不能为空"`
	Password   string `json:"password" form:"password" v:"password@required|#密码不能为空"`
	VerifyCode string `json:"verifyCode" form:"verifyCode"`
	VerifyKey  string `json:"verifyKey" form:"verifyKey"`
}

type SysUserListReq struct {
	DeptId             int64  `json:"deptId" form:"deptId"`
	Email              string `json:"email" form:"email"`
	Mobile             string `json:"mobile" form:"mobile"`
	Sex                int8   `json:"sex" form:"sex"`
	Status             int8   `json:"status" form:"status"`
	Keyword            string `json:"keyword" form:"keyword"`
	OrderBy            string `json:"orderBy" form:"orderBy"`
	Order              string `json:"order" form:"order"`
	request.Pagination `json:"pagination"`
}

type SysUserCreateReq struct {
	Id       int64   `json:"id,string"`
	UserName string  `json:"userName" v:"required#用户名不能为空"`
	Password string  `json:"password" v:"required#密码不能为空"`
	DeptId   int64   `json:"deptId"`                 //所属部门
	Email    string  `json:"email" v:"email#邮箱格式错误"` //邮箱
	NickName string  `json:"nickName"`
	Mobile   string  `json:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	Remark   string  `json:"remark"`
	Sex      int8    `json:"sex"`
	Status   int8    `json:"status"`
	IsAdmin  int     `json:"isAdmin"`
	Address  string  `json:"address"`
	PostIds  []int64 `json:"postIds"`
	RoleIds  []int64 `json:"roleIds"`
}

type SysUserUpdateReq struct {
	Id       int64   `json:"id,string" v:"required#用户名不能为空"`
	UserName string  `json:"userName" v:"required#用户名不能为空"`
	Password string  `json:"password" v:"required#密码不能为空"`
	DeptId   int64   `json:"deptId"`                 //所属部门
	Email    string  `json:"email" v:"email#邮箱格式错误"` //邮箱
	NickName string  `json:"nickName"`
	Mobile   string  `json:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	Remark   string  `json:"remark"`
	Sex      int8    `json:"sex"`
	Status   int8    `json:"status"`
	IsAdmin  int     `json:"isAdmin"`
	Address  string  `json:"address"`
	RoleIds  []int64 `json:"roleIds"`
	PostIds  []int64 `json:"postIds"`
}
