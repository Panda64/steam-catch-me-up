package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

// --------------------------------------------------------------------------------
// Structs
// --------------------------------------------------------------------------------

type Game struct {
	Title          string
	CurrentPlayers string
	PeakToday      string
	GameLink       string
}

type Video struct {
	Title   string
	Link    string
	Channel string
}

type Result struct {
	Game      Game
	GameVideo Video
}

type YoutubeResponse struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		Etag string `json:"etag"`
		ID   struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string    `json:"channelTitle"`
			LiveBroadcastContent string    `json:"liveBroadcastContent"`
			PublishTime          time.Time `json:"publishTime"`
		} `json:"snippet"`
	} `json:"items"`
}

// --------------------------------------------------------------------------------
// Helper Functions
// --------------------------------------------------------------------------------

func steam_stats() []Game {
	var games = []Game{}

	// Instantiate default collector
	c := colly.NewCollector()

	// Print status before visiting
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Getting real-time steam stats...")
	})

	// Add relevant data to games array (top 10 games)
	c.OnHTML("#detailStats > table > tbody", func(e *colly.HTMLElement) {
		for count := 3; count < 12; count++ {
			game := Game{
				Title:          e.ChildText("tr:nth-child(" + strconv.Itoa(count) + ") > td:nth-child(4)"),
				CurrentPlayers: e.ChildText("tr:nth-child(" + strconv.Itoa(count) + ") > td:nth-child(1)"),
				PeakToday:      e.ChildText("tr:nth-child(" + strconv.Itoa(count) + ") > td:nth-child(2)"),
				GameLink:       e.ChildAttr("tr:nth-child("+strconv.Itoa(count)+") > td:nth-child(4) > a", "href"),
			}
			games = append(games, game)
		}
	})

	// If something goes wrong, print the error
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Visit the steam statistics page
	c.Visit("https://store.steampowered.com/stats/Steam-Game-and-Player-Statistics")

	return games
}

func youtube(game_title string) Video {
	// Youtube API Key
	var developerKey = os.Getenv("YOUTUBE_API_KEY")

	// The date 7 days ago from today. Used to get the most recent videos
	var last_week = time.Now().AddDate(0, 0, -7).Format(time.RFC3339)
	last_week = strings.ReplaceAll(last_week, ":", "%3A")

	game_title = strings.ReplaceAll(game_title, " ", "%20")

	// API URL that includes all relevant parameters:
	// Max Results: 1
	// Order: View Count,
	// Published After: Last Week
	// Relevence Language: English
	var url = "https://youtube.googleapis.com/youtube/v3/search?part=snippet&maxResults=1&order=viewCount&publishedAfter=" + last_week + "&q=" + game_title + "&relevanceLanguage=en&key=" + developerKey

	// Sending GET request to YouTube API
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Reading response body
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Put JSON formatted response body into struct
	var result YoutubeResponse
	if err := json.Unmarshal(responseData, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	// Add relevant data to video struct
	video := Video{
		Title:   result.Items[0].Snippet.Title,
		Link:    "https://www.youtube.com/watch?v=" + result.Items[0].ID.VideoID,
		Channel: result.Items[0].Snippet.ChannelTitle,
	}

	return video
}

// --------------------------------------------------------------------------------
// Main
// --------------------------------------------------------------------------------

func main() {

	// Initialize dotenv
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get steam + YouTube data
	var games = steam_stats()
	var results = []Result{}

	// Add all stats to results array
	for _, game := range games {
		var video = youtube(game.Title)

		result := Result{
			Game:      game,
			GameVideo: video,
		}

		results = append(results, result)
	}

	// Print results
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Let's catch you up on what people are playing today...\n")
	for _, result := range results {
		fmt.Println("Game: " + result.Game.Title)
		fmt.Println("Current Players: " + result.Game.CurrentPlayers)
		fmt.Println("Peak Today: " + result.Game.PeakToday)
		fmt.Println("Trending Video: " + "'" + result.GameVideo.Title + "'" + " ---> by " + result.GameVideo.Channel)
		fmt.Println(result.GameVideo.Link)
		fmt.Println("\n")
	}
	fmt.Println("--------------------------------------------------------------------------------")
}
