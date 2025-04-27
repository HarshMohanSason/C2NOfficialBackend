package utils

import (
	"c2nofficialsitebackend/config"
	"errors"
	"net/http"
)

type Cookie struct {
	Name  string
	Value string
	Path  string
}

func SetAuthCookies(response http.ResponseWriter, cookies ...*Cookie) {
	switch config.Env.EnvType {
	case "PROD":
		for _, c := range cookies {
			http.SetCookie(response, &http.Cookie{
				Name:     c.Name,
				Value:    c.Value,
				HttpOnly: true,
				Secure:   true,
				Path:     "/",
				SameSite: http.SameSiteNoneMode,
			})
		}
	case "DEV":
		for _, c := range cookies {
			http.SetCookie(response, &http.Cookie{
				Name:     c.Name,
				Value:    c.Value,
				HttpOnly: false,
				Secure:   false,
				Path:     "/",
				SameSite: http.SameSiteLaxMode,
			})
		}
	default:
		for _, c := range cookies {
			http.SetCookie(response, &http.Cookie{
				Name:     c.Name,
				Value:    c.Value,
				HttpOnly: false,
				Secure:   false,
				Path:     "/",
				SameSite: http.SameSiteLaxMode,
			})
		}
		config.LogError(errors.New("warning! unknown error occured. defaulting to dev settings"))
	}
}
