[![Build Status](https://travis-ci.org/andefined/twitterfarm.svg?branch=master)](https://travis-ci.org/andefined/twitterfarm)
[![Go Report Card](https://goreportcard.com/badge/github.com/andefined/twitterfarm)](https://goreportcard.com/report/github.com/andefined/twitterfarm)

**Notice**: WIP

# twitterfarm
twitterfarm is a Twitter CLI tool written in Go. The goal is to collect and store data from Twitter Streaming API into an Elasticsearch index fast and easy. Before you begin you must have a working **elasticsearch** cluster and a **Twiiter Application** keys/secrets.

- [Installing Elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/5.x/install-elasticsearch.html)

    The easiest way to test elasticsearch is with docker
    ```bash
    docker pull docker.elastic.co/elasticsearch/elasticsearch:5.4.0
    docker run \
        -p 9200:9200 \
        -e "http.host=0.0.0.0" \
        -e "transport.host=127.0.0.1" \
        docker.elastic.co/elasticsearch/elasticsearch:5.4.0
    ```

- [Twitter Applications](https://apps.twitter.com/)

    Create an application and generate *Access Token*, *Access Secret* for every project you want to run in **twitterfarm**.


#### Installation
With Go
```bash
go install github.com/andefined/twitterfarm
```

#### How to use
```
NAME:
   twitterfarm - Collect data from Twitter

USAGE:
   twitterfarm [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     create   Create a new project
     list     List all projects
     test     Test project configuration: Twitter API Authorization, Elasticsearch Connection, Elasticsearcg Index
     rm       Remove a project
     start    Start a project
     exec     Execute a project
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

##### Create a project
```bash
twitterfarm create \
    --name "Trump" \
    --keywords "Trump, giant douche" \
    --elasticsearch-host "http://elastic:changeme@localhost:9200" \
    --elasticsearch-index "twitterfarm_trump" \
    --consumer-key $TWITTER_CONSUMER_KEY \
    --consumer-secret $TWITTER_CONSUMER_SECRET \
    --access-token $TWITTER_ACCESS_TOKEN \
    --access-token-secret $TWITTER_ACCESS_TOKEN_SECRET
```
