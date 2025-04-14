package utils

import (
	"net/http"
	"c2nofficialsitebackend/config"
	"errors"
)

type Cookie struct{
	Name string
	Value string
	Path string
}

func SetAuthCookies(response http.ResponseWriter, cookies ...*Cookie){
	switch config.Env.ENV_TYPE {
	case "PROD": 
		for _, c := range cookies{
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
    	for _, c := range cookies{
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
		config.LogError(errors.New("WARNING: Unknown ENV_TYPE. Defaulting to DEV settings."))
	}
}
