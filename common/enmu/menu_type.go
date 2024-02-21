package enmu

type MenuType int8

const (
	menuTypeDir  = 1
	menuTypeMenu = 2
	menuTypeBtn  = 3
)

var menuTypeMap map[int8]string

func init() {
	menuTypeMap = make(map[int8]string)
	menuTypeMap[menuTypeDir] = "目录"
	menuTypeMap[menuTypeMenu] = "菜单"
	menuTypeMap[menuTypeBtn] = "按钮"
}

func (m MenuType) Desc() string {
	return menuTypeMap[int8(m)]
}

func (m MenuType) Value() int8 {
	return int8(m)
}

func (m MenuType) Size() int {
	return len(menuTypeMap)
}

func (m MenuType) Equals(value int8) bool {
	return int8(m) == value
}
