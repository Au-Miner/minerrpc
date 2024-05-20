package serversServices

type TestImpl struct{}

func (s *TestImpl) Ping() (string, error) {
	return "pong", nil
}

func (s *TestImpl) Hello() (string, error) {
	return "name ", nil
}
