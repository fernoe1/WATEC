package config

import "time"

const GRPC_PORT = "GRPC_PORT"

type Config struct {
	Server
	Logger
}

type Server struct {
	Port              string
	Development       bool
	Timeout           time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
}

type Logger struct {
	Encoding   string
	Level      string
	Caller     bool
	Stacktrace bool
}
