package appctx

import (
	"database/sql"
	"log"
	"voca/internal/repositories"
	"voca/internal/router"
	"voca/pkg/util"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (config util.Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func Run() {

	var s util.Config
	cfg, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(s.DBDriver, s.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := repositories.NewStore(conn)
	server, err := router.NewServer(cfg, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(s.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
