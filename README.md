[![Build Status](https://travis-ci.org/andefined/twitterfarm.svg?branch=master)](https://travis-ci.org/andefined/twitterfarm)
[![Go Report Card](https://goreportcard.com/badge/github.com/andefined/twitterfarm)](https://goreportcard.com/report/github.com/andefined/twitterfarm)

**Notice**: WIP

# twitterfarm
twitterfarm is a Twitter CLI tool written in [Go](https://golang.org/). The goal is to collect and store data from Twitter Streaming API into an Elasticsearch index fast and easy. Before you begin you must have Elasticsearch up & running and Twitter Application keys/secrets.

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
   $VERSION

COMMANDS:
     init     Initialize twitterfarm. Will create a folder under $HOME/.twitterfarm
     create   Create a new project
     list     List all projects
     test     Test project configuration
     rm       Remove a project
     start    Start a project
     stop     Stop a project
     restart  Restart a project
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
You can create a project either by using the `--config` flag to load your custom [configuration](config/default.yml) or by providing individual flags. If successfully created it will return the project **ID**.

```bash
NAME:
   twitterfarm create - Create a new project

USAGE:
   twitterfarm create [command options]  

OPTIONS:
   --config value               Path to your .yml configuration file
   --name value                 Project Name
   --track value                Tracking Keywords
   --filter-level value         Filter Level (default: "none")
   --language value             Language (default: "en")
   --location value             Location
   --consumer-key value         Twitter Consumer Key
   --consumer-secret value      Twitter Consumer Secret
   --access-token value         Twitter Access Token
   --access-token-secret value  Twitter Access Secret
   --elasticsearch-host value   Comma Separated Elasticsearch Hosts
   --elasticsearch-index value  Elasticsearch Index
```

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

#### Sample configuration file

```yaml
# Project Name
name: us2016

# Streaming Specific (https://dev.twitter.com/streaming/overview/request-parameters)
track: trump,the giant douche,hillary,turd sandwich
filter-level: none
language: en
location: -126.56,22.88,-65.04,49.98
follow: ""
streaming-type: filter
stall-warnings: false

# Twitter Application Specific (https://apps.twitter.com/)
# Consumer Key
consumer-key: TWITTER_CONSUMER_KEY
# Consumer Secret
consumer-secret: TWITTER_CONSUMER_SECRET
# Access Token
access-token: TWITTER_ACCESS_TOKEN
 # Access Token Secret
access-token-secret: TWITTER_ACCESS_TOKEN_SECRET

# Elasticsearch Specific (https://www.elastic.co/)
# Elasticsearch Host (ex. http://elastic:changeme@host-a:9200, http://elastic:changeme@host-b:9200)
elasticsearch-host: http://elastic:changeme@localhost:9200
# Elasticsearch Index
elasticsearch-index: twitterfarm_trump_hillary
```

## List your projects
Will return **ID**, **PID**, **STATUS**, **NAME** and **TRACK**.

```bash
NAME:
   twitterfarm list - List all projects

USAGE:
   twitterfarm list [command options]  

OPTIONS:
   --quiet, -q  Print only ID
```

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

## Test a project
Test command is very useful for testing the connection with Twitter Streaming API and you Elasticsearch. By providing the `--create-index` flag it will create the corresponding Elasticsearch Index (fi not exists).

```bash
twitterfarm test $PROJECTID

ID           | TWITTER API  | ELASTIC HOST | ELASTIC INDEX
             -              -              -
28342C7FF7   | true         | false        | false
```

## Start a project
```bash
twitterfarm start $PROJECTID
```

## Restart a project
```bash
twitterfarm restart $PROJECTID
```

## Stop a project
```bash
twitterfarm stop $PROJECTID
```

## Remove a project
```bash
twitterfarm rm $PROJECTID
```

#### Useful Examples
Stop and remove all projects

```bash
twitterfarm stop $(twitterfarm list -q)
twitterfarm rm $(twitterfarm list -q)
```

## Contributing
1. Fork it
2. Create your feature branch: `git checkout -b feature-name`
3. Commit your changes: `git commit -am feature-name`
4. Push to the branch: `git push origin feature-name`
5. Submit a pull request

## Credits
- [urfave/cli](github.com/urfave/cli)
- [olivere/elastic.v5](gopkg.in/olivere/elastic.v5)
- [dghubble/go-twitter](github.com/dghubble/go-twitter)
- [dghubble/oauth1](github.com/dghubble/oauth1)
- [mitchellh/go-homedir](github.com/mitchellh/go-homedir)
