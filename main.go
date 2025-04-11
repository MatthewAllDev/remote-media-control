package main

import (
	"RemoteMediaControl/internal/server"
	"RemoteMediaControl/internal/utils"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug mode")
	port := flag.String("port", "80", "Port to run the server on")
	flag.Parse()
	log.SetOutput(os.Stdout)
	ip, err := utils.GetServerIp()
	if err != nil {
		log.Fatal(err)
	}
	app, err := server.NewApp(*debug)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server started at http://%v:%v", ip, *port)
	log.Fatal(http.ListenAndServe(":"+*port, app.Mux))
}
