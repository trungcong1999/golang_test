package route

import (
	"connect/config"
	"connect/handler"
	"connect/middleware"
	"connect/repo"
	"connect/service"
)

type Service struct {
	*config.App
}

func NewService() *Service {
	s := Service{
		config.NewApp(),
	}

	db := s.GetDB()
	repo := repo.NewRepo(db)

	userService := service.NewUser(repo)
	user := handler.NewUser(userService)
	migrate := handler.NewMigration(db)
	health := handler.NewHealth()

	router := s.Router
	v1 := router.Group("/v1")

	// user
	v1.POST("/login", middleware.CheckAdmin(), user.GetUsers)
	router.POST("/check-health", health.CheckHealth)

	// migration
	router.POST("/migrate", middleware.CheckAdmin(), migrate.Migrate)
	return &s
}
