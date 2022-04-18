package config

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// Config provides db credentials and app mode
type Config struct {
	DB          DBCreds      `json:"db"`
	MongoDB     MongoDBCreds `json:"mongo_db"`
	ServiceName string       `json:"service_name"`
	Log         ConfLog      `json:"log"`
}

// DbCreds describes db credentianls
type DBCreds struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"db_name"`
	LogMode  bool   `json:"log_mode"`
}

// MongoDbCreds creds for mongo db setup
type MongoDBCreds struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"db_name"`
	PoolSize uint64 `json:"pool_size"`
}

// ConfLog contains logger configuration details
type ConfLog struct {
	Level string `json:"level"`
}

// Load loads config in json format form file provided as configPath
func Load(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	config := &Config{}

	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	config.DB.DBName = getEnv("DB_NAME_PSQL", "")
	config.DB.Host = getEnv("HOST_PSQL", "")
	config.DB.LogMode = getEnvAsBool("DB_NAME", true)
	config.DB.Password = getEnv("PASSWORD_PSQL", "")
	config.DB.User = getEnv("USER_PSQL", "")
	config.DB.Port = getEnvAsInt("PORT_PSQL", 5432)

	return config, nil
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
