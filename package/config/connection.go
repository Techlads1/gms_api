package config

// func GetConnectionString() string {
// 	return "host=localhost dbname=teleradiology user=postgres password=password port=5432"

import (
	"github.com/tzdit/sample_api/package/log"
)

func GetDatabaseConnection() string {
	cfg, err := New()
	if err != nil {
		log.Errorf("error loading configuration file: %v", err)
		return ""
	}
	return cfg.GetDatabaseConnection()
}
