package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	elastic "gopkg.in/olivere/elastic.v5"

	"github.com/andefined/twitterfarm/projects"
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
	// Limit counter
	limit := -1

	path := c.Args().Get(0)
	// Create a temp project
	project := &projects.Project{}
	// Assign values from file
	project.Read(path)

	// A Context carries a deadline, a cancelation signal, and other values across API boundaries.
	ctx := context.Background()
	// Elasticsearch client
	esClient, err := elastic.NewClient(elastic.SetURL(project.ElasticsearchHost))
	utils.ExitOnError(err)

	// Http client & authentication
	consumer := oauth1.NewConfig(project.ConsumerKey, project.ConsumerSecret)
	token := oauth1.NewToken(project.AccessToken, project.AccessTokenSecret)
	httpClient := consumer.Client(oauth1.NoContext, token)

	// Twitter client
	twitterClient := twitter.NewClient(httpClient)

	// Demux Listener
	demux := twitter.NewSwitchDemux()
	// On Tweet
	demux.Tweet = func(tweet *twitter.Tweet) {
		// For debugging ONLY. It is not a good idea to log in your console large json files.
		// Use it with caution ex. `twitterfarm exec {{PROJECT_ID}}.yml --drop --verbose --limit 10`.
		if c.Bool("verbose") {
			out, _ := json.Marshal(tweet)
			fmt.Println(string(out))
		}

		// Skip indexing when `--drop` flag exists.
		if !c.Bool("drop") {
			// Save the tweet (should i ignore the error?).
			_, err = esClient.Index().
				Index(project.ElasticsearchIndex).
				Type("tweet").
				BodyJson(tweet).
				Refresh("true").
				Do(ctx)
			utils.ExitOnError(err)
		}

		// If certain limit reached will Exit
		if c.Int("limit") > -1 && c.Int("limit") >= limit {
			err = errors.New("Limit reached, exiting")
			utils.ExitOnError(err)
		}

		// Just increase the counter but only if `--limit` flag is passed.
		// In general this number can become a really very big number which is a bad practice for production.
		// Use it ONLY for debugging.
		if c.Int("limit") > -1 {
			limit++
		}
	}

	// Stream filter
	filterParams := &twitter.StreamFilterParams{
		Track:         strings.Split(project.Track, ","),
		StallWarnings: &project.StallWarnings,
		Locations:     strings.Split(project.Location, ","),
	}

	// Streamer
	stream, err := twitterClient.Streams.Filter(filterParams)
	utils.ExitOnError(err)

	// Tweets channel
	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")

	// Stop the stream
	stream.Stop()
	// Stop elasticsearch client
	esClient.Stop()
}
