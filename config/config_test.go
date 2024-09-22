package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willychavez/order-listing-challenge/config"
)

func TestGet(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("ENV", "development")
	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "dbname")
	os.Setenv("HTTP_PORT", "8080")
	os.Setenv("GRPC_PORT", "9090")
	os.Setenv("GRAPHQL_PORT", "7070")

	defer func() {
		// Clean up environment variables after test
		os.Unsetenv("ENV")
		os.Unsetenv("DB_DRIVER")
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_PASSWORD")
		os.Unsetenv("DB_NAME")
		os.Unsetenv("HTTP_PORT")
		os.Unsetenv("GRPC_PORT")
		os.Unsetenv("GRAPHQL_PORT")
	}()

	cfg := config.Get()

	assert.Equal(t, "development", cfg.Environment)
	assert.Equal(t, "postgres", cfg.DB.Driver)
	assert.Equal(t, "localhost", cfg.DB.Host)
	assert.Equal(t, "5432", cfg.DB.Port)
	assert.Equal(t, "user", cfg.DB.User)
	assert.Equal(t, "password", cfg.DB.Password)
	assert.Equal(t, "dbname", cfg.DB.Name)
	assert.Equal(t, "8080", cfg.HTTP.Port)
	assert.Equal(t, "9090", cfg.GRPC.Port)
	assert.Equal(t, "7070", cfg.GraphQL.Port)
}

func TestGet_MissingRequiredEnv(t *testing.T) {
	// Unset required environment variable to simulate missing env var
	os.Unsetenv("ENV")

	defer func() {
		// Recover from panic to continue test execution
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to missing required environment variable")
		}
	}()

	config.Get()
}
