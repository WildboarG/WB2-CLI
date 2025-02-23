package config

type Config struct {
	SDK          string `json:"sdk"`
	FromLocation string `json:"-"`
}
