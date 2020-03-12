package echo

import (
	"fmt"
	"github.com/labstack/echo"
	token "github.com/luckydine/session-store"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Get your custom fields from the context c.Get("custom_token") returns as a map
func ValidateSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isExpired, to, err := token.GetLdtTokenFromRequest(c.Request())

		if isExpired || err != nil {
			return c.JSON(401, "Unathorized")
		}

		c.Set("custom_token", to)
		return next(c)
	}
}

// Get your custom fields from the context c.Get("custom_token") returns as a map
func ValidateSessionOrRedirectToLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isExpired, to, err := token.GetLdtTokenFromRequest(c.Request())

		if isExpired || err != nil {
			loginTarget := fmt.Sprintf("%s?forward=%s", strings.Trim(os.Getenv("LOGIN_REDIRECT"), url.QueryEscape(c.Request().URL.String())))
			c.Redirect(302, loginTarget)
			return nil
		}

		c.Set("custom_token", to)
		return next(c)
	}
}

func RenewSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		to := c.Get("custom_token").(token.Token)
		isExpired, tokenString, err := to.Renew()

		if isExpired || err != nil {
			loginTarget := filepath.Join(os.Getenv("LOGIN_REDIRECT"), fmt.Sprintf("?forward=%s", url.QueryEscape(c.Request().URL.String())))
			c.Redirect(302, loginTarget)
			return nil
		}

		sessionCookie := &http.Cookie{
			Name:     os.Getenv("SESSION_COOKIE"),
			Value:    tokenString,
			Path:     "",
			Domain:   c.Request().Host,
			Expires:  time.Now().Add(15 * time.Minute),
			Secure:   true,
			HttpOnly: false,
		}
		c.SetCookie(sessionCookie)
		return next(c)
	}
}
