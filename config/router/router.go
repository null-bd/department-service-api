package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/null-bd/authn"
	"github.com/null-bd/authn/pkg/auth"

	"github.com/null-bd/department-service-api/config"
	"github.com/null-bd/department-service-api/internal/rest"
)

type Router struct {
	engine         *gin.Engine
	authMiddleware *authn.AuthMiddleware
	config         *config.Config
}

func NewRouter(cfg *config.Config, h *rest.Handler) (*Router, error) {
	// Load auth config
	authConfig := loadAuthConfig(cfg)

	// Initialize permission callback
	permCallback := func(orgId, branchId, role string) []string {
		// Customize this based on your needs
		return nil
	}

	// Create auth middleware
	authMiddleware, err := authn.NewAuthMiddleware(*authConfig, permCallback)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize auth middleware: %v", err)
	}

	// Create resource matcher
	resourceMatcher := authn.NewResourceMatcher(authConfig.Resources)

	// Set Gin mode
	gin.SetMode(getGinMode(cfg.App.Env))

	// Initialize router
	router := gin.New()

	// Add default middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// Add authentication middleware
	router.Use(authMiddleware.Authenticate())

	// Setup routes
	setupHealthRoutes(router, h)
	setupAPIRoutes(router, h, resourceMatcher)

	return &Router{
		engine:         router,
		authMiddleware: authMiddleware,
		config:         cfg,
	}, nil
}

func loadAuthConfig(cfg *config.Config) *authn.ServiceConfig {
	resources := make([]auth.ResourcePermission, 0, len(cfg.Auth.Resources))
	for _, resource := range cfg.Auth.Resources {
		resources = append(resources, auth.ResourcePermission{
			Path:    resource.Path,
			Method:  resource.Method,
			Actions: resource.Actions,
			Roles:   resource.Roles,
		})
	}

	publicPaths := make([]authn.PublicPath, 0, len(cfg.Auth.PublicPaths))
	for _, publicPath := range cfg.Auth.PublicPaths {
		publicPaths = append(publicPaths, authn.PublicPath{
			Path:   publicPath.Path,
			Method: publicPath.Method,
		})
	}

	authConfig := &authn.ServiceConfig{
		ServiceID:    cfg.Auth.ServiceID,
		ClientID:     cfg.Auth.ClientID,
		KeycloakURL:  cfg.Auth.KeycloakURL,
		Realm:        cfg.Auth.Realm,
		CacheEnabled: cfg.Auth.CacheEnabled,
		CacheURL:     cfg.Auth.CacheURL,
		Resources:    resources,
		PublicPaths:  publicPaths,
	}
	return authConfig
}

func (r *Router) Run() error {
	return r.engine.Run(r.config.App.GetAddress())
}

func setupAPIRoutes(router *gin.Engine, h *rest.Handler, resourceMatcher *authn.ResourceMatcher) {
	// v1 := router.Group("/api/v1")
	{
		// Example Resources routes with authl
		// TODO: Update the routes for service
		// resources := v1.Group("/resources")
		{
			// resources.GET("", h.GetResources)
		}
	}
}

func corsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func setupHealthRoutes(router *gin.Engine, h *rest.Handler) {
	router.GET("/health", h.HealthCheck)
}

func getGinMode(env string) string {
	switch env {
	case "production":
		return gin.ReleaseMode
	case "testing":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}
