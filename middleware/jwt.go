package middleware

import (
	"context"
	"net/http"
	"strings"

	"httpserver/utils"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &utils.Claims{}

		token, err := jwt.ParseWithClaims(
 			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				return utils.JwtSecret, nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "email", claims.Email)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}