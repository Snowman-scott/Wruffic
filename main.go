package main

import (
	"fmt"
	"log"
	"os"

	"wruffic/api"
	"wruffic/cli"
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
		log.Fatal("WRUFF WRUFF, Ewwow in main :3")
	}

	client := api.NewClient(host, user, pass)
	p, err := player.New(client)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	for {
		query, err := cli.GetSearchQuery()
		if err != nil {
			fmt.Println(err)
			break
		}

		searchResults, err := client.SearchReq(query)
		if err != nil {
			fmt.Println(err)
			break
		}

		id, kind, err := cli.PickSongs(&api.Client{},searchResults)
		if err != nil {
			fmt.Println(err)
			break
		}
		if id == "" {
			continue
		}

		var Songs []api.Song
		if kind == "album" {
			Songs, err = client.GetAlbum(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else if kind == "song" {
			Songs, err = client.GetSong(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		if err := p.Player(Songs); err != nil {
			fmt.Println(err)
			break
		}

	}
	select {}
}
