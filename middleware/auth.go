package middleware

import (
	"context"
	"fmt"
	"go-backend/initializers"
	"go-backend/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get cookie off req
		tokenString, err := r.Cookie("Authorization")
		if err != nil {
			fmt.Println("Error getting cookie: ", err)
			http.Error(w, "Error Getting Cookie", http.StatusUnauthorized)
			return
		}

		// decode & validate it
		token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("TokenSecret")), nil
		})

		if err != nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		// check expiration
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Error Getting Claims", http.StatusUnauthorized)
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			http.Error(w, "Cookie Expired", http.StatusUnauthorized)
			return
		}

		// find user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			http.Error(w, "User Not Found", http.StatusUnauthorized)
			return
		}

		// attach to req
		ctx := context.WithValue(r.Context(), "user", &user)
		r = r.WithContext(ctx)

		// contiue
		next.ServeHTTP(w, r)
	})
}
