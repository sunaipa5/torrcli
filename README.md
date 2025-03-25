# Torrcli

CLI for YouROK/TorrServer (MatriX)

## Features

- List all torrents
- View torrent episodes
- Get torrent media play link
- Add new torrent
- Delete torrent

## Roadmap

- View torrent buffer

## Installation

```bash
git clone https://github.com/sunaipa5/torrcli
```

```bash
cd torrcli
```

```bash
go build
```

Run the program in the terminal

```bash
torrcli
```

You can create .env file (recommended) where the program is located or you can start program with env

`SERVER_URL must`, if you are using basic auth you need to set `BASIC_AUTH_USERNAME` and `BASIC_AUTH_PASSWORD` env

```env
SERVER_URL=https://server.url
BASIC_AUTH_USERNAME=john
BASIC_AUTH_PASSWORD=johndoe123
```

## CLI view

```
user@linux:~$ torrcli
  _______                  _ _
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
```
