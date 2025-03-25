package main

type Torrents struct {
	Title string `json:"title"`
	Hash  string `json:"hash"`
	Data  string `json:"data"`
}

type File struct {
	ID     int    `json:"id"`
	Path   string `json:"path"`
	Length int64  `json:"length"`
}

type Files struct {
	Files []File `json:"files"`
}

type Torrent struct {
	TorrServer Files `json:"torrServer"`
}

type Status struct {
	Stat string `json:"stat_string"`
}
