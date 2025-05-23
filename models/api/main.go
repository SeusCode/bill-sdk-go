package api

type Enviroment uint
type Language string

const (
	// Enviroment Types
	TESTING    Enviroment = 1
	PRODUCTION Enviroment = 0

	// Language Types
	SPANISH Language = "es"
	ENGLISH Language = "en"
)
