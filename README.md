[![Build Status](https://travis-ci.org/andefined/twitterfarm.svg?branch=master)](https://travis-ci.org/andefined/twitterfarm)
[![Go Report Card](https://goreportcard.com/badge/github.com/andefined/twitterfarm)](https://goreportcard.com/report/github.com/andefined/twitterfarm)

**Notice**: WIP

# twitterfarm
twitterfarm is a Twitter CLI tool written in [Go](https://golang.org/). The goal is to collect and store data from Twitter Streaming API into an Elasticsearch index fast and easy. Before you begin you must have Elasticsearch up & running and Twiiter Application keys/secrets.

- [Installing Elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/5.x/install-elasticsearch.html)
- [Twitter Applications](https://apps.twitter.com/)


## Installation
With Go
```bash
go install github.com/andefined/twitterfarm
```

## How to use
```
NAME:
   twitterfarm - Collect data from Twitter

USAGE:
   twitterfarm [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     init     Initialize twitterfarm. Will create a folder under $HOME/.twitterfarm
     create   Create a new project
     list     List all projects
     test     Test project configuration
     rm       Remove a project
     start    Start a project
     exec     Execute a project
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
## Initialize twitterfarm
Before you begin you need to initialize twitterfarm for the first time. The command simple creates a folder under `$HOME/.twitterfarm` where we store the configuration files for every project.
```bash
twitterfarm init
```
## Create a project
You can create a project either by using the `--config` flag to load your custom [configuration](config/test.yml) file or by providing indivual flags. If succesfully created it will return the project **ID**.
```bash
twitterfarm create --config config/default.yml
```
OR
```bash
twitterfarm create \
    --name "us2016" \
    --track "trump,the giant douche,hillary,turd sandwich" \
    --elasticsearch-host "http://elastic:changeme@localhost:9200" \
    --elasticsearch-index "twitterfarm_trump_hillary" \
    --consumer-key $TWITTER_CONSUMER_KEY \
    --consumer-secret $TWITTER_CONSUMER_SECRET \
    --access-token $TWITTER_ACCESS_TOKEN \
    --access-token-secret $TWITTER_ACCESS_TOKEN_SECRET
```

## List your projects
Will return **ID**, **PID**, **STATUS**, **NAME** and **TRACK**.
```bash
twitterfarm list

ID         | PID   | STATUS           | NAME             | TRACK
           -       -                  -                  -
215BC106C5 | 0     | not initialized  | us2016           | trump,the giant douch...
247FBA4667 | 0     | not initialized  | us2016           | trump,the giant douch...
```
OR if you want to return only the **ID**.
```bash
twitterfarm list -q

215BC106C5
247FBA4667

```

## Contributing
1. Fork it
2. Create your feature branch: `git checkout -b feature-name`
3. Commit your changes: `git commit -am 'feature-name'`
4. Push to the branch: `git push origin feature-name`
5. Submit a pull request

## Credits
- [urfave/cli](github.com/urfave/cli)
- [olivere/elastic.v5](gopkg.in/olivere/elastic.v5)
- [dghubble/go-twitter](github.com/dghubble/go-twitter)
- [dghubble/oauth1](github.com/dghubble/oauth1)
- [mitchellh/go-homedir](github.com/mitchellh/go-homedir)
