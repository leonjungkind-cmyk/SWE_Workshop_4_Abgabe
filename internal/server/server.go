// Package server wires up the Gin engine, route groups, and middleware.
package server

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"swe-workshop-api/internal/config"
	"swe-workshop-api/internal/middleware"
)

// NewRouter builds the Gin engine with public and secured route groups.
//
// /api/public  - reachable without authentication (e.g. health checks)
// /api/secured - requires a valid Keycloak-issued Bearer token
//
// Keycloak is optional for this workshop step: if the OIDC provider cannot
// be reached at startup, the secured group is still mounted but every
// request to it is rejected with 401, and the reason is logged.
func NewRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	public := router.Group("/api/public")
	registerPublicRoutes(public, db)

	secured := router.Group("/api/secured")
	authMiddleware, err := middleware.RequireAuth(context.Background(), cfg.Keycloak)
	if err != nil {
		log.Printf("warning: Keycloak auth disabled, secured routes will reject all requests: %v", err)
		authMiddleware = func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "auth not available"})
		}
	}
	secured.Use(authMiddleware)
	registerSecuredRoutes(secured, db)

	return router
}

func registerPublicRoutes(group *gin.RouterGroup, _ *gorm.DB) {
	group.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}

func registerSecuredRoutes(_ *gin.RouterGroup, _ *gorm.DB) {
	// Placeholder: protected handlers will be added in a later workshop step.
}
