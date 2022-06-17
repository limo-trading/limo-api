package main

import (
	"api/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

var routes = []Route{
	{method: "get", path: "/ping", handler: handler.PingHandler},
}
