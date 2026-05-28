package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/altfoxie/drpc"
)

type TrackInfo struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Img    string `json:"img"`
}

var client *drpc.Client

func buildExternalImageURL(rawURL string) string {
	if rawURL == "" || !strings.HasPrefix(rawURL, "http") {
		return "logo"
	}

	cleanURL := rawURL
	if idx := strings.Index(cleanURL, "?"); idx != -1 {
		cleanURL = cleanURL[:idx]
	}

	if idx := strings.Index(cleanURL, "=w"); idx != -1 {
		cleanURL = cleanURL[:idx]
	}
	cleanURL = cleanURL + "=w512-h512-l90-rj"

	return cleanURL
}

func updateDiscordRPC(track TrackInfo) {
	if client == nil {
		return
	}

	imageURL := buildExternalImageURL(track.Img)

	err := client.SetActivity(drpc.Activity{
		Details: track.Title,
		State:   "Artist: " + track.Artist,
		Assets: &drpc.Assets{
			LargeImage: imageURL,
			LargeText:  "YouTube Music",
		},
		Timestamps: &drpc.Timestamps{
			Start: time.Now(),
		},
		Buttons: []drpc.Button{
			{
				Label: "Open YouTube Music",
				URL:   "https://music.youtube.com",
			},
		},
	})
	if err != nil {
		log.Printf("Error updating Discord RPC: %v", err)
	} else {
		fmt.Printf("Playing: %s — %s\n", track.Title, track.Artist)
		fmt.Printf("Image URL: %s\n", imageURL)
	}
}

func trackHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://music.youtube.com")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		var track TrackInfo
		err := json.NewDecoder(r.Body).Decode(&track)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("RAW JSON received: title=%q artist=%q img=%q\n", track.Title, track.Artist, track.Img)

		updateDiscordRPC(track)
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	client = drpc.New("1509651617434042368")

	http.HandleFunc("/track", trackHandler)
	http.HandleFunc("/update", trackHandler)

	fmt.Println("Listening on port 8080")
	fmt.Println("Wait track")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
