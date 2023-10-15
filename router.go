package main

import (
	"path"

	"mosaics-web/api/process"
	"mosaics-web/cable"
	"mosaics-web/env"

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
	r.Static("/img", path.Join(env.GetBaseDir(), env.GetResultDir()))
}
