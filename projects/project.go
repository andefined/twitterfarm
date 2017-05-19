package projects

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/andefined/twitterfarm/utils"
	"github.com/dghubble/oauth1"

	elastic "gopkg.in/olivere/elastic.v5"
	yaml "gopkg.in/yaml.v2"
)

// Project : Struct
type Project struct {
	ID                 string    `yaml:"id"`
	Name               string    `yaml:"name"`
	Track              string    `yaml:"track"`
	FilterLevel        string    `yaml:"filter-level"`
	Language           string    `yaml:"language"`
	Location           string    `yaml:"location"`
	Follow             string    `yaml:"follow"`
	StreamingType      string    `yaml:"streaming-type"`
	StallWarnings      bool      `yaml:"stall-warnings"`
	ConsumerKey        string    `yaml:"consumer-key"`
	ConsumerSecret     string    `yaml:"consumer-secret"`
	AccessToken        string    `yaml:"access-token"`
	AccessTokenSecret  string    `yaml:"access-token-secret"`
	ElasticsearchHost  string    `yaml:"elasticsearch-host"`
	ElasticsearchIndex string    `yaml:"elasticsearch-index"`
	DateCreated        time.Time `yaml:"date-created"`
	PID                int       `yaml:"pid"`
	Config             string    `yaml:"config"`
}

// SetStatic ...
func (p *Project) SetStatic() {
	p.SetID()

	p.Follow = ""
	p.StreamingType = "filter"
	p.StallWarnings = false
	p.DateCreated = time.Now()
	p.PID = 0

	if p.Name == "" {
		p.Name = p.ID
	}

	if p.ElasticsearchIndex == "" {
		p.ElasticsearchIndex = strings.ToLower("twitterfarm_" + p.ID)
	}
}

// SetID ...
func (p *Project) SetID() {
	b := make([]byte, 5)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	p.ID = fmt.Sprintf("%X", b)
}

// Init ...
func (p *Project) Init(config string) {
	if config != "" {
		configFile, err := ioutil.ReadFile(config)
		utils.ExitOnError(err)

		err = yaml.Unmarshal(configFile, &p)
		utils.ExitOnError(err)
	}

	p.SetStatic()

	projectConfigPath := utils.GetHomeDir() + "/" + p.ID + ".yml"
	p.Config = projectConfigPath

	if _, err := os.Stat(projectConfigPath); err == nil {
		utils.ExitOnError(errors.New("File allready exists"))
	}

	f, err := os.Create(projectConfigPath)
	utils.ExitOnError(err)

	defer f.Close()

	data, err := yaml.Marshal(p)
	utils.ExitOnError(err)

	// _, err = f.Write(data)
	err = ioutil.WriteFile(projectConfigPath, data, 0644)
	utils.ExitOnError(err)
}

// Save ...
func (p *Project) Save(path string) {
	data, err := yaml.Marshal(&p)
	utils.ExitOnError(err)

	err = ioutil.WriteFile(path, data, 0644)
	utils.ExitOnError(err)
}

// Read ...
func (p *Project) Read(path string) {
	configFile, err := ioutil.ReadFile(path)
	utils.ExitOnError(err)

	err = yaml.Unmarshal(configFile, &p)
	utils.ExitOnError(err)
}

// TestElasticsearch : Test if there is a living connection for elasticsearch && if index created
func (p *Project) TestElasticsearch(create bool) (bool, bool) {
	isHost := false
	isIndex := false

	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(p.ElasticsearchHost))
	if err != nil {
		return false, false
	}
	isHost = client.IsRunning()

	if isHost {
		exists, err := client.IndexExists(p.ElasticsearchIndex).Do(ctx)
		if err != nil {
			utils.ExitOnError(err)
		}
		if exists {
			isIndex = true
		}

		if !exists && create {
			_, err := client.CreateIndex(p.ElasticsearchIndex).Do(ctx)
			if err != nil {
				isIndex = false
				utils.ExitOnError(err)
			}
			isIndex = false
		}
	}

	client.Stop()
	return isHost, isIndex
}

// TestTwitter : ...
func (p *Project) TestTwitter() bool {
	consumer := oauth1.NewConfig(p.ConsumerKey, p.ConsumerSecret)
	token := oauth1.NewToken(p.AccessToken, p.AccessTokenSecret)
	httpClient := consumer.Client(oauth1.NoContext, token)
	resp, err := httpClient.Get("https://api.twitter.com/1.1/search/tweets.json")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return false
	}

	return true
}
