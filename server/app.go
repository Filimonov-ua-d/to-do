package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Filimonov-ua-d/to-do/pkg"
	dhttp "github.com/Filimonov-ua-d/to-do/pkg/delivery/http"
	"github.com/Filimonov-ua-d/to-do/pkg/repository/postgres"
	pkgUC "github.com/Filimonov-ua-d/to-do/pkg/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server
	pkgUC      pkg.UseCase
}

func NewApp() *App {

	var db *sqlx.DB
	var err error

	user := viper.GetString("POSTGRES_USER")
	password := viper.GetString("POSTGRES_PASSWORD")
	dbname := viper.GetString("POSTGRES_DB")
	// sslmode := viper.GetString("SSLMODE")
	host := viper.GetString("POSTGRES_HOST")
	port := viper.GetString("POSTGRES_PORT")

	// dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
	// user, password, dbname, sslmode, host, port)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s",
		user, password, dbname, host, port)

	if db, err = sqlx.Connect("postgres", dsn); err != nil {
		log.Panic(err)
	}

	pkgRepo := postgres.NewPkgRepository(db)

	return &App{
		pkgUC: pkgUC.NewPkgUseCase(
			pkgRepo,
			[]byte(viper.GetString("SIGNING_KEY")),
			viper.GetString("HASH_SALT"),
			viper.GetDuration("TOKEN_TTL"),
			viper.GetString("PORT")),
	}
}

func (a *App) Run(port string) error {

	router := gin.Default()
	router.Use(CORSMiddleware())

	dhttp.RegisterHTTPEndpoints(router, a.pkgUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("running server error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the required CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, restrict in production
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// If the request method is OPTIONS, respond with 200 OK
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Process the next handler
		c.Next()
	}
}
