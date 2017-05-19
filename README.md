[![Build Status](https://travis-ci.org/andefined/twitterfarm.svg?branch=master)](https://travis-ci.org/andefined/twitterfarm)
[![Go Report Card](https://goreportcard.com/badge/github.com/andefined/twitterfarm)](https://goreportcard.com/report/github.com/andefined/twitterfarm)

# twitterfarm
twitterfarm is a Twitter CLI tool written in [Go](https://golang.org/). The goal is to collect and store data from Twitter Streaming API into an Elasticsearch index fast and easy. Before you begin you must have Elasticsearch up & running and Twitter Application keys/secrets.

- [Installing Elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/5.x/install-elasticsearch.html)
- [Twitter Applications](https://apps.twitter.com/)


## Installation
You can download the binaries from the [releases](/releases) section, or you can install it with Go.

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
To start using twitterfarm you need to create a working directory. The command simple creates a folder under `$HOME/.twitterfarm` where we store the configuration files for every project and needs to be run only once.

```bash
twitterfarm init
```

## Create a project
You can create a project either by using the `--config` flag to load your custom [configuration](config/default.yml) or by providing individual flags. If successfully created it will return the project **ID**.

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
Test command is very useful for testing the connection with Twitter Streaming API and your Elasticsearch. By providing the `--create-index` flag it will create the corresponding Elasticsearch Index (fi not exists).

```bash
twitterfarm test $PROJECT_ID

ID           | TWITTER API  | ELASTIC HOST | ELASTIC INDEX
             -              -              -
28342C7FF7   | true         | false        | false
```

## Start a project
```bash
twitterfarm start $PROJECT_ID
```

## Restart a project
```bash
twitterfarm restart $PROJECT_ID
```

## Stop a project
```bash
twitterfarm stop $PROJECT_ID
```

## Remove a project
The command will try to stop the project `proc.Kill()` and then will remove the configuration file from `$HOME/.twitterfarm/{{$PROJECT_ID}}.yml`. It will NOT remove any data from your elasticsearch index.

```bash
twitterfarm rm $PROJECT_ID
```

## TO DO

- [ ] Global Logger & Log Command
- [ ] Twitter Streaming API Request Parameters
- [ ] Twitter Streaming API User/Site
- [ ] Elasticsearch Tweet Mappings
- [ ] Before Save Pipe Script

## Contributing
1. Fork it
2. Create your feature branch: `git checkout -b feature-name`
3. Commit your changes: `git commit -am feature-name`
4. Push to the branch: `git push origin feature-name`
5. Submit a pull request

## Credits
- ![urfave/cli](https://github.com/urfave/cli)
- ![olivere/elastic.v5](https://github.com/olivere/elastic/tree/v5.0.38)
- ![dghubble/go-twitter](https://github.com/dghubble/go-twitter)
- ![dghubble/oauth1](https://github.com/dghubble/oauth1)
- ![mitchellh/go-homedir](https://github.com/mitchellh/go-homedir)
