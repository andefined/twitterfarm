package commands

import (
	"context"
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

	path := c.Args().Get(0)
	// Create a temp project
	project := &projects.Project{}
	// Assign values from file
	project.Read(path)

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
		// out, _ := json.Marshal(tweet)
		// fmt.Println(string(out))
		_, err = esClient.Index().Index(project.ElasticsearchIndex).Type("tweet").BodyJson(tweet).Refresh("true").Do(ctx)
		utils.ExitOnError(err)
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
