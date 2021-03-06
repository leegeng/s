package main

import (
	"fmt"
	"os"

	// Load necessary providers.
	_ "github.com/zquestz/s/providers/8tracks"
	_ "github.com/zquestz/s/providers/amazon"
	_ "github.com/zquestz/s/providers/arxiv"
	_ "github.com/zquestz/s/providers/atmospherejs"
	_ "github.com/zquestz/s/providers/baidu"
	_ "github.com/zquestz/s/providers/bandcamp"
	_ "github.com/zquestz/s/providers/bing"
	_ "github.com/zquestz/s/providers/codepen"
	_ "github.com/zquestz/s/providers/coursera"
	_ "github.com/zquestz/s/providers/crunchyroll"
	_ "github.com/zquestz/s/providers/digg"
	_ "github.com/zquestz/s/providers/dockerhub"
	_ "github.com/zquestz/s/providers/dribbble"
	_ "github.com/zquestz/s/providers/duckduckgo"
	_ "github.com/zquestz/s/providers/dumpert"
	_ "github.com/zquestz/s/providers/facebook"
	_ "github.com/zquestz/s/providers/flickr"
	_ "github.com/zquestz/s/providers/flipkart"
	_ "github.com/zquestz/s/providers/foursquare"
	_ "github.com/zquestz/s/providers/gist"
	_ "github.com/zquestz/s/providers/github"
	_ "github.com/zquestz/s/providers/gmail"
	_ "github.com/zquestz/s/providers/go"
	_ "github.com/zquestz/s/providers/godoc"
	_ "github.com/zquestz/s/providers/google"
	_ "github.com/zquestz/s/providers/hackernews"
	_ "github.com/zquestz/s/providers/ietf"
	_ "github.com/zquestz/s/providers/ifttt"
	_ "github.com/zquestz/s/providers/imdb"
	_ "github.com/zquestz/s/providers/imgur"
	_ "github.com/zquestz/s/providers/inbox"
	_ "github.com/zquestz/s/providers/kickasstorrents"
	_ "github.com/zquestz/s/providers/libgen"
	_ "github.com/zquestz/s/providers/linkedin"
	_ "github.com/zquestz/s/providers/lmgtfy"
	_ "github.com/zquestz/s/providers/macports"
	_ "github.com/zquestz/s/providers/mdn"
	_ "github.com/zquestz/s/providers/msdn"
	_ "github.com/zquestz/s/providers/naver"
	_ "github.com/zquestz/s/providers/nhaccuatui"
	_ "github.com/zquestz/s/providers/npm"
	_ "github.com/zquestz/s/providers/npmsearch"
	_ "github.com/zquestz/s/providers/nvd"
	_ "github.com/zquestz/s/providers/packagist"
	_ "github.com/zquestz/s/providers/php"
	_ "github.com/zquestz/s/providers/pinterest"
	_ "github.com/zquestz/s/providers/python"
	_ "github.com/zquestz/s/providers/quora"
	_ "github.com/zquestz/s/providers/reddit"
	_ "github.com/zquestz/s/providers/rubygems"
	_ "github.com/zquestz/s/providers/soundcloud"
	_ "github.com/zquestz/s/providers/spotify"
	_ "github.com/zquestz/s/providers/stackoverflow"
	_ "github.com/zquestz/s/providers/steam"
	_ "github.com/zquestz/s/providers/taobao"
	_ "github.com/zquestz/s/providers/thepiratebay"
	_ "github.com/zquestz/s/providers/twitchtv"
	_ "github.com/zquestz/s/providers/twitter"
	_ "github.com/zquestz/s/providers/unity3d"
	_ "github.com/zquestz/s/providers/upcloud"
	_ "github.com/zquestz/s/providers/vimeo"
	_ "github.com/zquestz/s/providers/wikipedia"
	_ "github.com/zquestz/s/providers/wolframalpha"
	_ "github.com/zquestz/s/providers/yahoo"
	_ "github.com/zquestz/s/providers/yandex"
	_ "github.com/zquestz/s/providers/youtube"

	"github.com/zquestz/s/cmd"
)

func main() {
	setupSignalHandlers()

	if err := cmd.SearchCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
