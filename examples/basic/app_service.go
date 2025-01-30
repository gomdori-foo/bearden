package main

type AppService struct {}

// @Injectable()
func NewAppService() *AppService {
	return &AppService{}
}

func (s *AppService) GetHello() string {
	return "Hello, World!"
}
