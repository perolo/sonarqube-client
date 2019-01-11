package main

import (
	"flag"
	"github.com/perolo/sonarqube-client/sonarclient"
)

var config = sonarclient.SonarQubeConfig{}


func main() {
	flag.StringVar(&config.Username, "u", "", "Confluence username")
	flag.StringVar(&config.Password, "p", "", "Confluence password")
	flag.StringVar(&config.URL, "s", "", "The base URL of the Confluence page")
	flag.Parse()
}

