package app

import (
	"fmt"
	"github.com/nyumbapoa/backend/configs"
)

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

type SMTP struct {
	Port        string
	Host        string
	Username    string
	Password    string
	FromName    string
	FromAddress string
}

type Pesaswap struct {
	ConsumerKey string
	ApiKey      string
}

type Mpesa struct {
	Sandbox        string
	ConsumerKey    string
	ConsumerSecret string
	Initiator      string
	Paybill        string
	Shortcode      string
	Passkey        string
	ValidationUrl  string
	CallbackUrl    string
}

type SMS struct {
	Username       string
	ConsumerSecret string
	Shortcode      string
}

func (d Database) String(sslmode string) string {
	return fmt.Sprintf(""+
		"host=%s "+
		"port=%s "+
		"user=%s "+
		"dbname=%s "+
		"password=%s "+
		"sslmode=%v"+
		"", d.Host, d.Port, d.User, d.DBName, d.Password, sslmode)
}

type Config struct {
	DB       Database
	Smtp     SMTP
	PesaSwap Pesaswap
	MPesa    Mpesa
	Sms      SMS

	Secret string
}

func GetConfig(cfg configs.YamlConfig) Config {
	return Config{
		DB: Database{
			User:     cfg.Database.User,
			Password: cfg.Database.Password,
			Host:     cfg.Database.Host,
			Port:     cfg.Database.Port,
			DBName:   cfg.Database.DBName,
		},
		Smtp: SMTP{
			Port:        cfg.SMTP.Port,
			Host:        cfg.SMTP.Host,
			Username:    cfg.SMTP.Username,
			Password:    cfg.SMTP.Password,
			FromName:    cfg.SMTP.FromName,
			FromAddress: cfg.SMTP.FromAddress,
		},
		PesaSwap: Pesaswap{
			ConsumerKey: cfg.Pesaswap.ConsumerKey,
			ApiKey:      cfg.Pesaswap.ApiKey,
		},
		MPesa: Mpesa{
			Sandbox:        cfg.Mpesa.Sandbox,
			ConsumerKey:    cfg.Mpesa.ConsumerKey,
			ConsumerSecret: cfg.Mpesa.ConsumerSecret,
			Initiator:      cfg.Mpesa.Initiator,
			Paybill:        cfg.Mpesa.Paybill,
			Shortcode:      cfg.Mpesa.Shortcode,
			Passkey:        cfg.Mpesa.Passkey,
			ValidationUrl:  cfg.Mpesa.ValidationUrl,
			CallbackUrl:    "",
		},
		Sms: SMS{
			Username:       cfg.SMS.Username,
			ConsumerSecret: cfg.SMS.ConsumerSecret,
			Shortcode:      cfg.SMS.Shortcode,
		},
		Secret: cfg.AppSecret,
	}
}