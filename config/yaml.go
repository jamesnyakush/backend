package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path"
)

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

/**
* Port smtp port
* Host smtp host e.g. smtp.gmail.com
* Username smtp username.
* Password smtp password.
* FromName email sender's name.
* FromAddress email sender's email address
 */
type SMTPConfig struct {
	Port        string `yaml:"port"`
	Host        string `yaml:"host"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	FromName    string `yaml:"from_name"`
	FromAddress string `yaml:"from_address"`
}

/**
* Port smtp port
* Host smtp host e.g. smtp.gmail.com
 */
type PesaswapConfig struct {
	ConsumerKey string `yaml:"consumer_key"`
	ApiKey      string `yaml:"apikey"`
}

/**
* Port smtp port
* Host smtp host e.g. smtp.gmail.com
* Username smtp username.
* Password smtp password.
* FromName email sender's name.
* FromAddress email sender's email address
 */
type MpesaConfig struct {
	Sandbox        string `yaml:"sandbox"`
	ConsumerKey    string `yaml:"consumer_key"`
	ConsumerSecret string `yaml:"consumer_secret"`
	Initiator      string `yaml:"initiator"`
	Paybill        string `yaml:"paybill"`
	Shortcode      string `yaml:"shortcode"`
	Passkey        string `yaml:"passkey"`
	ValidationUrl  string `yaml:"validation_url"`
	CallbackUrl    string `yaml:"callback_url"`
}

type SMSConfig struct {
	Username       string `yaml:"username"`
	ConsumerSecret string `yaml:"consumer_secret"`
	Shortcode      string `yaml:"shortcode"`
}

/**
* This maps the configuration in the yaml file into a struct
 */
type YamlConfig struct {
	Database  DatabaseConfig `yaml:"database"`
	SMTP      SMTPConfig     `yaml:"smtp"`
	Pesaswap  PesaswapConfig `yaml:"pesaswap"`
	Mpesa     MpesaConfig    `yaml:"mpesa"`
	SMS       SMSConfig      `yaml:"sms"`
	AppSecret string         `yaml:"app_secret_key"`
}

func ReadYaml(path string) *YamlConfig {
	if path == "" {
		path = defaultYamlConfigPath()
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer func() { _ = f.Close() }()

	var cfg YamlConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Printf("error reading yaml file into config struct: %s\n", err)
		os.Exit(2)
	}
	return &cfg
}

/**
* Reads the path of the current executable
* goes up 2 directories and appends config.yaml to the path.
 */
func defaultYamlConfigPath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Printf("error encountered reading path: %s\n", err)
		os.Exit(2)
	}

	filename := "config.yml"
	dir := path.Dir(path.Dir(ex))
	dir = path.Join(dir, filename)
	return dir
}
