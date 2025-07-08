package app

import (
	"log"

	"github.com/dangLuan01/rebuild-api-movie28/internal/config"
	"github.com/dangLuan01/rebuild-api-movie28/internal/middleware"
	"github.com/dangLuan01/rebuild-api-movie28/internal/repository/redis"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	"github.com/dangLuan01/rebuild-api-movie28/internal/validation"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/elastic/go-elasticsearch/v7"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config *config.Config
	router *gin.Engine
}

func NewApplication(cfg *config.Config, DB *goqu.Database, ES *elasticsearch.Client) *Application {

	if err := validation.InitValidator(); err != nil {
		log.Fatalf("Validation init failed %v:", err)
	}
	redisRepo := redis.NewRedisRepository(cfg.Redis)
	
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	modules := []Module{
		NewUserModule(DB, redisRepo),
		NewGenreModule(DB, redisRepo),
		NewMovieModule(DB, redisRepo),
		NewCategoryModule(DB),
		NewThemeModule(DB, redisRepo),
		NewSearchModule(ES),
	}

	routes.RegisterRoute(r, getModuleRoutes(modules)...)
	return &Application{
		config: cfg,
		router: r,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}
func LoadEnv()  {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}