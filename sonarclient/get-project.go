package sonarclient

import (
	"fmt"
)


// .../api/projects/search_my_projects?ps=
func (c *SonarQubeClient) GetProjects(limit int) (results *SonarProjects) {
	var path string
	path = fmt.Sprintf("/api/projects/search_my_projects?ps=%v", limit)

	results = &SonarProjects{}
	c.doRequest("GET", path, nil, results)
	return results
}

