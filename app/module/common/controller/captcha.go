package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxg422/XishangOS-backend/app/module/common/service"
	"github.com/wangxg422/XishangOS-backend/common/result"
	"github.com/wangxg422/XishangOS-backend/global"
)

type CaptchaController struct {
}

func (m *CaptchaController) GetCaptchaImgString(c *gin.Context) {
	// 检查验证码功能是否启用
	if !global.AppConfig.Captcha.Enabled {
		result.FailWithMessage("验证码功能未启用", c)
		return
	}
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
