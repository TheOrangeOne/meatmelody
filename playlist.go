package main

type Playlist struct {
	items []string
}

func newPlaylist() *Playlist {
	return &Playlist{
		items: make([]string, 50),
	}
}
