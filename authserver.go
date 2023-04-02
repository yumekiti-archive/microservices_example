package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/callback", func(c echo.Context) error {
		// JWT トークンを生成する
		token := jwt.New(jwt.SigningMethodHS256)

		// トークンにペイロードを追加する
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "John Doe"
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

		// トークンに署名を追加する
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		// トークンをレスポンスとして返す
		return c.JSON(http.StatusOK, map[string]string{
			"token": tokenString,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
