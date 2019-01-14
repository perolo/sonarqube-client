# sonarqube-client

Simple REST API Wrapper for SonarQube

## Features

* Authentication 
* GetProjects
* GetQualityGateByProject
* GetProjectStatus

## Compatible SonarQube versions

This package was tested against 6.7.6 (build 38781) 

## Installation

It is go gettable

    $ go get github.com/perolo/sonarqube-client


## API

The [latest SonarQube documentation](https://docs.sonarqube.org/latest/) was the base document for this package.


### Login

```go
var sonarClient *sonarclient.SonarQubeClient

var sonaConfig = sonarclient.SonarQubeConfig{}
sonaConfig.Username = "User"
sonaConfig.Password = "Pass"
sonaConfig.URL = "https://SomeURL"

sonarClient = sonarclient.Client(&sonaConfig)
```

### Get Projects

```go
projects := sonarClient.GetProjects(20)
fmt.Printf("ProjCount: %v\n", projects.Paging.Total)
```


## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).