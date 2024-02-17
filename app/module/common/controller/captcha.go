package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/common/service"
	"github.com/wangxg422/XishangOS-backend/common/result"
)

type CaptchaController struct {
}

func (m *CaptchaController) GetCaptchaImgString(c *gin.Context) {
	idKeyC, base64stringC, err := service.AppCaptchaService.GetCaptchaImgString()
	if err != nil {
		result.FailWithMessage(err.Error(), c)
		return
	}
	codeMap := make(result.Map)
	codeMap["key"] = idKeyC
	codeMap["img"] = base64stringC
	result.OkWithData(codeMap, c)
}
