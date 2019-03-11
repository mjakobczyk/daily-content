package newsapi

// Config for managing connection with News API.
type Config struct {
	Host struct {
		Key string `envconfig:"default="`
		API string `envconfig:"default=https://newsapi.org/v2"`
	}

	HTTP struct {
		Client struct {
			Timeout struct {
				Seconds int `envconfig:"default=10"`
			}
		}
	}
}
