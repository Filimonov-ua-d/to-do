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
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type App struct {
	httpServer *http.Server
	pkgUC      pkg.UseCase
}

func NewApp() *App {

	var db *sqlx.DB
	var err error

	loggerUC := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("Layer:", "usecase").
		Str("Service:", "Home_finances").
		Logger()

	loggerRepo := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("Layer:", "repository").
		Str("Service:", "Home_finances").
		Logger()

	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.password")
	dbname := viper.GetString("postgres.dbname")
	sslmode := viper.GetString("postgres.sslmode")
	host := viper.GetString("postgres.host")
	port := viper.GetString("postgres.port")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		user, password, dbname, sslmode, host, port)

	if db, err = sqlx.Connect("postgres", dsn); err != nil {
		log.Panic(err)
	}

	pkgRepo := postgres.NewPkgRepository(db, &loggerRepo)

	return &App{
		pkgUC: pkgUC.NewPkgUseCase(
			pkgRepo,
			&loggerUC,
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetString("auth.hash_salt"),
			viper.GetDuration("auth.token_ttl"),
			viper.GetString("port")),
	}
}

func (a *App) Run(port string) error {

	router := gin.Default()

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
