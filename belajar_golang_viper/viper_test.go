package belajar_golang_viper

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Library Viper digunakan untuk membaca file konfigurasi seperti env, json, yaml
*/

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigType("json")
	config.SetConfigName("config")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "belajar_golang_viper", config.GetString("app.name"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "belajar_golang_viper", config.GetString("APP_NAME"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOWSQL"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
}

func TestFromENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "belajar_golang_viper", config.GetString("APP_NAME"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOWSQL"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
	assert.Equal(t, "Hi", config.GetString("FROM_ENV"))
}
