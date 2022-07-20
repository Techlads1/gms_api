package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tzdit/sample_api/package/log"
	"github.com/tzdit/sample_api/services/database"
	"github.com/tzdit/sample_api/webserver"

	"github.com/tzdit/sample_api/package/config"
)

func init() {
	path := config.LoggerPath()
	fmt.Println(path)
	log.SetOptions(
		log.Development(),
		log.WithCaller(true),
		log.WithLogDirs(path),
	)
}
func main() {
	//migrations.Up()
	database.Connect()
	defer database.Close() //close database pool TODO: add this into graceful app close function
	go webserver.StartWebserver()
	// go routes.HandleAPIRequests()

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
	<-stop
	log.Infoln("Sample API is shutting down...  👋 !")
	fmt.Println("Sample API is shutting down .... 👋 !")
	database.Close()

	go func() {
		<-stop
		log.Fatalln("Sample API is terminating...")
	}()

	defer os.Exit(0)
}
