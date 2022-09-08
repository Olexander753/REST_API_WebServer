package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/Olexander753/REST_API_WebServer/internal/config"
	"github.com/Olexander753/REST_API_WebServer/internal/user"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Create router")
	router := httprouter.New()

	log.Println("Read config")
	cfg := config.GetConfig()

	// log.Println("Create NewClient")
	// mongoDBClient, err := mongodb.NewClient(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port,
	// 	cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthDB)
	// if err != nil {
	// 	//log.Println(err)
	// 	panic(err)
	// }

	// log.Println("Create NewStorage")

	// storage := db.NewStorage(mongoDBClient, cfg.MongoDB.Database, cfg.MongoDB.Collection)

	log.Println("Register user handler")
	handler := user.NewHanler()
	handler.Register(router)

	Start(router, cfg)

}

func Start(router *httprouter.Router, cfg *config.Config) {
	log.Println("Start app")

	listner, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.IP, cfg.Listen.Port))
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server listening %s\n", fmt.Sprintf("%s:%s", cfg.Listen.IP, cfg.Listen.Port))
	server.Serve(listner)
}
