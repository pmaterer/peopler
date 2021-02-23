package config

type Config struct {
	Server Server
}

type Server struct {
	ListenAddress string
	ListenPort    int64
}
