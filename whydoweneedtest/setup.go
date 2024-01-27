package whydoweneedtest

import (
	"github.com/LGROW101/LGROW-Microservices/config"
)

func NewTestConfig() *config.Config {
	cfg := config.LoadConfig("../env/test/.env")
	return &cfg
}
