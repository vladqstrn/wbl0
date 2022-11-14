package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/vladqstrn/l0/internal/wbl0"
)

func main() {
	//настройка и чтение конфига
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Unable to read config file")
	}

	dbConn := wbl0.StartDatabaseConnection()
	log.Println("database started!")

	cm := wbl0.CreateCacheManager(dbConn)
	log.Println("сache restored!")

	natsConn := wbl0.NatsConnection()
	log.Println("nats-streaming started!")

	domain := viper.GetString("AppServer.domain")
	port := viper.GetString("AppServer.port")

	srv := http.Server{
		Addr:    domain + ":" + port,
		Handler: wbl0.InitRouter(cm),
	}

	app := wbl0.CreateApp(dbConn, natsConn, &srv, cm)

	app.Run()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit

}
