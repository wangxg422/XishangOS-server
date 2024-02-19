package param

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ParamInt64(key string, c *gin.Context) (int64, error) {
	v := c.Param(key)
	return strconv.ParseInt(v, 10, 64)
}
