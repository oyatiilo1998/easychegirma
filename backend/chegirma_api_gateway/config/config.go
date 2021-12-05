package config

import (
	"log"

	"github.com/MrWebUzb/goenv"
)

type Config struct {
	Environment string `env:"ENVIRONMENT" default:"develop"`

	UserServiceHost string `env:"USER_SERVICE_HOST" default:"localhost"`
	UserServicePort int    `env:"USER_SERVICE_PORT" default:"8002"`

	SettingServiceHost string `env:"SETTING_SERVICE_HOST" default:"localhost"`
	SettingServicePort int    `env:"SETTING_SERVICE_PORT" default:"8003"`

	EntityServiceHost string `env:"ENTITY_SERVICE_HOST" default:"localhost"`
	EntityServicePort int    `env:"ENTITY_SERVICE_PORT" default:"8004"`

	IntegrationServiceHost string `env:"INTEGRATION_SERVICE_HOST" default:"localhost"`
	IntegrationServicePort int    `env:"INTEGRATION_SERVICE_PORT" default:"8009"`

	DiscussionLogicServiceHost string `env:"DISCUSSION_LOGIC_SERVICE_HOST" default:"localhost"`
	DiscussionLogicServicePort int    `env:"DISCUSSION_LOGIC_SERVICE_PORT" default:"8005"`

	ContentServiceHost string `env:"CONTENT_SERVICE_HOST" default:"localhost"`
	ContentServicePort int    `env:"CONTENT_SERVICE_PORT" default:"7004"`

	GeoServiceHost string `env:"GEO_SERVICE_HOST" default:"localhost"`
	GeoServicePort int    `env:"GEO_SERVICE_PORT" default:"8008"`

	RedisHost             string `env:"REDIS_HOST" default:"localhost"`
	RedisPassword         string `env:"REDIS_PASSWORD" default:"redis"`
	RedisPort             int    `env:"REDIS_PORT" default:"6379"`
	RedisTimeoutInSeconds int    `env:"REDIS_TIMEOUT_IN_SECONDS" default:"120"`

	LogLevel            string `env:"LOG_LEVEL" default:"debug"`
	HttpPort            string `env:"HTTP_PORT" default:":8000"`
	BucketName          string `env:"MINIO_BUCKET_NAME" default:"ekadastr"`
	MinioDomain         string `env:"MINIO_DOMAIN" default:"cdn.ekadastr.udevs.io"`
	MinioAccessKeyID    string `env:"MINIO_ACCESS_KEY" default:"8DbGdJfNjQmSqVsXv2x4z7C9EbHeKgNkRnTrWtYv3y5A7DaFcJfMhPmSpUrXuZw3z6B8EbGdJgNjQmTqVsXv2x4A7C"`
	MinioSecretAccesKey string `env:"MINIO_SECRET_KEY" default:"QmSpUsXuZx4z6B9EbGdKgNjQnTqVtYv2x5A7C9FcHeKhPkRpUrWtZw3y5B8DaFdJfMjQmSpVsXuZx4z6B9EbGeKgNj"`

	LoginSecretAccessKey  string `env:"LOGIN_ACCESS_SECRET_KEY" default:"dWRldnMgZGV2ZWxvcGVkIGVsZWN0cm9uIGthZGFzdHIK"`
	LoginSecretRefreshKey string `env:"LOGIN_REFRESH_SECRET_KEY" default:"ZWxlY3Ryb24ga2FkYXN0ciAgLT4gdWRldnMgZGV2ZWxvcGVkCg"`

	OneIdClientID      string `env:"ONE_ID_CLIENT_ID" default:"yerelektron"`
	OneIdClientSecret  string `env:"ONE_ID_CLIENT_SECRET" default:"o930RPG5hUHDK8fzkbvB4WLm"`
	OneIdRedirectURI   string `env:"ONE_ID_REDIRECT_URI" default:"https://test.yerelektron.uz/"`
	OneIdScope         string `env:"ONE_ID_SCOPE" default:"yerelektron"`
	OneIdUrl           string `env:"ONE_ID_URL" default:"https://sso2.egov.uz:8443/sso/oauth/Authorization.do"`
	AnalyticServiceUrl string `env:"ANALYTIC_SERVICE_URL" default:"127.0.0.1:8081"`
	KafkaUrl           string `env:"KAFKA_URL" default:"localhost:9092"`

	ReportService     string `env:"REPORT_SERVICE_HOST" default:"localhost"`
	ReportServicePort int    `env:"REPORT_SERVICE_PORT" default:"8007"`

	LandTaxPropertyId        string `env:"LAND_TAX_PROPERTY_ID" default:"61890231bdb53f8f36ea29e8"`
	LandLocationPriceId      string `env:"LAND_LOCATION_PRICE_ID" default:"6189029abdb53f8f36ea29e9"`
	LandCommunicationPriceId string `env:"LAND_COMMUNICATION" default:"61890346bdb53f8f36ea29ea"`
	LandAreaPropertyId       string `env:"LAND_AREA_PROPERTY_ID" default:"60ee8a5ec30b719d5f4e4686"`
	AreaPropertyId           string `env:"AREA_PROPERTY_ID" default:"611359b073bf6fe15aaef568"`
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
