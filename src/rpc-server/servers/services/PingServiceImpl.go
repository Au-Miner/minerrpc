package services

type PingServiceImpl struct{}

func (s *PingServiceImpl) Ping() string {
	return "pong"
}

func (s *PingServiceImpl) ping(name string) string {
	return "pong " + name
}
