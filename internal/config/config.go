package config

import (
	"os"
	"time"

	_ "github.com/lib/pq" // postgres driver don`t delete
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	URL string `mapstructure:"url"`
}

type ServerConfig struct {
	Port     string `mapstructure:"port"`
	BasePath string `mapstructure:"base_path"`
}

type JWTConfig struct {
	AccessToken struct {
		SecretKey     string        `mapstructure:"secret_key"`
		TokenLifetime time.Duration `mapstructure:"token_lifetime"`
	} `mapstructure:"access_token"`
}

type MongoDBConfig struct {
	URL      string `mapstructure:"url"`
	database string `mapstructure:"database"`
}

type RabbitMQConfig struct {
	URL      string `mapstructure:"url"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Exchange string `mapstructure:"exchange"`
}

type CloudinaryConfig struct {
	CloudName string `mapstructure:"cloud_name"`
	APIKey    string `mapstructure:"api_key"`
	APISecret string `mapstructure:"api_secret"`
}

type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type UrlConfig struct {
	UserStorage struct {
		GetUser string `mapstructure:"get_user"`
	} `mapstructure:"user_storage"`
}

type SwaggerConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	URL     string `mapstructure:"url"`
	Port    string `mapstructure:"port"`
}

type CORSConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedMethods []string `mapstructure:"allowed_methods"`
	AllowedHeaders []string `mapstructure:"allowed_headers"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type RateLimitConfig struct {
	MaxRequests int           `mapstructure:"max_requests"`
	TimeWindow  time.Duration `mapstructure:"time_window"`
	Expiration  time.Duration `mapstructure:"expiration"`
}

type Config struct {
	Database DatabaseConfig   `mapstructure:"database"`
	MongoDB  MongoDBConfig    `mapstructure:"mongodb"`
	Server   ServerConfig     `mapstructure:"server"`
	JWT      JWTConfig        `mapstructure:"jwt"`
	Rabbit   RabbitMQConfig   `mapstructure:"rabbit"`
	Storage  CloudinaryConfig `mapstructure:"cloudinary"`
	Logging  LoggingConfig    `mapstructure:"logging"`
	Rate     RateLimitConfig  `mapstructure:"rate_limit"`
	Swagger  SwaggerConfig    `mapstructure:"swagger"`
	CORS     CORSConfig       `mapstructure:"cors"`
	Redis    RedisConfig      `mapstructure:"redis"`
	Url      UrlConfig        `mapstructure:"url"`
}

func LoadConfig() (*Config, error) {
	configPath := os.Getenv("KV_VIPER_FILE")
	if configPath == "" {
		return nil, errors.New("KV_VIPER_FILE env var is not set")
	}
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Errorf("error reading config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.Errorf("error unmarshalling config: %s", err)
	}

	return &config, nil
}
