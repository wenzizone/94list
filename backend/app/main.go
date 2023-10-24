package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

func init() {
	viper.setconfigfile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
