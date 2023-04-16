package main

import (
	"fmt"
	"getdolar/alertzy"
	"getdolar/dolar"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	key := os.Getenv("ALERTZY_API_KEY")
	if key == "" {
		fmt.Println("API key for Alertzy was not provided")
		os.Exit(-1)
	}
	fmt.Println("API key loaded just fine")
	r, err := dolar.GetDolar()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	fmt.Println("dolar info loaded fine")
	msg := fmt.Sprintf(`
	Dolar blue: %.2[1]f
	Dolar oficial: %.2[2]f
	`, r.BlueUSD.Sell, r.OficialUSD.Sell)
	err = alertzy.SendNotification(key, "Dolar Values", msg, 0, "Dolar Values")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	fmt.Println("dolar info sent")
}
