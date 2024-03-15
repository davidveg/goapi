package config

import (
	"github.com/magiconair/properties"
	"log"
)

func CreateProperties() *properties.Properties {
	p, err := LoadProperties()
	if err != nil {
		log.Fatalf("ERROR : %v", err)
		return nil
	}
	return p
}

func LoadProperties() (*properties.Properties, error) {
	// Carregar o arquivo de propriedades
	p, err := properties.LoadFile("config.properties", properties.UTF8)
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo de propriedades: %v", err)
		return nil, err
	}
	return p, nil
}
