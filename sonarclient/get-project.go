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

// api/user_groups/search
func (c *SonarQubeClient) GetGroups() (results *SonarUsers) {
	var path string
	path = fmt.Sprintf("/api/user_groups/search")

	results = &SonarUsers{}
	c.doRequest("GET", path, nil, results)
	return results
}

func (c *SonarQubeClient) GetGroupMembers(id int) (results *SonarGroupMembers) {
	var path string
	path = fmt.Sprintf("/api/user_groups/users?id=%v", id)

	results = &SonarGroupMembers{}
	c.doRequest("GET", path, nil, results)
	return results
}

