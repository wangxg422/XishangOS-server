package enmu

type UserStatus int8

const (
	userStatus_disable = 0
	userStatus_normal  = 1
)

func (m UserStatus) userStatusMap() map[int8]string {
	userStatus := make(map[int8]string)
	userStatus[0] = "禁用"
	userStatus[1] = "正常"

	return userStatus
}

func (m UserStatus) Desc() string {
	return m.userStatusMap()[int8(m)]
}

func (m UserStatus) Value() int8 {
	return int8(m)
}

func (m UserStatus) Size() int {
	return len(m.userStatusMap())
}

func (m UserStatus) Equals(value int8) bool {
	return int8(m) == value
}
