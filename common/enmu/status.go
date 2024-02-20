package enmu

type Status int8

const (
	statusNormal   = 1
	statusDisabled = 2
)

var statusMap map[int8]string

func init() {
	statusMap = make(map[int8]string)
	statusMap[statusNormal] = "正常"
	statusMap[statusDisabled] = "停用"
}

func (m Status) Desc() string {
	return statusMap[int8(m)]
}

func (m Status) Value() int8 {
	return int8(m)
}

func (m Status) Size() int {
	return len(statusMap)
}

func (m Status) Equals(value int8) bool {
	return int8(m) == value
}
