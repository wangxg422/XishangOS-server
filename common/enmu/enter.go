package enmu

const (
	DelFlagNormal  DelFlag = DelFlag(delFlagNormal)
	DelFlagDeleted DelFlag = DelFlag(delFlagDeleted)

	StatusNormal   Status = Status(statusNormal)
	StatusDisabled Status = Status(statusDisabled)

	MenuTypeDir  MenuType = MenuType(menuTypeDir)
	MenuTypeMenu MenuType = MenuType(menuTypeMenu)
	MenuTypeBTN  MenuType = MenuType(menuTypeBtn)

	MenuIsFrame    MenuFrame = MenuFrame(Menu_Frame)
	MenuIsNotFrame MenuFrame = MenuFrame(Menu_Not_Frame)

	MenuIsCache    MenuCache = MenuCache(Menu_Cache)
	MenuIsNotCache MenuCache = MenuCache(Menu_Not_Cache)

	SexM = Sex(Sex_M)
	SexF = Sex(Sex_F)
	SexU = Sex(Sex_U)
)
