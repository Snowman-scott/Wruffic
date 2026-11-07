package main

import (
	"log"
	"os"

	"wruffic/api"
	"wruffic/player"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Your Wruffword Failed 3:")
	}

	host := os.Getenv("NAVI_HOST")
	user := os.Getenv("NAVI_USERNAME")
	pass := os.Getenv("NAVI_PASSWORD")
	if host == "" || user == "" || pass == "" {
		log.Fatal("missing NAVI_HOST / NAVI_USERNAME / NAVI_PASSWORD")
	}

	client := api.NewClient(host, user, pass)

	Songs, err := client.GetID()
	if err != nil {
		log.Fatal(err)
	}

	p, err := player.New(client)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	if err := p.PlayAlbum(Songs); err != nil {
		log.Fatal(err)
	}

	select {}
}
