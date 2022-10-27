package request

type Login struct {
	Username  string `json:"username" binding:"required"`  // 用户名
	Password  string `json:"password" binding:"required"`  // 密码
	Captcha   string `json:"captcha" binding:"required"`   // 验证码
	CaptchaId string `json:"captchaId" binding:"required"` // 验证码ID
}

type ChangePassword struct {
	Password string `json:"password" binding:"required"` // 新密码
}
