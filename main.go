package main

import (
	"fmt"
	"github.com/yeqianmen/go-gin-example/pkg/setting"
	"github.com/yeqianmen/go-gin-example/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed to start: %v", err)
	}
}
