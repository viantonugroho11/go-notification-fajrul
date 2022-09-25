package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	consulEndpoint = "127.0.0.1:8500"
	consulPath     = "NOTIF-ENGINE" //Consul Path for load remote config
)

type Config struct {
	Env      string            `mapstructure:"env"`
	Port     int               `mapstructure:"port"`
	Database SQL              `mapstructure:"database"`
	Email    Email             `mapstructure:"email"`
	Queue    Queue             `mapstructure:"queue"`
	Firebase Firebase          `mapstructure:"firebase"`
	Server   ApplicationConfig `mapstructure:"server"`
}

type Queue struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dial     string `mapstructure:"dial"`
}

type Email struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	Email         string `mapstructure:"email"`
	Password      string `mapstructure:"password"`
	AddressMail   string `mapstructure:"address_mail"`
	FromNameEmail string `mapstructure:"from_name_email"`
}

type Firebase struct {
	FCMClientID string `mapstructure:"fcm_client_id"`
}

type ApplicationConfig struct {
	UserAddress     string `mapstructure:"user_address"` //grpc address
	UserAddressHttp string `mapstructure:"user_address_http"`
	ArticleAddress  string `mapstructure:"article_address"`
}

type PostgresConfig struct {
	ConnMaxLifetime    int  `mapstructure:"connMaxLifetime"`
	MaxOpenConnections int  `mapstructure:"maxOpenConnections"`
	MaxIdleConnections int  `mapstructure:"maxIdleConnections"`
	ConnectTimeout     int  `mapstructure:"connectTimeout"`
	Master             SQL `mapstructure:"master"`
	Slave              SQL `mapstructure:"slave"`
}

type SQL struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
	DBName   string `mapstructure:"dbName"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func New() (conf Config) {
	var once sync.Once
	once.Do(func() {
		v := viper.New()
		retried := 0
		err := InitialiseRemote(v, retried)
		if err != nil {
			log.Printf("No remote server configured will load configuration from file and environment variables: %+v", err)
			if err := InitialiseFileAndEnv(v, "config.local"); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					configFileName := fmt.Sprintf("%s.yaml", "config.local")
					log.Printf("No '" + configFileName + "' file found on search paths. Will either use environment variables or defaults")
				} else {
					log.Fatalf("Error occured during loading config: %s", err.Error())
				}
			}
		}
		err = v.Unmarshal(&conf)
		if err != nil {
			log.Fatalf("%v", err)
		}
	})
	return conf
}

func InitialiseRemote(v *viper.Viper, retried int) error {
	if consulEnv := os.Getenv("CONSUL_URL"); consulEnv != "" {
		consulEndpoint = consulEnv
	}
	log.Printf("Initialising remote config, consul endpoint: %s, consul path: %s, retried: %d", consulEndpoint, consulPath, retried)
	v.AddRemoteProvider("consul", consulEndpoint, consulPath)
	v.SetConfigType("yaml")
	err := v.ReadRemoteConfig()
	if err != nil && retried < 1 {
		time.Sleep(500 * time.Millisecond)
		return InitialiseRemote(v, retried+1)
	}
	return err
}

func InitialiseFileAndEnv(v *viper.Viper, configName string) error {
	var searchPath = []string{
		"/notif-engine",
		"$HOME/.notif-engine",
		".",
	}
	v.SetConfigName(configName)
	for _, path := range searchPath {
		v.AddConfigPath(path)
	}
	v.SetEnvPrefix("notif-engine")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	return v.ReadInConfig()
}
