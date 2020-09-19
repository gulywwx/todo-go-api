package conf

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Jwt struct {
			Secret     string        `yaml:"secret"`
			ExpireTime time.Duration `yaml:"expiretime"`
		} `yaml:"jwt"`
		Port         int           `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"readtimeout"`
		WriteTimeout time.Duration `yaml:"writetimeout"`
	} `yaml:"server"`
	Log struct {
		Path       string `yaml:"path"`
		FileExt    string `yaml:"fileext"`
		TimeFormat string `yaml:"timeformat"`
	} `yaml:"log"`
}

type Jwt struct {
	Secret     string
	ExpireTime time.Duration
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}
var JwtSetting = &Jwt{}

func Load() (*Config, error) {
	config := &Config{}
	// Open config file
	file, err := os.Open("todo.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	ServerSetting.HttpPort = config.Server.Port
	ServerSetting.ReadTimeout = config.Server.ReadTimeout
	ServerSetting.WriteTimeout = config.Server.WriteTimeout

	JwtSetting.Secret = config.Server.Jwt.Secret
	JwtSetting.ExpireTime = config.Server.Jwt.ExpireTime

	return config, nil
}
