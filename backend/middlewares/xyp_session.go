package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/pengye91/xieyuanpeng.in/backend/configs"
)

var Store, _ = sessions.NewRedisStore(10, "tcp", configs.REDIS_URL, "", []byte("secret"))

func MyStore(s sessions.Store) sessions.Store {
	s.Options(
		sessions.Options{
			Path:     "/",
			Domain:   configs.BASE_DOMAIN,
			MaxAge:   86400,
			Secure:   false,
			HttpOnly: false,
		})
	return s
}

var Session_middleware = sessions.Sessions("sessionid", MyStore(Store))
