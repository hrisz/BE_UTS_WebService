package _anidata

import (
	"fmt"
	"testing"
)

func TestInsertDataAnime(t *testing.T) {
	uid := "GI_170424_V"
	animation := AnimationData{
		UniversalID: "GI_170424",
		Title:       "'The Song Burning in the Embers' Full Animated Short",
		Producer:    "Hoyo Anime",
		Games:       "Genshin Impact",
		Dates:       "April, 17th 2024",
		URLVid:      "https://youtu.be/tN5JACOEJFM",
	}
	views := 6054984
	likes := 398950
	comments := 12297

	insertData := InsertDataAnime(uid, animation, views, likes, comments)
	fmt.Println(insertData)
}

func TestGetDataAnimeWithUID(t *testing.T) {
	uid := "GI_170424_V"
	animation := GetDataAnimeWithUID(uid)
	fmt.Println(animation)
}

func TestGetAll(t *testing.T) {
	data := GetAllDataAnime()
	fmt.Println(data)
}
