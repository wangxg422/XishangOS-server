package service

import (
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mojocn/base64Captcha"
	"github.com/wangxg422/XishangOS-backend/app/module/common/util/captcha"
)

type CaptchaService struct {
}

func (m *CaptchaService) GetCaptchaImgString() (idKeyC string, base64stringC string, err error) {
	driver := captcha.CaptchaDriver().ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, captcha.GetRedisStore())
	idKeyC, base64stringC, _, err = c.Generate()
	return
}

// VerifyString 验证验证码是否正确
func (m *CaptchaService) VerifyString(id, answer string) bool {
	c := base64Captcha.NewCaptcha(captcha.CaptchaDriver(), captcha.GetRedisStore())
	answer = gstr.ToLower(answer)
	return c.Verify(id, answer, true)
}
