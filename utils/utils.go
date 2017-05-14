package utils

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dghubble/oauth1"
	homedir "github.com/mitchellh/go-homedir"

	yaml "gopkg.in/yaml.v2"
)

// Project ...
type Project struct {
	ID                 string    `json:"id" yaml:"id"`
	Name               string    `json:"name" yaml:"name"`
	ConsumerKey        string    `json:"consumer-key" yaml:"consumer-key"`
	ConsumerSecret     string    `json:"consumer-secret" yaml:"consumer-secret"`
	AccessToken        string    `json:"access-token" yaml:"access-token"`
	AccessTokenSecret  string    `json:"access-token-secret" yaml:"access-token-secret"`
	ElasticsearchHost  string    `json:"elasticsearch-host" yaml:"elasticsearch-host"`
	ElasticsearchIndex string    `json:"elasticsearch-index" yaml:"elasticsearch-index"`
	Keywords           string    `json:"keyword" yaml:"keyword"`
	DateCreated        time.Time `json:"date-created" yaml:"date-created"`
	PID                int       `json:"pid" yaml:"pid"`
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

// CreateProject : Save configuration file
func CreateProject(path string, content []byte) (*Project, error) {
	if _, err := os.Stat(path); err == nil {
		return nil, err
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return nil, err
	}

	project := ReadProject(path)

	return project, nil
}

// ReadProject ...
func ReadProject(path string) *Project {
	project := &Project{}
	y, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(y, project)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return project
}

// SaveFile ...
func SaveFile(path string, content []byte) error {
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

// GetHomeDir : Return $HOME Directory or Error
func GetHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	path := home + "/.twitterfarm"
	if _, err = os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	return path
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
