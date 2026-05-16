package models

type Config struct {
	Name        string `json:"name" yaml:"name"`
	Port        int    `json:"port" yaml:"port"`
	Environment string `json:"environment" yaml:"environment"`
}
