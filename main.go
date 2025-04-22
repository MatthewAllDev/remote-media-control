package main

import (
	"RemoteMediaControl/internal/server"
	"RemoteMediaControl/internal/utils"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const version = "v0.1"

func main() {
	debug := flag.Bool("debug", false, "Enable debug mode")
	port := flag.String("port", "80", "Port to run the server on")
	staticPath := flag.String("static-path", "./static", "Path to the directory with static files")
	about := flag.Bool("about", false, "Show information about this program")
	versionFlag := flag.Bool("version", false, "Show the version of the program")
	flag.Parse()
	if *versionFlag {
		fmt.Println("RemoteMediaControl", version)
		os.Exit(0)
	}
	if *about {
		fmt.Println("RemoteMediaControl", version)
		fmt.Println("Copyright (c) 2025 Ilia Kuvarzin")
		fmt.Println("\nReleased under the MIT License")
		fmt.Println("https://opensource.org/licenses/MIT")
		fmt.Println("\nSource: https://github.com/MatthewAllDev/remote-media-control")
		fmt.Println("Or: https://gitverse.ru/MatthewAllDev/remote-media-control")
		os.Exit(0)
	}
	if *debug {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
	ip, err := utils.GetServerIp()
	if err != nil {
		log.Fatal(err)
	}
	app, err := server.NewApp(*staticPath, *debug)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server started at http://%v:%v\n", ip, *port)
	log.Fatal(http.ListenAndServe(":"+*port, app.Mux))
}
