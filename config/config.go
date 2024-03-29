package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct for webapp config
type Config struct {
	Sever struct {
		ServerLevelDB struct {
			PathFile string `yaml:"pathfile"`
		} `yaml:"server_levedb"`

		ServerMySql struct {
			DBUserName string `yaml:"dbusername"`

			DBPassword string `yaml:"dbpassword"`

			DBName string `yaml:"dbname"`

			DBHost string `yaml:"dbhost"`

			DBPort string `yaml:"dbport"`
		} `yaml:"server_mysql"`

		ServerPostgersSql struct {
			DBUserName string `yaml:"dbusername"`

			DBPassword string `yaml:"dbpassword"`

			DBName string `yaml:"dbname"`

			DBHost string `yaml:"dbhost"`

			DBPort int `yaml:"dbport"`
		} `yaml:"server_postgressql"`

		ServerMongoDB struct {
			DBUserName string `yaml:"dbusername"`

			DBHost string `yaml:"dbhost"`

			DBPort string `yaml:"dbport"`
		} `yaml:"server_mongodb"`
	} `yaml:"server"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
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

	return config, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config/config.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

// Func main should be as small as possible and do as little as possible by convention
func GetConfig() *Config {

	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
