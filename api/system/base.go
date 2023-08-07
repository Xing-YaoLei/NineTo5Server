package system

import (
	"NineTo5Server/global"
	"NineTo5Server/model/common/response"
	systemReq "NineTo5Server/model/system/request"
	systemRes "NineTo5Server/model/system/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router    /base/captcha [post]
func (ba *BaseApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.NineTo5_CONFIG.Captcha.ImgHeight, global.NineTo5_CONFIG.Captcha.ImgWidth, global.NineTo5_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global.NineTo5_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.NineTo5_CONFIG.Captcha.KeyLong,
	}, "验证码获取成功", c)
}

// Login
// @Tags      Base
// @Summary   用户登录
// @Produce   application/json
// @Success   200  {object}  responseResponse{data=string,msg=string}  "登录成功"
// @Router    /base/login [post]
func (ba *BaseApi) Login(c *gin.Context) {
	// TODO
	var login systemReq.Login
	_ = c.ShouldBindJSON(&login)
	validate := validator.New()
	// 校验数据
	if err := validate.Struct(&login); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.NineTo5_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if store.Verify(login.CaptchaId, login.Captcha, true) {
		// TODO
		response.OkWithMessage("登录成功", c)
	} else {
		response.FailWithMessage("验证码错误", c)
		return
	}

}
