package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName string           `envconfig:"service_name"`
	HttpServer  HttpServerConfig `envconfig:"http_server"`

	Database      DatabaseConfig      `envconfig:"database"`
	Redis         RedisConfig         `envconfig:"redis"`
	MessageStream MessageStreamConfig `envconfig:"message_stream"`
	OpenTelemetry OpenTelemetryConfig `envconfig:"open_telemetry"`
	GrpcServer    GrpcServerConfig    `envconfig:"grpc"`
	GrpcClient    GrpcClientConfig    `envconfig:"grpc_client"`
	HttpClient    HttpClientConfig    `envconfig:"http_client"`
	Logger        LoggerConfig        `envconfig:"logger"`
	ObjectStorage MinioConfig         `envconfig:"object_storage"`
}

type HttpServerConfig struct {
	Host string `envconfig:"http_server_host"`
	Port string `envconfig:"http_server_port"`
}

type MinioConfig struct {
	Endpoint        string `envconfig:"endpoint"`
	AccessKeyID     string `envconfig:"access_key_id"`
	SecretAccessKey string `envconfig:"secret_access_key"`
	UseSSL          bool   `envconfig:"use_ssl"`
	BucketName      string `envconfig:"bucket_name"`
}

type HttpClientConfig struct {
	Host                string  `envconfig:"http_client_host"`
	Port                string  `envconfig:"http_client_port"`
	Timeout             int     `envconfig:"http_client_timeout"`
	ConsecutiveFailures int     `envconfig:"http_client_consecutive_failures"`
	ErrorRate           float64 `envconfig:"http_client_error_rate"` // 0.001 - 0.999
	Threshold           int     `envconfig:"http_client_threshold"`
	Type                string  `envconfig:"http_client_type"` // consecutive, error_rate
}

type LoggerConfig struct {
	IsVerbose       bool   `envconfig:"logger_is_verbose"`
	LoggerCollector string `envconfig:"logger_logger_collector"`
}

type GrpcServerConfig struct {
	Host string `envconfig:"grpc_host"`
	Port string `envconfig:"grpc_port"`
}

type GrpcClientConfig struct {
	Host string `envconfig:"grpc_client_host"`
	Port string `envconfig:"grpc_client_port"`
}

type DatabaseConfig struct {
	Host         string `envconfig:"database_host"`
	Port         int    `envconfig:"database_port"`
	Username     string `envconfig:"database_username"`
	Password     string `envconfig:"database_password"`
	DBName       string `envconfig:"database_db_name"`
	SSL          string `envconfig:"database_ssl"`
	SchemaName   string `envconfig:"database_schema_name"`
	MaxIdleConns int    `envconfig:"database_max_idle_conns"`
	MaxOpenConns int    `envconfig:"database_max_open_conns"`
	Timeout      int    `envconfig:"database_timeout"`
}

type OpenTelemetryConfig struct {
	ExporterType             string `envconfig:"open_telemetry_exporter_type"`
	OtelExporterOLTPEndpoint string `envconfig:"open_telemetry_otel_exporter_otlp_endpoint"`
	OtelExporterOTLPInsecure bool   `envconfig:"open_telemetry_otel_exporter_otlp_insecure"`
}

type MessageStreamConfig struct {
	Host           string `envconfig:"message_stream_host"`
	Port           string `envconfig:"message_stream_port"`
	Username       string `envconfig:"message_stream_username"`
	Password       string `envconfig:"message_stream_password"`
	ExchangeName   string `envconfig:"message_stream_exchange_name"`
	PublishTopic   string `envconfig:"message_stream_publish_topic"`
	SubscribeTopic string `envconfig:"message_stream_subscribe_topic"`
}

type RedisConfig struct {
	Host            string        `envconfig:"redis_host"`
	Port            string        `envconfig:"redis_port"`
	Username        string        `envconfig:"redis_username"`
	Password        string        `envconfig:"redis_password"`
	DB              int           `envconfig:"redis_db"`
	MaxRetries      int           `envconfig:"redis_max_retries"`
	PoolFIFO        bool          `envconfig:"redis_pool_fifo"`
	PoolSize        int           `envconfig:"redis_pool_size"`
	PoolTimeout     time.Duration `envconfig:"redis_pool_timeout"`
	MinIdleConns    int           `envconfig:"redis_min_idle_conns"`
	MaxIdleConns    int           `envconfig:"redis_max_idle_conns"`
	ConnMaxIdleTime time.Duration `envconfig:"redis_conn_max_idle_time"`
	ConnMaxLifetime time.Duration `envconfig:"redis_conn_max_lifetime"`
}

// InitConfig reads config file from path

func InitConfig() *Config {
	var Cfg Config

	err := envconfig.Process("fww_core", &Cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Cfg
}
