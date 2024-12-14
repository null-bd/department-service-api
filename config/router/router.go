package router

// import (
// 	"fmt"
// 	"time"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"

// 	// "github.com/null-bd/authmiddleware"

// 	"github.com/null-bd/microservice-name/config"
// 	"github.com/null-bd/microservice-name/internal/rest"
// )

// type Router struct {
// 	engine         *gin.Engine
// 	authMiddleware *authmiddleware.AuthMiddleware
// 	config         *config.Config
// }

// func NewRouter(cfg *config.Config, h *rest.Handler) (*Router, error) {
// 	// Initialize auth middleware
// 	authConfig, err := authmiddleware.NewConfigLoader("resources/config.yaml").Load()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load auth config: %v", err)
// 	}

// 	// Initialize permission callback
// 	permCallback := func(orgId, branchId, role string) []string {
// 		// Customize this based on your needs
// 		return nil
// 	}

// 	// Create auth middleware
// 	authMiddleware, err := authmiddleware.NewAuthMiddleware(*authConfig, permCallback)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to initialize auth middleware: %v", err)
// 	}

// 	// Create resource matcher
// 	resourceMatcher := authmiddleware.NewResourceMatcher(authConfig.Resources)

// 	// Set Gin mode
// 	gin.SetMode(getGinMode(cfg.App.Env))

// 	// Initialize router
// 	router := gin.New()

// 	// Add default middleware
// 	router.Use(gin.Logger())
// 	router.Use(gin.Recovery())
// 	router.Use(corsMiddleware())

// 	// Add authentication middleware
// 	router.Use(authMiddleware.Authenticate())

// 	// Setup routes
// 	setupHealthRoutes(router, h)
// 	setupAPIRoutes(router, h, resourceMatcher)

// 	return &Router{
// 		engine:         router,
// 		authMiddleware: authMiddleware,
// 		config:         cfg,
// 	}, nil
// }

// func (r *Router) Run() error {
// 	return r.engine.Run(r.config.App.GetAddress())
// }

// func setupAPIRoutes(router *gin.Engine, h *rest.Handler, resourceMatcher *authmiddleware.ResourceMatcher) {
// 	// v1 := router.Group("/api/v1")
// 	{
// 		// Example Resources routes with authl
// 		// TODO: Update the routes for service
// 		// resources := v1.Group("/resources")
// 		{
// 			// resources.GET("", h.GetResources)
// 		}
// 	}
// }

// func corsMiddleware() gin.HandlerFunc {
// 	return cors.New(cors.Config{
// 		AllowOrigins:     []string{"*"},
// 		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		MaxAge:           12 * time.Hour,
// 	})
// }

// func setupHealthRoutes(router *gin.Engine, h *rest.Handler) {
// 	router.GET("/health", h.HealthCheck)
// }

// func getGinMode(env string) string {
// 	switch env {
// 	case "production":
// 		return gin.ReleaseMode
// 	case "testing":
// 		return gin.TestMode
// 	default:
// 		return gin.DebugMode
// 	}
// }
