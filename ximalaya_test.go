package lovepodcast

import "testing"

func GetPodcast_test(t *testing.T){
	ximalayaPodcast := GetPodcast("http://www.ximalaya.com/6725746/album/238927/")

	t.Log(ximalayaPodcast.SoundIds)
}

func GetEpisode_test(t *testing.T){
	ximalayaEpisode := GetEpisode("http://www.ximalaya.com/tracks/86846841.json")


}
