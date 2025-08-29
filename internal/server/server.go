package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/juanjerrah/go-project/internal/config"
	"github.com/juanjerrah/go-project/internal/models"
	"github.com/juanjerrah/go-project/pkg/database"
)

type Server struct {
	cfg *config.Config
	router *gin.Engine
	db *database.Database
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg:   cfg,
		router: gin.Default(),
	}
}

func (s *Server) InitializeDatabase() error {
	db, err := database.NewDatabase(s.cfg.DbHost, s.cfg.DbPort, s.cfg.DbUser, s.cfg.DbPassword, s.cfg.DbName)

	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	s.db = db

	
	if err := s.db.Migrate(&models.User{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

func (s *Server) initRoutes() {
	// Inicializar repositórios e serviços
	// userRepo := repositories.NewUserRepository(s.db.DB)
	// userService := services.NewUserService(userRepo)
	// userHandler := handlers.NewUserHandler(userService)
	
	// Rotas da API
	api := s.router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			// users.POST("", userHandler.CreateUser)
			// users.GET("", userHandler.GetAllUsers)
			// users.GET("/:id", userHandler.GetUser)
			// users.PUT("/:id", userHandler.UpdateUser)
			// users.DELETE("/:id", userHandler.DeleteUser)
		}
	}
	
	// Health check
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
}

func (s *Server) Run() error {
	if err := s.InitializeDatabase(); err != nil {
		return err
	}
	defer s.db.Close()

	s.initRoutes()

	return s.router.Run(":" + s.cfg.ServerPort)
}
