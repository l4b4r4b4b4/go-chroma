// TODO add creacefullness from https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/
package main

import (
	"log"
	"net/http"

	handler "github.com/l4b4r4b4b4/go-chroma/server/handler/collection"
	"github.com/l4b4r4b4b4/kitchen/go-chroma/server/service/collection"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}