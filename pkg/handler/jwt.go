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

type user struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
	Role  string `form:"role" json:"role" binding:"required"`
}

var phone = "phone"
var role = "role"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(phone)
	c.JSON(200, gin.H{
		"userID": claims[phone],
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
		IdentityKey: phone,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			log.Println("PayloadFunc")
			if v, ok := data.(*user); ok {
				return jwt.MapClaims{
					phone: v.Phone,
					role:  v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			log.Println("IdentityHandler", claims[phone].(string), claims[role].(string))
			return &user{
				Phone: claims[phone].(string),
				Role:  claims[role].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			log.Println("Authenticator")
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			phone := loginVals.Phone
			password := loginVals.Password
			role := loginVals.Role
			u, err := h.storage.User.GetByPhone(phone)
			if err != nil {
				return nil, err
			}
			//заглушка
			u.Phone = "996550269716"
			u.Password = "123456"
			if (phone == u.Phone) && (password == u.Password) {
				u := &user{
					Phone: phone,
					Role:  role,
				}
				return u, nil
			}
			//заглушка
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			log.Println("Authorizator")
			log.Println(data)
			if v, ok := data.(*user); ok && (v.Role == client || v.Role == psychologist || v.Role == moderator) {
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
