package authentication

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/theghostmac/bankie.go/api/rest"
	"net/http"
	"os"
)

// WithJWTAuth is a middleware function that provides JWT authentication for HTTP handlers.
func WithJWTAuth(authHandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Calling JWT Auth Middleware...")
		tokenString := request.Header.Get("x-jwt-token")
		_, err := ValidateJWT(tokenString)
		if err != nil {
			rest.WriteJSON(writer, http.StatusForbidden, rest.APIError{Error: "Invalid Token"})
		}
		authHandlerFunc(writer, request)
	}
}

// ValidateJWT validates JWT tokens.
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})
}
