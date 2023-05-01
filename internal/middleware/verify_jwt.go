package middleware

import (
	"deall-alfon/internal"
	"deall-alfon/internal/constant"
	"deall-alfon/internal/usecase"
	"deall-alfon/pkg/api"
	"deall-alfon/pkg/config"
	"deall-alfon/pkg/util/converter"
	"deall-alfon/pkg/util/fn"
	"deall-alfon/pkg/util/logger"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	cfg      config.ConfigStr
	usecases usecase.UseCases
}

func NewMiddleware(cfg config.ConfigStr, usecases usecase.UseCases) internal.MiddlewareItf {
	return &Middleware{
		cfg:      cfg,
		usecases: usecases,
	}
}

func (m *Middleware) VerifyJWTAccess() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			op := fn.Name()
			req := c.Request()
			header := req.Header

			accessToken := header[constant.HeaderJWTAccessToken]
			var realAccessToken string
			if len(accessToken) <= 0 {
				response := api.DefaultResponse(c)
				return response.ResponseBadRequestError(constant.AccessTokenNotFound)
			}
			realAccessToken = converter.ToString(accessToken[0])

			var jwtclaim jwt.MapClaims
			jwttoken, err := jwt.Parse(realAccessToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("invalid jwt token: %v", token.Header["alg"])
				}

				return []byte(m.cfg.JWT.SecretKey), nil
			})

			jwtclaim = jwttoken.Claims.(jwt.MapClaims)
			email := converter.ToString(jwtclaim[constant.KeyTokenEmail])

			if err != nil {
				logger.Printf("[%v] error parse token for email: %v, error: %v", op, email, err)
				response := api.DefaultResponse(c)
				return response.ResponseUnauthorized(constant.AccessTokenInvalidError)
			}

			userLevel := converter.ToInt64(jwtclaim[constant.KeyTokenUserLevel])
			if userLevel != constant.UserLevelAdmin && userLevel != constant.UserLevelRoot {
				logger.Printf("[%v] user dont have access email: %v, error: %v", op, email, err)
				response := api.DefaultResponse(c)
				return response.ResponseUnauthorized(constant.AccessTokenUnauthorized)
			}

			return next(c)
		}
	}
}
