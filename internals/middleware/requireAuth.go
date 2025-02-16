package middleware

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shevchenko-nikita/EDiary/internals/repository"
	"net/http"
	"os"
	"time"
)

func RequireAuth(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, err := c.Cookie("Authorization")

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			user, err := repository.FindUserByUsername(db, claims["username"].(string))

			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("user", user)

			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
