package sonarclient

//ConfluenceConfig holds the current client configuration
type SonarQubeConfig struct {
	Username string
	Password string
	URL      string
	Debug    bool
}
