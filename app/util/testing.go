package util

import "github.com/jacobstr/confer"

// GetTestingConfig returns a Confer configuration object with some useful
// testing defaults.
func GetTestingConfig() *confer.Config {
	config := confer.NewConfig()

	config.Set("database.driver", "sqlite")
	config.Set("database.file", ":memory:")

	return config
}
