package platform

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ctxKey string

// UserIDKey carries the authenticated user id through the request context.
const UserIDKey ctxKey = "userID"

// AuthMiddleware reads X-User-ID and stores it on the request context.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.GetHeader("X-User-ID"))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(c.Request.Context(), UserIDKey, id)
		c.Request = c.Request.WithContext(ctx)
		c.Set("userID", id)
		c.Next()
	}
}

// UserID returns the authenticated user id from ctx.
func UserID(ctx context.Context) int {
	id, _ := ctx.Value(UserIDKey).(int)
	return id
}
