package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {


	env, err := LoadEnv(".")

	assert.NoError(t, err)

	expectedEnv := EnvConfig{
		Port:                   ":8080",
		DBSource:               "mysql://username:password@localhost:3306/database",
		MigrationURL:           "file://db/migrations",
		CrawlerAddress:         "localhost:8000",
		GoogleClientID:         "your-google-client-id",
		GoogleClientSecret:     "your-google-client-secret",
		GoogleOAuthRedirectUrl: "http://localhost:8080/auth/callback",
		FrontEndOrigin:         "http://localhost:3000",
	}
	assert.Equal(t, expectedEnv, env)
}
