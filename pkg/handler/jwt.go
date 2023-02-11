package handler

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/serjbibox/friend-service/pkg/models"
)

type login struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}

var identityKey = "phone"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID": claims[identityKey],
		"Phone":  user.(*models.User).Phone,
		"text":   "Hello World.",
	})
}

func (h *Handler) newAuthMiddleWare() (*jwt.GinJWTMiddleware, error) {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			log.Println("PayloadFunc")
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			log.Println("IdentityHandler")
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Role: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			log.Println("Authenticator")
			u, err := h.signIn(c)
			return u, err
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			log.Println("Authorizator")
			log.Println(data)
			if v, ok := data.(*models.User); ok && v.Role == client {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			log.Println("Unauthorized")
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}
