package newsapi

// Config for managing connection with News API.
type Config struct {
	Host struct {
		Key string `envconfig:"default=97c8beb3e81b44e5a7b290f41d02928f"`
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
