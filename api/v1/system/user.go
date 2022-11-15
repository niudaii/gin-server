package system

import (
	"gin-server/model/common/request"
	"gin-server/model/common/response"
	systemModel "gin-server/model/system"
	"gin-server/utils"
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

// Login 用户登录
func (a *BaseApi) Login(c *gin.Context) {
	var req request.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMessage("参数校验失败", err, c)
		return
	}
	data := map[string]interface{}{
		"isLogin": true,
	}
	// 判断验证码
	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		response.UnAuth(data, "验证码错误", c)
		return
	}
	// 判断用户是否存在
	user, err := userService.Login(req.Username, req.Password)
	if err != nil {
		response.UnAuth(data, "账户或密码错误", c)
		return
	}
	// 登录成功,返回 token
	j := utils.NewJwt()
	claims := j.CreateClaims(utils.BaseClaims{
		UUID:        user.UUID,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	var token string
	token, err = j.CreateToken(claims)
	if err != nil {
		response.ErrorWithMessage("生成token失败", err, c)
	} else {
		data["token"] = token
		response.Ok(data, "登录成功", c)
	}
}

// Logout 用户注销
func (a *BaseApi) Logout(c *gin.Context) {
	response.OkWithMessage("注销成功", c)
}

type UserApi struct{}

// GetInfo 查询用户信息
func (a *UserApi) GetInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	if user, err := userService.Select(uuid); err != nil {
		response.ErrorWithMessage("查询用户信息失败", err, c)
	} else {
		data := map[string]interface{}{
			"name": user.Username,
			// antd-vue
			"role": map[string]interface{}{
				"permissions": []map[string]interface{}{
					{
						"permissionId": user.Authority.AuthorityName,
					},
				},
			},
			// vue-element-admin
			"roles": []string{
				user.Authority.AuthorityName,
			},
			"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		}
		response.Ok(data, "查询用户信息成功", c)
	}
}

// UserMenu 查询用户菜单
func (a *UserApi) GetMenu(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	user, err := userService.Select(uuid)
	if err != nil {
		response.ErrorWithMessage("查询用户菜单失败", err, c)
		return
	}
	if err = authorityService.GetMenus(&user.Authority); err != nil {
		response.ErrorWithMessage("查询用户菜单失败", err, c)
	} else {
		response.Ok(user.Authority.Menus, "", c)
	}
}

// Create 创建用户
func (a *UserApi) Create(c *gin.Context) {
	var req systemModel.User
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMessage("参数校验失败", err, c)
		return
	}
	if err := userService.Insert(req); err != nil {
		response.ErrorWithMessage("创建用户失败", err, c)
	} else {
		response.OkWithMessage("创建用户成功", c)
	}
}

// FindList 查询用户列表
func (a *UserApi) FindList(c *gin.Context) {
	var req request.PageInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMessage("参数校验失败", err, c)
		return
	}
	if list, total, err := userService.SelectList(&req); err != nil {
		response.ErrorWithMessage("查询用户列表失败", err, c)
	} else {
		response.Ok(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "查询用户列表成功", c)
	}
}

// Delete 删除用户
func (a *UserApi) Delete(c *gin.Context) {
	var req request.UUID
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMessage("参数校验失败", err, c)
		return
	}
	if err := userService.Delete(req.UUID); err != nil {
		response.ErrorWithMessage("删除用户失败", err, c)
	} else {
		response.OkWithMessage("删除用户成功", c)
	}
}

// ResetPassword 重置密码
func (a *UserApi) ResetPassword(c *gin.Context) {
	var req request.UUID
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMessage("参数校验失败", err, c)
		return
	}
	if err := userService.ResetPassword(req.UUID); err != nil {
		response.ErrorWithMessage("重置密码失败", err, c)
	} else {
		response.OkWithMessage("重置密码成功", c)
	}
}

// ChangePassword 修改密码
func (a *UserApi) ChangePassword(c *gin.Context) {
	var req request.ChangePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMessage("参数校验失败", err, c)
		return
	}
	uuid := utils.GetUserUuid(c)
	if err := userService.ChangePassword(uuid, req.Password); err != nil {
		response.ErrorWithMessage("修改密码失败", err, c)
	} else {
		response.OkWithMessage("修改密码成功", c)
	}
}
