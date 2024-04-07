package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("PORT", "8080")
	os.Setenv("MONGO_URI", "mongodb://test:password@localhost:27017")
	os.Setenv("DB_NAME", "testdb")

	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("MONGO_URI")
		os.Unsetenv("DB_NAME")
	}()

	config := LoadConfig()

	assert.Equal(t, "8080", config.Port, "they should be equal")
	assert.Equal(t, "mongodb://test:password@localhost:27017", config.MongoURI, "they should be equal")
	assert.Equal(t, "testdb", config.DbName, "they should be equal")
}
