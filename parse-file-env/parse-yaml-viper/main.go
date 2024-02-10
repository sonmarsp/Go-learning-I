package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("DB.HOST", "127.0.0.1")

	// getting env variables DB.PORT
	// viper.Get() returns an empty interface{}
	// so we have to do the type assertion, to get the value
	DBHost, ok1 := viper.Get("DB.HOST").(string)
	DBuser, ok2 := viper.Get("DB.USERNAME").(string)
	DBPassword, ok3 := viper.Get("DB.PASWORD").(string)
	DBName, ok4 := viper.Get("DB.NAME").(string)

	// if type assert is not valid it will throw an error
	if !ok1 {
		log.Fatalf("Invalid type assertion")
	}
	if !ok2 {
		log.Fatalf("Invalid type assertion")
	}
	if !ok3 {
		log.Fatalf("Invalid type assertion")
	}
	if !ok4 {
		log.Fatalf("Invalid type assertion")
	}

	fmt.Println(DBHost)
	fmt.Println(DBuser)
	fmt.Println(DBPassword)
	fmt.Println(DBName)
}
