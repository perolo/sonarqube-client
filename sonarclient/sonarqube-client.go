package sonarclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//ConfluenceClient is the primary client to the Confluence API
type SonarQubeClient struct {
	username string
	password string
	baseURL  string
	debug    bool
	client   *http.Client
}


//Client returns a new instance of the client
func Client(config *SonarQubeConfig) *SonarQubeClient {
	return &SonarQubeClient{
		username: config.Username,
		password: config.Password,
		baseURL:  config.URL,
		debug:    config.Debug,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (c *SonarQubeClient) doRequest(method, url string, content, responseContainer interface{}) ([]byte,  *http.Response){
	b := new(bytes.Buffer)
	if content != nil {
		json.NewEncoder(b).Encode(content)
	}
	furl := c.baseURL + url
	if c.debug {
		log.Println("Full URL", furl)
		log.Println("JSON Content:", b.String())
	}
	request, err := http.NewRequest(method, furl, b)
	request.SetBasicAuth(c.username, c.password)
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		log.Fatal(err)
	}
	if c.debug {
		log.Println("Sending request to services...")
	}
	response, err := c.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if c.debug {
		log.Println("Response received, processing response...")
		log.Println("Response status code is", response.StatusCode)
		log.Println(response.Status)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode < 200 || response.StatusCode > 300 {
		log.Println("Bad response code received from server: ", response.Status)
	} else {
		json.Unmarshal(contents, responseContainer)
	}
	return contents, response
}


func (c *SonarQubeClient) doGetPage(method, url string, content interface{}) ([]byte,  *http.Response){
	b := new(bytes.Buffer)
	if content != nil {
		json.NewEncoder(b).Encode(content)
	}
	furl := c.baseURL + url
	if c.debug {
		log.Println("Full URL", furl)
		log.Println("JSON Content:", b.String())
	}
	request, err := http.NewRequest(method, furl, b)
	request.SetBasicAuth(c.username, c.password)
	if c.debug {
		log.Println("Sending request to services...")
	}
	response, err := c.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if c.debug {
		log.Println("Response received, processing response...")
		log.Println("Response status code is", response.StatusCode)
		log.Println(response.Status)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode < 200 || response.StatusCode > 300 {
		log.Println("Bad response code received from server: ", response.Status)
	}
	return contents, response
}
