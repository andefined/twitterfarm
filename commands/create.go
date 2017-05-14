package commands

import (
	"fmt"
	"log"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/urfave/cli"

	utils "github.com/andefined/twitterfarm/utils"
)

// Create ...
func Create(c *cli.Context) error {
	p := utils.Project{
		ID:                 utils.ID(5),
		Name:               c.String("name"),
		ConsumerKey:        c.String("consumer-key"),
		ConsumerSecret:     c.String("consumer-secret"),
		AccessToken:        c.String("access-token"),
		AccessTokenSecret:  c.String("access-token-secret"),
		ElasticsearchHost:  c.String("elasticsearch-host"),
		ElasticsearchIndex: c.String("elasticsearch-index"),
		Keywords:           c.String("keywords"),
		DateCreated:        time.Now(),
		PID:                0,
	}

	if p.ConsumerKey == "" ||
		p.ConsumerSecret == "" ||
		p.AccessToken == "" ||
		p.AccessTokenSecret == "" ||
		p.ElasticsearchHost == "" ||
		p.Keywords == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	if p.Name == "" {
		p.Name = p.ID
	}

	if p.ElasticsearchIndex == "" {
		p.ElasticsearchIndex = strings.ToLower("twitterfarm_" + p.ID)
	}

	y, err := yaml.Marshal(p)
	if err != nil {
		log.Print(err)
		return err
	}

	config := utils.GetHomeDir() + "/" + p.ID + ".yml"
	project, _ := utils.CreateProject(config, y)
	fmt.Printf("%s\n", project.ID)
	return nil
}
