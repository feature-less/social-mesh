package main

import (
	"log"
	"os"
	api "threads/internal/delivery/http"

	"server/pkg/router"
	"server/pkg/service"

	"github.com/casbin/casbin/v2"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Config struct {
	DbUrl   string `json:"url"`
	Port    string `json:"port"`
	Address string `json:"address"`
}

func main() {
	log.SetPrefix("[Threads Service]: ")

	var err error
	var httpRouter router.RouterConfig
	var worker service.Service

	enforcer, err := casbin.NewEnforcer("./authorization_model.conf", "./policy.csv")

	if err != nil {
		log.Fatalf(" unhandled casbin enforcer error\n enforcer error: %s", err)
	}

	viper.SetDefault("port", "9999")
	viper.SetDefault("address", "localhost")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("No config.yml was found %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode config into struct %s", err)
	}

	db, err := sqlx.Connect("pgx", viper.GetString("dbUrl"))
	if err != nil {
		log.Fatalf("Unable to connect to database, make sure your database is reachable or is running: %v", err)
	}
	defer db.Close()

	os.Setenv("DATABASE_URL", viper.GetString("dbUrl"))
	os.Setenv("PORT", viper.GetString("port"))
	os.Setenv("ADDRESS", viper.GetString("address"))

	log.Println("using database  ", config.DbUrl)
	log.Println("using port ", config.Port)

	httpRouter.Enforcer = enforcer
	httpRouter.Origin = os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")
	httpRouter.RootRoute = api.Root
	worker.HTTP.Addr = os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")
	worker.HTTP.Handler = httpRouter.Set()
	worker.Spawn()
}
