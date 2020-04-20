package apiserver

//APIServer ...
type APIServer struct{
	config *Config
}

//New ...
func New (config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

//Start APIServer
func (s *APIServer) Start() error {
	return nil
}
