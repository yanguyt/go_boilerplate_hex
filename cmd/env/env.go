package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//AllEnvs responsible to manage all .env file
type AllEnvs struct {
	Port string
	Mode string
}

// ReadAllEnvs is responsible to read and fill the struct with all Envs
func ReadAllEnvs() *AllEnvs {
	envs := &AllEnvs{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading . env File")
	}

	envs.Port = os.Getenv("PORT")
	envs.Mode = os.Getenv("MODE")
	if envs.Mode == "" {
		envs.Mode = "dev"
	}

	return envs
}
