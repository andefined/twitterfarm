package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

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

	/*ctx := context.Background()
	esclient, err := elastic.NewClient(elastic.SetURL(project.ElasticsearchHost))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if !utils.TwitterConnectionEstablished(project) {
		log.Fatal("Unable to connect with Twitter API")
		os.Exit(1)
	}*/

	consumer := oauth1.NewConfig(project.ConsumerKey, project.ConsumerSecret)
	token := oauth1.NewToken(project.AccessToken, project.AccessTokenSecret)
	httpClient := consumer.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		out, _ := json.Marshal(tweet)
		fmt.Println(string(out))
		/*t := tweet
		_, err = esclient.Index().Index(strings.ToLower(project.ElasticsearchIndex)).Type("tweet").BodyJson(t).Refresh("true").Do(ctx)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}*/
	}
	filterParams := &twitter.StreamFilterParams{
		Track:         strings.Split(project.Track, ","),
		StallWarnings: &project.StallWarnings,
	}

	stream, err := client.Streams.Filter(filterParams)
	utils.ExitOnError(err)

	go demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
	// esclient.Stop()
}
