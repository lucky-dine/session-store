package main

import (
	"github.com/labstack/echo"
	token "github.com/luckydine/session-store"
	echo2 "github.com/luckydine/session-store/middleware/echo"
	"net/http"
	"os"
	"time"
)

func main() {
	os.Setenv("SESSION_COOKIE", "lucky_session")
	os.Setenv("LOGIN_REDIRECT", "/health")
	os.Setenv("LDT_SECRET", `ZAJ6PpUtEySe^=Yad-nhH-3%9xh6g78E#?xhZx?M7z5e*qqUA57!9g28rKn+K6gxH+P$5UAh#Zs%PZ?!%h=RPjGU6**at&UYdcFAV9j#N3!NJwh&j!!Au6aQ*Ag*6u8xkh+ehUBQ#6QTa7u!+tpvAXgk7xp%DDaa6jV-BXE&e^3U7y$PzAc*smQv#CXZ=KpyekwYmx2rU2qDh^DKsukzcJuLMq!EX*m-XmNeL!WwDnkpSs$+jt?K?ja-Nfxpz7a@`)

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, "Healthy")
	})

	e.GET("/test", func(c echo.Context) error {
		return c.JSON(200, c.Get("custom_token"))
	}, echo2.ValidateSessionOrRedirectToLogin, echo2.RenewSession)

	e.GET("/new_token", func(c echo.Context) error {
		ts := token.New(map[string]string{"secret_message": "Pickles are great!"})

		sessionCookie := &http.Cookie{
			Name:     os.Getenv("SESSION_COOKIE"),
			Value:    ts,
			Path:     "",
			Domain:   c.Request().Host,
			Expires:  time.Now().Add(15 * time.Minute),
			Secure:   true,
			HttpOnly: false,
		}
		c.SetCookie(sessionCookie)

		return c.JSON(200, "Should have a shiny brand new token.")
	})

	e.Start(":5000")
}
