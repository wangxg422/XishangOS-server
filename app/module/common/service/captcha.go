package service

import (
	"github.com/mojocn/base64Captcha"
	"github.com/wangxg422/XishangOS-backend/app/module/common/util/captcha"
	"strings"
)

type Captcha struct {
}

func (m *Captcha) GetCaptchaImgString() (idKeyC string, base64stringC string, err error) {
	driver := captcha.CaptchaDriver().ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, captcha.GetRedisStore())
	idKeyC, base64stringC, _, err = c.Generate()
	return
}

// VerifyString 验证验证码是否正确
func (m *Captcha) VerifyString(id, answer string) bool {
	c := base64Captcha.NewCaptcha(captcha.CaptchaDriver(), captcha.GetRedisStore())
	answer = strings.ToLower(answer)
	return c.Verify(id, answer, true)
}
