package app

import (
	"fmt"
	"github.com/nyumbapoa/backend/config"
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
	DB   Database
	Smtp SMTP

	Secret string
}

func GetConfig(cfg config.YamlConfig) Config {
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

		Secret: cfg.AppSecret,
	}
}
