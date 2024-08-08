package configs

import (
	"log"
	"os"
	"reflect"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	ENV          string `env:"ENV"`
	PORT         string `env:"PORT"`
	DATABASE_URL string `env:"DATABASE_URL"`
}

var (
	envInstance *Env
	envOnce     sync.Once
)

// loadEnv loads environment variables from a .env file if it exists, and checks if all required variables are present.
func loadEnv() *Env {
	envOnce.Do(func() {
		// Attempt to load .env file
		if err := godotenv.Load(); err != nil {
			log.Println(".env file not found, using system environment variables")
		}

		env := &Env{}
		missingVars := []string{}

		// Use reflection to iterate over the fields of the Env struct
		v := reflect.ValueOf(env).Elem()
		typeOfEnv := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := typeOfEnv.Field(i)
			envTag := fieldType.Tag.Get("env")

			value, exists := os.LookupEnv(envTag)
			if !exists {
				missingVars = append(missingVars, envTag)
			} else {
				field.SetString(value)
			}
		}

		if len(missingVars) > 0 {
			log.Fatalf("Missing required environment variables: %v", missingVars)
		}

		envInstance = env
		log.Print("env loaded")
	})

	return envInstance
}

// GetEnv returns the singleton instance of Env.
func GetEnv() *Env {
	if envInstance == nil {
		return loadEnv()
	}
	return envInstance
}
