package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/config"
	"github.com/tapiaw38/alutec-inventory-be/internal/platform/token"
)

// RequireAuth rejects requests without a valid session token. Every failure
// answers the same way: distinguishing "expired" from "forged" would tell a
// caller probing the API which tokens are worth guessing at.
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := config.GetConfigService().AuthConfig.Secret
		if secret == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    "common:internal-server-error",
				"message": "auth is not configured",
			})
			return
		}

		raw, found := strings.CutPrefix(c.GetHeader("Authorization"), "Bearer ")
		if !found {
			unauthorized(c)
			return
		}

		claims, err := token.Verify(strings.TrimSpace(raw), secret)
		if err != nil {
			unauthorized(c)
			return
		}

		c.Set("auth_subject", claims.Subject)
		c.Next()
	}
}

func unauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":    "common:unauthorized",
		"message": "sesión inválida o expirada",
	})
}
