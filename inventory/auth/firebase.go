package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"firebase.google.com/go/auth"
)

const bearerSchema = "Bearer "

func AuthMiddleware(cl *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		token := authHeader[len(bearerSchema):]
		if token == "" {
			respondWithError(c, http.StatusUnauthorized, "API token required")
			return
		}
		ctx := c.Request.Context()
		t, err := cl.VerifyIDToken(ctx, token)
		if err != nil {
			respondWithError(c, http.StatusUnauthorized, "Invalid API token")
			return
		}
		user, err := cl.GetUser(ctx, t.UID)
		if err != nil {
			respondWithError(c, http.StatusUnauthorized, "User not found")
			return
		}
		c.Request.Header.Set("UserID", user.UID)
		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
