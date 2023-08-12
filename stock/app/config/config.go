package config

type Config struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	CertFile     string
	KeyFile      string
}
