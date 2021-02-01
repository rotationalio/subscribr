package config

type Config struct {
	Name string `enconfig:"SUBSCRIBR_NAME" default:"subscriber"`
}
