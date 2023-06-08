package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorf/gorf"
	"github.com/go-gorf/gorf/backends/gormi"
	"log"
)

import "template/apps/hello"

// add all the apps
var apps = []gorf.GorfApp{
	&hello.HelloApp,
}

func LoadSettings() {
	// jwt secret key
	gorf.Settings.SecretKey = "GOo8Rs8ht7qdxv6uUAjkQuopRGnql2zWJu08YleBx6pEv0cQ09a"
	gorf.Settings.DbBackends = &gormi.GormSqliteBackend{
		Name: "db.sqlite",
	}
}

// bootstrap server
func BootstrapRouter() *gin.Engine {
	gorf.Apps = append(apps)
	LoadSettings()
	err := gorf.InitializeDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	gorf.SetupApps()
	r := gin.Default()
	gorf.RegisterApps(r)
	return r
}
