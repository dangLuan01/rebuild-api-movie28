package app

import (
	"log"

	"github.com/dangLuan01/rebuild-api-movie28/internal/config"
	"github.com/dangLuan01/rebuild-api-movie28/internal/db"
	"github.com/dangLuan01/rebuild-api-movie28/internal/middleware"
	"github.com/dangLuan01/rebuild-api-movie28/internal/routes"
	"github.com/dangLuan01/rebuild-api-movie28/internal/validation"
	"github.com/doug-martin/goqu/v9"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config *config.Config
	router *gin.Engine
	//module []Module
}

type ModuleContext struct {
	DB *goqu.Database
	Redis *redis.Client
	ES *elasticsearch.Client
}

func NewApplication(cfg *config.Config) *Application {

	if err := validation.InitValidator(); err != nil {
		log.Fatalf("Validation init failed %v:", err)
	}

	if err := db.InitDB(); err != nil {
		log.Fatalf("unable to connect to sql")
	}

	redisClient := config.NewRedisClient()
	elasticsearchClient := config.NewElasticSearchClient()
	
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	ctx := &ModuleContext{
		DB: db.DB,
		Redis: redisClient,
		ES: elasticsearchClient,
	}


	modules := []Module{
		NewUserModule(ctx),
		NewGenreModule(ctx, ctx.Redis),
		NewMovieModule(ctx, ctx.Redis),
		NewCategoryModule(ctx),
		NewThemeModule(ctx, ctx.Redis),
		NewSearchModule(ctx.ES),
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}