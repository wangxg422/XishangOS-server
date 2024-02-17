package captcha

import (
	"github.com/mojocn/base64Captcha"
	"github.com/wangxg422/XishangOS-backend/global"
)

var driver *base64Captcha.DriverString

func CaptchaDriver() *base64Captcha.DriverString {
	if driver == nil {
		d := &base64Captcha.DriverString{
			Height:          global.AppConfig.Captcha.ImgHeight,
			Width:           global.AppConfig.Captcha.ImgWidth,
			NoiseCount:      global.AppConfig.Captcha.NoiseCount,
			ShowLineOptions: global.AppConfig.Captcha.ShowLineOptions,
			Length:          global.AppConfig.Captcha.KeyLong,
			Source:          global.AppConfig.Captcha.Source,
			Fonts:           []string{"chromohv.ttf"},
		}
		driver = d
	}
	return driver
}
