package commands

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/andefined/twitterfarm/utils"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/urfave/cli"
)

// Exec ...
func Exec(c *cli.Context) {
	if c.Args().Get(0) == "" {
		cli.ShowSubcommandHelp(c)
	}
	config := c.Args().Get(0)
	project := utils.ReadFile(config)

	consumer := oauth1.NewConfig(project.ConsumerKey, project.ConsumerSecret)
	token := oauth1.NewToken(project.AccessToken, project.AccessTokenSecret)
	httpClient := consumer.Client(oauth1.NoContext, token)

	if !TwitterConnectionEstablished(httpClient) {
		log.Fatal("Unable to connect with Twitter Streaming API")
		os.Exit(1)
	}
	// Twitter client
	client := twitter.NewClient(httpClient)
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		// out, _ := json.Marshal(tweet)
		// fmt.Println(string(out))
	}
	filterParams := &twitter.StreamFilterParams{
		Track:         strings.Split(project.Keywords, ","),
		StallWarnings: twitter.Bool(true),
	}

	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}

// TwitterConnectionEstablished ...
func TwitterConnectionEstablished(httpClient *http.Client) bool {
	resp, err := httpClient.Get("https://api.twitter.com/1.1/search/tweets.json")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return false
	}

	return true
}
