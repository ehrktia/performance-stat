package main

import (
	"fmt"
	"net/http"

	"github.com/ehrktia/performance-monitor/adapter/http_adapter"
)

func main() {
	fmt.Println("from main")
	routeHandler := http_adapter.New()
	routes := http_adapter.InitRoutes()
	routes["/"] = http.HandlerFunc(http_adapter.HomeHandler)

	if err := http.ListenAndServe(":9999", routeHandler.GetRouter()); err != nil {
		panic(fmt.Errorf("error starting http server in port `9999`:%v", err))
	}
}
