package conf

var Conf struct {
	HTTPListen string `env:"SERVER_ADDRESS"`
	BaseURL    string `env:"BASE_URL"`
	LogLevel   string `env:"LOG_LEVEL"`
	PgDsn      string `env:"DATABASE_DSN"`
	JWTSecret  string `env:"JWT_SECRET"`
}

func confLoad() {

}
