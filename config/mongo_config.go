package config

import "github.com/spf13/viper"

type MongoConfig struct {
	ConnectionString string
	Database string
}

func NewMongoConfig() MongoConfig {
	connectionString := viper.GetString("mongo_connection_string")
	database := viper.GetString("mongo_database")
	return MongoConfig { connectionString, database }
}