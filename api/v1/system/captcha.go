package system

import (
	"github.com/niudaii/gin-server/global"
	"github.com/niudaii/gin-server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

// Captcha 生成验证码
func (a *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.Server.Captcha.ImgHeight, global.Server.Captcha.ImgWidth, global.Server.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		response.ErrorWithMessage("生成验证码失败", err, c)
	} else {
		response.Ok(response.CaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.Server.Captcha.KeyLong,
		}, "", c)
	}
}
