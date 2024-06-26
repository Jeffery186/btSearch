package main

import (
	"github.com/Bmixo/btSearch/common"
	"github.com/Bmixo/btSearch/model"
	"github.com/Bmixo/btSearch/service"
	"github.com/flosch/pongo2"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	service.InitConfig()
	common.Init()
	model.Init()
	service.Init()
	server := service.NewServer()
	pongo2.RegisterFilter("locFilter", server.FilterAddLoc)
	pongo2.RegisterFilter("keyFilter", server.FilterGetdbDataValueByKey)
	server.Router.GET("/search", server.Search)
	server.Router.GET("/movie/:id", server.Movie)
	server.Router.GET("/about", server.About)
	server.Router.GET("/", server.Index)
	server.Router.GET("/details/:objectid", server.Details)

	go server.Timer()
	go server.SyncDbHotSearchTimer()

	server.Router.Run(service.ConfigData.WebServerAddr)

}
