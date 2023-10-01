package main

import (
	"mosaics-web/api/env"
	"mosaics-web/api/helloworld"
	"mosaics-web/api/process"
	"mosaics-web/cable"

	"github.com/gin-gonic/gin"
)

const (
	GET  = "GET"
	POST = "POST"
)

type EndpointConfig struct {
	Resolver gin.HandlerFunc
	Path     string
	Method   string
}

type RouteConfig struct {
	Name      string
	Endpoints []EndpointConfig
}

var routesConfig = []RouteConfig{
	{
		"HelloWorldEndpoints",
		[]EndpointConfig{
			{Path: "/helloworld", Method: GET, Resolver: helloworld.ShowHelloWorld},
		},
	},
	{
		"EnvInfoEndpoints",
		[]EndpointConfig{
			{Path: "/envs", Method: GET, Resolver: env.ShowAllEnvs},
		},
	},
	{
		"ImageProcessorEndpoints",
		[]EndpointConfig{
			{Path: "/process", Method: POST, Resolver: process.ProcessUpload},
		},
	},
	{
		"WebsocketEndpoints",
		[]EndpointConfig{
			{Path: "/cable", Method: GET, Resolver: cable.HandleWebsocket},
		},
	},
}

func InitRouter(r *gin.Engine) {
	for _, route := range routesConfig {
		for _, endpoint := range route.Endpoints {
			r.Handle(endpoint.Method, endpoint.Path, endpoint.Resolver)
		}
	}
}
