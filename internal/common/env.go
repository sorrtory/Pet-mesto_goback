package common

import (
	"log"
	"os"
)

func GetEnv(env string) string {
	v := os.Getenv(env)
	if v == "" {
		log.Printf("%v isn't set!\n", env)
	}
	return v
}
