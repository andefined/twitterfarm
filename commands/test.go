package commands

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	elastic "gopkg.in/olivere/elastic.v5"

	"github.com/andefined/twitterfarm/utils"
	"github.com/dghubble/oauth1"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// Test ...
func Test(c *cli.Context) error {
	if c.Args().Get(0) == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}
	home, err := utils.GetHomeDir()
	if err != nil {
		return err
	}

	config := home + "/" + c.Args().Get(0) + ".yml"
	project := utils.ReadFile(config)

	consumer := oauth1.NewConfig(project.ConsumerKey, project.ConsumerSecret)
	token := oauth1.NewToken(project.AccessToken, project.AccessTokenSecret)
	httpClient := consumer.Client(oauth1.NoContext, token)

	isESHost := true
	isESIndex := false

	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(project.ElasticsearchHost))
	if err != nil {
		isESHost = false
	}

	if isESHost {
		isESIndex = true
		exists, err := client.IndexExists(strings.ToLower(project.ElasticsearchIndex)).Do(context.Background())
		if err != nil {
			isESIndex = false
			log.Fatal(err)
			return err
		}
		if !exists {
			isESIndex = false
		}
	}

	if c.Bool("fix") {
		if isESHost && !isESIndex {
			isESIndex = true
			_, err := client.CreateIndex(project.ElasticsearchIndex).Do(ctx)
			if err != nil {
				isESIndex = false
				log.Fatal(err)
				return err
			}
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "NAME", "TWITTER CONNECTION", "ES CONNECTION", "ES INDEX"})
	table.Append([]string{
		project.ID,
		project.Name,
		strconv.FormatBool(utils.TwitterConnectionEstablished(httpClient)),
		strconv.FormatBool(isESHost),
		strconv.FormatBool(isESIndex),
	})

	table.Render()

	return nil
}
