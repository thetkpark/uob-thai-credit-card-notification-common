package config

import (
	"github.com/Netflix/go-env"
	"log"
)

func LoadConfigFromENV(c interface{}) interface{} {
	if _, err := env.UnmarshalFromEnviron(c); err != nil {
		log.Fatal(err)
	}
	return c
}
