package enmu

type DelFlag int8

const (
	delFlagNormal  = 1
	delFlagDeleted = 2
)

var delFlagMap map[int]string

func init() {
	delFlagMap = make(map[int]string)
	delFlagMap[delFlagNormal] = "正常"
	delFlagMap[delFlagDeleted] = "已删除"
}

func (m DelFlag) Desc() string {
	return delFlagMap[int(m)]
}

func (m DelFlag) Value() int8 {
	return int8(m)
}

func (m DelFlag) Size() int {
	return len(delFlagMap)
}

func (m DelFlag) Equals(value int8) bool {
	return int8(m) == value
}
