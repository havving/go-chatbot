package apiserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-chatbot/internal/app/apiserver/config"
	"go-chatbot/internal/models"
	"net/http"
)

// Start starts the echo HTTP server
func Start() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     config.Server.CORS.AllowOrigins,
		AllowMethods:     config.Server.CORS.AllowMethods,
		AllowHeaders:     config.Server.CORS.AllowHeaders,
		AllowCredentials: config.Server.CORS.AllowCredentials,
	}))

	e.POST("/api/hello", Hello)

	e.Start(":8080")
}

func Hello(ctx echo.Context) error {
	fmt.Println("야호")
	var req kakao.Request
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	params := req.UserRequest.Utterance
	msg := fmt.Sprintf("아하~ %s님 반가워요. 앞으로 저랑 대화하려면, 아래의 동그란 버튼을 눌러주시면 돼요.", params)

	var quickReplies kakao.Kakao
	quickReplies.Add(kakao.QuickReply{}.New("동그라미", "동그라미"))

	return ctx.JSON(http.StatusOK, kakao.SimpleText{}.Build(msg, quickReplies))
}
