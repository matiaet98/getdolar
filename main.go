package main

import (
	"fmt"
	"getdolar/alertzy"
	"getdolar/dolar"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	key := os.Getenv("ALERTZY_API_KEY")
	if key == "" {
		log.Fatal("API key for Alertzy was not provided")
	}
	log.Println("API key loaded just fine")
	r, err := dolar.GetDolar()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("dolar info loaded fine")
	msg := fmt.Sprintf(`
	Dolar blue: %.2[1]f
	Dolar oficial: %.2[2]f
	`, r.BlueUSD.Sell, r.OficialUSD.Sell)
	err = alertzy.SendNotification(key, "Dolar Values", msg, 0, "Dolar Values")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("dolar info sent")
}
