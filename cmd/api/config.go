package main

import (
	"flag"
	"io/ioutil"
	"os"
	"sync"

	"github.com/romarq/visualtez-storage/internal/logger"
	"gopkg.in/yaml.v2"
)

// Config holds API configurations
type Config struct {
	Port string         `yaml:"port,omitempty"`
	DB   DatabaseConfig `yaml:"database,omitempty"`
	Log  LogConfig      `yaml:"log,omitempty"`
}

// DatabaseConfig holds database configurations
type DatabaseConfig struct {
	URL string `yaml:"url,omitempty"`
}

// LogConfig holds logging configuration
type LogConfig struct {
	Location string `yaml:"location,omitempty"`
	Level    string `yaml:"level,omitempty"`
}

// EnvironmentProperty - Known environment properties
type EnvironmentProperty string

const (
	postgresURL EnvironmentProperty = "POSTGRES_URL"
	logLocation EnvironmentProperty = "LOG_LOCATION"
	apiPort     EnvironmentProperty = "API_PORT"
)

var once sync.Once
var singleton Config

// GetConfig - Get Configurations (Singleton pattern)
func GetConfig() Config {
	// Load environment variables only once
	once.Do(func() {
		var configPath string
		flag.StringVar(&configPath, "config", "./config/api.yaml", "API config file location")
		flag.Parse()
		singleton = load(configPath)
	})

	return singleton
}

// Load configuration from yaml and environment variables
func load(file string) Config {
	logger.Info("Loading configurations from: %s", file)

	// Config instance
	c := Config{}

	// Load config from YAML file
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Warn("Error reading configuration file: %s. %v", file, err)
	}
	if err := yaml.Unmarshal(fileContents, &c); err != nil {
		logger.Warn("Failed to parse configuration file: %s", file, err)
	}

	// Override configurations by ENV values (if provided)

	postgresURLFromEnv := os.Getenv(string(postgresURL))
	if postgresURLFromEnv != "" {
		c.DB.URL = postgresURLFromEnv
	}

	logLocationFromEnv := os.Getenv(string(logLocation))
	if logLocationFromEnv != "" {
		c.Log.Location = logLocationFromEnv
	}

	apiPortFromEnv := os.Getenv(string(apiPort))
	if apiPortFromEnv != "" {
		c.Port = apiPortFromEnv
	}

	if c.DB.URL == "" || c.Log.Location == "" || c.Port == "" {
		logger.Fatal("Failed to load configuration, missing mandatory configs: %v", c)
	}

	return c
}
