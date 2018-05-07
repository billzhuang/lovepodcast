package main

import (
	"github.com/billzhuang/lovepodcast"
	"fmt"
)

func main() {
	ximalayaPodcast := lovepodcast.GetPodcast("http://www.ximalaya.com/6725746/album/238927/")
	fmt.Print("hello")

	fmt.Print(ximalayaPodcast.SoundIds + "\n")

	ximalayaEpisode := lovepodcast.GetEpisode("http://www.ximalaya.com/tracks/86846841.json")
	fmt.Print(ximalayaEpisode.Title + "\n")
}