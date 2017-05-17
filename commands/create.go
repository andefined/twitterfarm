package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/andefined/twitterfarm/utils"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

func createWithFile(c *cli.Context) *utils.Project {
	p := utils.ReadProject(c.String("config"))
	p.ID = utils.ID(5)
	p.Follow = ""
	p.StreamingType = "filter"
	p.StallWarnings = false
	p.DateCreated = time.Now()
	p.PID = 0

	return p
}

func createWithFlags(c *cli.Context) *utils.Project {
	p := &utils.Project{
		ID:                 utils.ID(5),
		Name:               c.String("name"),
		Track:              c.String("track"),
		FilterLevel:        c.String("filter-level"),
		Language:           c.String("language"),
		Location:           c.String("location"),
		Follow:             "",
		StreamingType:      "filter",
		StallWarnings:      false,
		ConsumerKey:        c.String("consumer-key"),
		ConsumerSecret:     c.String("consumer-secret"),
		AccessToken:        c.String("access-token"),
		AccessTokenSecret:  c.String("access-token-secret"),
		ElasticsearchHost:  c.String("elasticsearch-host"),
		ElasticsearchIndex: c.String("elasticsearch-index"),
		DateCreated:        time.Now(),
		PID:                0,
	}

	return p
}

// Create : Create a new project
func Create(c *cli.Context) error {
	var p *utils.Project

	if c.String("config") != "" {
		p = createWithFile(c)
	} else {
		p = createWithFlags(c)
	}

	if p.ConsumerKey == "" ||
		p.ConsumerSecret == "" ||
		p.AccessToken == "" ||
		p.AccessTokenSecret == "" ||
		p.ElasticsearchHost == "" ||
		p.Track == "" {
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
	utils.ExitOnError(err)

	config := utils.GetHomeDir() + "/" + p.ID + ".yml"
	project := utils.CreateProject(config, y)
	fmt.Printf("%s\n", project.ID)

	return nil
}
