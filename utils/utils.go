package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dghubble/oauth1"
	homedir "github.com/mitchellh/go-homedir"

	yaml "gopkg.in/yaml.v2"
)

// Project : Struct
type Project struct {
	ID                 string    `json:"id" yaml:"id"`
	Name               string    `json:"name" yaml:"name"`
	Track              string    `json:"track" yaml:"track"`
	FilterLevel        string    `json:"filter-level" yaml:"filter-level"`
	Language           string    `json:"language" yaml:"language"`
	Location           string    `json:"location" yaml:"location"`
	Follow             string    `json:"follow" yaml:"follow"`
	StreamingType      string    `json:"streaming-type" yaml:"streaming-type"`
	StallWarnings      bool      `json:"stall-warnings" yaml:"stall-warnings"`
	ConsumerKey        string    `json:"consumer-key" yaml:"consumer-key"`
	ConsumerSecret     string    `json:"consumer-secret" yaml:"consumer-secret"`
	AccessToken        string    `json:"access-token" yaml:"access-token"`
	AccessTokenSecret  string    `json:"access-token-secret" yaml:"access-token-secret"`
	ElasticsearchHost  string    `json:"elasticsearch-host" yaml:"elasticsearch-host"`
	ElasticsearchIndex string    `json:"elasticsearch-index" yaml:"elasticsearch-index"`
	DateCreated        time.Time `json:"date-created" yaml:"date-created"`
	PID                int       `json:"pid" yaml:"pid"`
}

// ExitOnError : Terminate Program with Error
func ExitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ID : Generate Id String
func ID(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

// CreateProject : Save configuration file. Returns *Project or Error
func CreateProject(path string, content []byte) *Project {
	if _, err := os.Stat(path); err == nil {
		err = errors.New("File allready exists")
		ExitOnError(err)
	}

	f, err := os.Create(path)
	ExitOnError(err)

	defer f.Close()

	_, err = f.Write(content)
	ExitOnError(err)

	project := ReadProject(path)

	return project
}

// ReadProject : Read Project configuration from file. Returns *Project
func ReadProject(path string) *Project {
	project := &Project{}
	y, err := ioutil.ReadFile(path)
	ExitOnError(err)

	err = yaml.Unmarshal(y, project)
	ExitOnError(err)

	return project
}

// SaveFile ...
func SaveFile(path string, content []byte) error {
	err := ioutil.WriteFile(path, content, 0644)
	ExitOnError(err)

	return nil
}

// SetHomeDir : Create the .twitterfarm directory under $HOME
func SetHomeDir() {
	home, err := homedir.Dir()
	ExitOnError(err)

	path := home + "/.twitterfarm"
	if _, err = os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	fmt.Printf("twitterfarm configuration folder: %s\n", path)
}

// GetHomeDir :
func GetHomeDir() string {
	home, _ := homedir.Dir()
	return home + "/.twitterfarm"
}

// TwitterConnectionEstablished ...
func TwitterConnectionEstablished(project *Project) bool {
	consumer := oauth1.NewConfig(project.ConsumerKey, project.ConsumerSecret)
	token := oauth1.NewToken(project.AccessToken, project.AccessTokenSecret)
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
