package ryserver

import (
	"RuoYi-Go/internal/response"
	"RuoYi-Go/pkg/captcha"
	"RuoYi-Go/pkg/redis"
	"fmt"
	"github.com/kataras/iris/v12"
	"strings"
	"time"
)

var server *iris.Application
var redisService *ryredis.RedisStruct

func InitServer(s *iris.Application, r *ryredis.RedisStruct) {
	server = s
	redisService = r
}

func StartServer() {
	// 定义路由
	server.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello, Iris!")
	})

	server.Get("/captchaImage", func(ctx iris.Context) {
		// 获取当前时间并格式化为HTTP日期格式
		currentTime := time.Now().UTC().Format(time.RFC1123)
		ctx.Header("Date", currentTime) // 设置Date头
		ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate")
		ctx.Header("Cache-Control", "post-check=0, pre-check=0")
		ctx.Header("Pragma", "no-cache")
		ctx.ContentType("image/jpeg")

		id, b64s, a, err := captcha.GenerateCaptcha()
		if err != nil {
			ctx.JSON(response.Error(500, "生成验证码失败"))
			return
		}

		redisService.Set(fmt.Sprintf("captcha:%d", id), a, time.Minute*5)

		user := response.CaptchaImage{
			Code:    200,
			Uuid:    id,
			Img:     b64s[strings.Index(b64s, ",")+1:],
			Message: "操作成功",
		}
		// 使用 ctx.JSON 自动将user序列化为JSON并写入响应体
		ctx.JSON(user)
	})

	// 定义路由
	server.Post("/login", func(ctx iris.Context) {

		var l loginStruct
		// Attempt to read and bind the JSON request body to the 'user' variable
		if err := ctx.ReadJSON(&l); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"message": "Invalid JSON", "error": err.Error()})
			return
		}

		v, error := redisService.Get(fmt.Sprintf("captcha:%d", l.Uuid))
		if error != nil || v == "" {
			ctx.JSON(response.Error(500, "验证码错误或已失效"))
			return
		}

		if v != "" && captcha.VerifyCaptcha(l.Uuid, l.Code) {
			ctx.WriteString("Hello, Iris!")
		} else {
			ctx.JSON(response.Error(500, "验证码错误"))
			return
		}

	})
}

type loginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
	Uuid     string `json:"uuid"`
}
