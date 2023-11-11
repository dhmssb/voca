package appctx

import (
	"voca/pkg/util"

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

// func Run() {

// 	cfg, err := LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("cannot load config: ", err)
// 	}

// 	conn, err := sql.Open(util.Config.DBDriver, Config.DBSource)
// 	if err != nil {
// 		log.Fatal("cannot connect to db: ", err)
// 	}

// 	store := repositories.NewStore(conn)
// 	// server, err :=

// }
