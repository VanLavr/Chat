package config

import (
	"log"

	"github.com/spf13/viper"
)

var Con = NewConfigurator()

type Configurator struct {
	v *viper.Viper
}

func NewConfigurator() *Configurator {
	v := viper.New()
	v.SetConfigFile("./.env")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	return &Configurator{
		v: v,
	}
}

func (c *Configurator) GetPort() (port string) {
	port = c.v.GetString("PORT")
	return
}

func (c Configurator) GetAddress() (addr string) {
	addr = c.v.GetString("ADDR")
	return
}

func (c Configurator) GetLogFiles() (string, string, string) {
	errorPath := c.v.GetString("ERRORFILE")
	warningPath := c.v.GetString("WARNINGFILE")
	infoPath := c.v.GetString("INFOFILE")

	return errorPath, warningPath, infoPath
}

func (c Configurator) GetDBConnectionString() string {
	return c.v.GetString("CONNECTIONSTRING")
}

func (c Configurator) GetSecret() string {
	return c.v.GetString("SECRET")
}

func (c Configurator) GetOrigin() string {
	return c.v.GetString("ORIGIN")
}

func (c Configurator) GetPostgres() string {
	return c.v.GetString("DSN")
}
