package config

// ServerBlock :
type ServerBlock struct {
	CORS struct {
		AllowOrigins     []string `env:"CORS_ALLOW_ORIGINS" envSeparator:"," envDefault:"*" json:"allowOrigins,omitempty"`
		AllowMethods     []string `env:"CORS_ALLOW_METHODS" envSeparator:"," envDefault:"GET,HEAD,PUT,PATCH,POST,DELETE" json:"allowMethods,omitempty"`
		AllowHeaders     []string `env:"CORS_ALLOW_HEADERS" envSeparator:"," envDefault:"*" json:"allowHeaders,omitempty"`
		AllowCredentials bool     `env:"CORS_ALLOW_CREDENTIALS" envDefault:"false" json:"allowCredentials,omitempty"`
	}
}

// Server :
var Server ServerBlock
