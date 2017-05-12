package commands

import (
	"fmt"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/urfave/cli"

	utils "github.com/andefined/twitterfarm/utils"
)

// Create ...
func Create(c *cli.Context) error {
	project := utils.Project{
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
	if project.Name == "" {
		project.Name = project.ID
	}

	if project.ConsumerKey == "" || project.ConsumerSecret == "" || project.AccessToken == "" || project.AccessTokenSecret == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	if project.ElasticsearchHost == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	if project.ElasticsearchIndex == "" {
		project.ElasticsearchIndex = "twitterfarm" + "-" + project.Name + "-" + project.ID
	}

	if project.Keywords == "" {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	y, err := yaml.Marshal(project)
	if err != nil {
		return err
	}

	home, err := utils.GetHomeDir()
	if err != nil {
		return err
	}

	config := home + "/" + project.ID + ".yml"
	if _, err = os.Stat(config); err == nil {
		return err
	}

	err = utils.CreateFile(config, y)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", project.ID)
	return nil
}
