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


func SetEnv(env string, value string)  {
    os.Setenv(env, value)
}
