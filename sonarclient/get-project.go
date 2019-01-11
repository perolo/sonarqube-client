package sonarclient

import (
	"fmt"
)




func (c *SonarQubeClient) GetProjects() (results *SonarProjects) {
	var path string
	path = fmt.Sprintf("/api/projects/search_my_projects?ps=250")

	results = &SonarProjects{}
	c.doRequest("GET", path, nil, results)
	return results
}

