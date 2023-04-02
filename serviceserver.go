// serviceserver.go
package main

import (
    "fmt"
    "net/http"

    "github.com/golang-jwt/jwt"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    // Echoのインスタンスを作成する
    e := echo.New()

    // 認証機能のミドルウェアを作成する
    authMiddleware := middleware.JWT([]byte("secret"))

    // "hello" を返すエンドポイントを作成する
    e.GET("/hello", func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        name := claims["name"].(string)
        return c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
    }, authMiddleware)

    // サーバーを開始する
    e.Logger.Fatal(e.Start(":8081"))
}