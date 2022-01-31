package config

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	// Conf ...
	Conf *Config
)

// Init init config
func Init(confPath string) error {
	err := initConfig(confPath)
	if err != nil {
		return err
	}
	return nil
}

// initConfig init config from conf file
func initConfig(confPath string) error {
	if confPath != "" {
		viper.SetConfigFile(confPath)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config.dev")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}

	// parse to config struct
	err := viper.Unmarshal(&Conf)
	if err != nil {
		return err
	}
	log.Info().Msgf("config:(%#v)", Conf)
	watchConfig()

	return nil
}

// watchConfig ...
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info().Msgf("Config file changed: %s", e.Name)
	})
}

// AppConfig ...
type AppConfig struct {
	Name      string `json:"name"`
	RunMode   string `json:"run_mode"`
	Addr      string `json:"addr"`
	JwtSecret string `json:"jwt_secret"`
	// JWTExpirationTime day
	JwtExpirationTime int `json:"jwt_expiration_time"`
}

type DBConfig struct {
	Uri             string `json:"uri"`
	MaxIdleConn     int    `json:"max_idel_conn"`
	MaxOpenConn     int    `json:"max_open_conn"`
	ConnMaxLifeTime int    `json:"conn_max_lifetime"`
}

type RedisConfig struct {
	Uri          string `json:"uri"`
	DialTimeout  int    `json:"dial_timeout"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	PoolSize     int    `json:"pool_size"`
}

// Config global config
// include common and biz config
type Config struct {
	// common
	App   AppConfig   `json:"app"`
	DB    DBConfig    `json:"db"`
	Redis RedisConfig `json:"redis"`
}
