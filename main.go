package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
)

var basic_auth_username string
var basic_auth_password string
var server_url string

func main() {
	err := loadEnv(".env")
	if err != nil {
		log.Fatal(err)
	}

	basic_auth_username = os.Getenv("BASIC_AUTH_USERNAME")
	basic_auth_password = os.Getenv("BASIC_AUTH_PASSWORD")
	server_url = os.Getenv("SERVER_URL")

	if server_url == "" {
		fmt.Println("Server Url require, please set .env file or environment variable")
		return
	}

	list := flag.Bool("ls", false, "list all torrents")
	view := flag.Int("view", 0, "view torrent detail")

	add := flag.Bool("add", false, "add magnet")
	magnetUrl := flag.String("u", "", "magnet url")
	magnetTitle := flag.String("t", "", "torrent title (not require)")

	link := flag.Int("link", 0, "")

	delete := flag.Int("rm", 0, "delete torrent")

	flag.Parse()

	if flag.NFlag() == 0 {
		help()
		os.Exit(1)
	}

	if err := req_torrents(); err != nil {
		fmt.Println(err)
		return
	}

	if *list {
		idWidth := 2
		nameWidth := 100

		printTableHeader(idWidth, nameWidth, "ID", "Name")
		for i, t := range torrents {
			printTableRow(i+1, t.Title, idWidth, nameWidth)
		}
		printTableFooter(idWidth, nameWidth)
	}

	if *view > 0 {
		torr := torrents[*view-1]
		if err := json.Unmarshal([]byte(torr.Data), &torrent); err != nil {
			fmt.Println(err)
			return
		}

		if *link == 0 {
			idWidth := 2
			nameWidth := 100

			printTableHeader(idWidth, nameWidth, "", torrents[*view-1].Title)
			for i, t := range torrent.TorrServer.Files {
				printTableRow(i+1, path.Base(t.Path), idWidth, nameWidth)
			}
			printTableFooter(idWidth, nameWidth)
		}
	}

	if *add {
		if *magnetUrl == "" {
			fmt.Println("--u Magnet Url require!")
			return
		}

		if err := req_add(*magnetUrl, *magnetTitle); err != nil {
			fmt.Println(err)
			return
		}
	}

	if *link > 0 {
		selected := torrent.TorrServer.Files[*link-1]

		filename := path.Base(selected.Path)

		idWidth := 2

		filenameEscaped := url.QueryEscape(filename)
		hashEscaped := url.QueryEscape(torrents[*view-1].Hash)

		link := fmt.Sprintf(server_url+"/stream/%s?link=%s&index=%d&play", filenameEscaped, hashEscaped, *link)

		printTableHeader(idWidth, len(link), "", fmt.Sprint(torrents[*view-1].Title))
		printTableRow(0, link, idWidth, len(link))
		printTableFooter(idWidth, len(link))

	}

	if *delete > 0 {
		var confirm string

		fmt.Println(torrents[*delete-1].Title)
		fmt.Println(torrents[*delete-1].Hash)

		fmt.Println("\nDo you wanna delete this torrent? (y or n)")
		fmt.Scanln(&confirm)

		if confirm == "y" || confirm == "Y" {
			req_delete(torrents[*delete-1].Hash)
			fmt.Println("Successfully deleted.")

		} else if confirm == "n" || confirm == "N" {
			fmt.Println("Canceled")
		} else {
			fmt.Println("Wrong answer!")
		}
	}

	fmt.Println("Done.")
}

func help() {
	helpMessage := `  _______                  _ _
 |__   __|                | (_)
    | | ___  _ __ _ __ ___| |_
    | |/ _ \| '__| '__/ __| | |
    | | (_) | |  | | | (__| | |
    |_|\___/|_|  |_|  \___|_|_|

╔═════════╦══════════════════════════════╗
║   Arg   ║            Usage             ║
╠═════════╬══════════════════════════════╣
║ -ls     ║ List all torrents            ║
╠═════════╬══════════════════════════════╣
║ -view   ║ View torrent episodes        ║
║ -link   ║ Get episode source link      ║
╠═════════╬══════════════════════════════╣
║ -add    ║ Add new torrent from magnet  ║
║   -u    ║   New torrent magnet url     ║
║   -t    ║   New torrent title          ║
╠═════════╬══════════════════════════════╣
║ -rm     ║ Delete Torrent               ║
╚═════════╩══════════════════════════════╝

╔═════════════════════╦══════════════════╗
║     Environment     ║       Must       ║
╠═════════════════════╬══════════════════╣
║ SERVER_URL          ║ Yes              ║
║ BASIC_AUTH_USERNAME ║ No               ║
║ BASIC_AUTH_PASSWORD ║ No               ║
╚═════════════════════╩══════════════════╝
`
	fmt.Println(helpMessage)
}
