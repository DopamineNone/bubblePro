package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sync"
)

const (
	EnvTypeKey    = "GO_MOD"
	TestEnv       = "test"
	ProductionEnv = "production"
)

// singleton patter
var (
	conf    = new(AppConfig)
	once    sync.Once
	EnvType string
)

type AppConfig struct {
	Name           string         `yaml:"name"`
	Version        string         `yaml:"version"`
	StartTime      string         `yaml:"start_time"`
	MachineID      int64          `yaml:"machine_id"`
	Port           int            `yaml:"port"`
	MySQL          MySQL          `yaml:"mysql"`
	Log            Log            `yaml:"log"`
	Security       Security       `yaml:"security"`
	Authentication Authentication `yaml:"auth"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Log struct {
	FilePath string
}

// Security contains security setting like hash salt length, all measured in bytes
type Security struct {
	SaltLength int `yaml:"salt_length"`
	HashLength int `yaml:"hash_length"`
	Iterations int `yaml:"iterations"`
}

type Authentication struct {
	AccessTokenExpireDuration  int `yaml:"access_token_expire_duration"`
	RefreshTokenExpireDuration int `yaml:"refresh_token_expire_duration"`
}

func GetConf() *AppConfig {
	once.Do(initConfig)
	return conf
}

func initConfig() {
	// get environment type
	EnvType = os.Getenv(EnvTypeKey)
	if len(EnvType) == 0 {
		EnvType = TestEnv
	}

	// load test environment variables
	if EnvType == TestEnv {
		err := godotenv.Load(filepath.Join("src/config", EnvType, ".env"))
		if err != nil {
			panic(err)
		}
	}

	// marshal config
	f, err := os.Open(filepath.Join("src/config", EnvType, "config.yaml"))
	if err != nil {
		panic(err)
	}
	err = yaml.NewDecoder(f).Decode(conf)
	if err != nil {
		panic(err)
	}
}

func GetEnv(key string) string {
	once.Do(initConfig)
	return os.Getenv(key)
}
