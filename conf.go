package shoppe

import (
	_ "embed"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/shoppehub/shoppe/util"
)

//go:embed conf/.env
var defaultEnvString string

const (
	// 应用的端口
	PORT = "PORT"
)

func init() {

	envMap, _ := godotenv.Unmarshal(defaultEnvString)

	for key, value := range envMap {
		os.Setenv(key, strings.TrimSpace(value))
	}

	env := os.Getenv("SHOPPE_MODE")
	if "" == env {
		env = "development"
	}

	err := godotenv.Overload("conf/.env." + env)
	if err != nil {
		log.Fatal("Error loading .env."+env+" file", err)
	}

	log.Println("Shoppe Running in " + env)

	local := "conf/.env.local"

	if util.Exists(local) {
		godotenv.Overload(local)
		log.Println("Shoppe load " + local)
	}

}
