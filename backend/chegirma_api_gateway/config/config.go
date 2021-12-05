package config

import (
	"log"

	"github.com/MrWebUzb/goenv"
)

type Config struct {
	Environment string `env:"ENVIRONMENT" default:"develop"`

	SettingServiceHost string `env:"SETTING_SERVICE_HOST" default:"localhost"`
	SettingServicePort int    `env:"SETTING_SERVICE_PORT" default:"8003"`

	RedisHost             string `env:"REDIS_HOST" default:"localhost"`
	RedisPassword         string `env:"REDIS_PASSWORD" default:"redis"`
	RedisPort             int    `env:"REDIS_PORT" default:"6379"`
	RedisTimeoutInSeconds int    `env:"REDIS_TIMEOUT_IN_SECONDS" default:"120"`

	LogLevel            string `env:"LOG_LEVEL" default:"debug"`
	HttpPort            string `env:"HTTP_PORT" default:":8000"`
	BucketName          string `env:"MINIO_BUCKET_NAME" default:"chegirma"`
	MinioDomain         string `env:"MINIO_DOMAIN" default:"cdn.chegirma.udevs.io"`
	MinioAccessKeyID    string `env:"MINIO_ACCESS_KEY" default:"8DbGdJfNjQmSqVsXv2x4z7C9EbHeKgNkRnTrWtYv3y5A7DaFcJfMhPmSpUrXuZw3z6B8EbGdJgNjQmTqVsXv2x4A7C"`
	MinioSecretAccesKey string `env:"MINIO_SECRET_KEY" default:"QmSpUsXuZx4z6B9EbGdKgNjQnTqVtYv2x5A7C9FcHeKhPkRpUrWtZw3y5B8DaFdJfMjQmSpVsXuZx4z6B9EbGeKgNj"`

	LoginSecretAccessKey  string `env:"LOGIN_ACCESS_SECRET_KEY" default:"dWRldnMgZGV2ZWxvcGVkIGVsZWN0cm9uIGthZGFzdHIK"`
	LoginSecretRefreshKey string `env:"LOGIN_REFRESH_SECRET_KEY" default:"ZWxlY3Ryb24ga2FkYXN0ciAgLT4gdWRldnMgZGV2ZWxvcGVkCg"`

	KafkaUrl string `env:"KAFKA_URL" default:"localhost:9092"`

	ReportService     string `env:"REPORT_SERVICE_HOST" default:"localhost"`
	ReportServicePort int    `env:"REPORT_SERVICE_PORT" default:"8007"`
}

func Load(fileNames ...string) Config {
	config := Config{}

	env, err := goenv.New(fileNames...)

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	if err := env.Parse(&config, nil); err != nil {
		panic(err)
	}

	return config
}
