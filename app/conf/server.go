package conf

type Server struct {
	Bind string `env:"SERVER_BIND" envDefault:":3000"`
}
