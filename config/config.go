package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warnln(".env not found, falling back to OS variables", err)
	}
}
