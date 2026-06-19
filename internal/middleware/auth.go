// Package middleware contains Gin middleware, including OIDC/JWT
// authentication against the existing Keycloak setup (see keycloak/ folder).
package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"

	"swe-workshop-api/internal/config"
)

// RequireAuth builds a Gin middleware that verifies a Bearer JWT issued by
// the Keycloak realm configured via KEYCLOAK_ISSUER_URL / KEYCLOAK_CLIENT_ID.
//
// The OIDC provider/verifier is resolved once, at server startup, so a
// reachability problem with Keycloak fails fast instead of on every request.
// This is a placeholder for later workshop steps: it does not yet inspect
// claims beyond issuer/audience checks performed by go-oidc itself.
func RequireAuth(ctx context.Context, cfg config.KeycloakConfig) (gin.HandlerFunc, error) {
	provider, err := oidc.NewProvider(ctx, cfg.IssuerURL)
	if err != nil {
		return nil, fmt.Errorf("connecting to Keycloak OIDC provider at %s: %w", cfg.IssuerURL, err)
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: cfg.ClientID})

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			return
		}
		rawToken := strings.TrimPrefix(authHeader, bearerPrefix)

		if _, err := verifier.Verify(c.Request.Context(), rawToken); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Next()
	}, nil
}
