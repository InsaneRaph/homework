package config

import "flag"

type Config struct {
	JwtSecret  string
	DataSource string
	Port       int
	LogLevel   string
}

var AppConfig Config

func init() {
	port := flag.Int("port", 8888, "listen port")
	loglevel := flag.String("loglevel", "info", "log level [trace,debug,info,warn,error,fatal,panic]")
	secret := flag.String("secret", "changeme", "Jwt api secret")
	datasource := flag.String("ds",
		"root:root@tcp(127.0.0.1:3306)/homework?charset=utf8&parseTime=True",
		"mysql connection url",
	)

	flag.Parse()

	AppConfig = Config{
		JwtSecret:  *secret,
		DataSource: *datasource,
		Port:       *port,
		LogLevel:   *loglevel,
	}
}
