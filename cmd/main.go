// @title           Gin Tattoo API
// @version         1.0
// @description     API sobre estilos e curiosidades do mundo da tatuagem.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://github.com/santzin/gin-tattoo

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/santzin/gin-tattoo/docs"
	appdb "github.com/santzin/gin-tattoo/internal/db"
	"github.com/santzin/gin-tattoo/internal/handlers"
)

func main() {
	ctx := context.Background()

	pool, err := appdb.Connect(ctx)
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer pool.Close()

	if err := appdb.Migrate(ctx, pool); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	h := &handlers.H{DB: pool}

	r := gin.Default()

	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/", func(c *gin.Context) { c.File("./static/index.html") })

	r.GET("/health", h.HealthCheck)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/styles", h.ListStyles)
		v1.GET("/styles/:id", h.GetStyle)
		v1.GET("/curiosities", h.ListCuriosities)
		v1.GET("/curiosities/:id", h.GetCuriosity)
	}

	log.Println("listening on :8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server: %v", err)
	}
}
