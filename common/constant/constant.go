package constant

const (
	EnvConfigFile     = "CONFIG_FILE"
	DefaultConfigFile = "config.yml"

	DEBUG   = "debug"
	RELEASE = "release"

	HTTP  = "http://"
	HTTPS = "https://"

	// LAYOUT Layout组件标识
	LAYOUT = "Layout"
	// INNER_LINK InnerLink组件标识
	INNER_LINK = "InnerLink"
	// PARENT_VIEW ParentView组件标识
	PARENT_VIEW = "ParentView"

	NoRedirect = "noRedirect"

	Comma = ","

	DelFlagField = "del_flag"

	TokenName = "X-ACCESS-TOKEN"

	// DefaultPageLimit 默认查询分页每页大小
	DefaultPageLimit = 10

	TraceCtx = "TraceCtx"
	TraceKey = "TraceKey"
)

var UpdateOmit = []string{"create_time", "create_by"}
