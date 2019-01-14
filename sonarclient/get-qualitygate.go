package sonarclient

import (
	"fmt"
)



//.../api/qualitygates/get_by_project?project=

func (c *SonarQubeClient) GetQualityGateByProject(key string) (results *QualityGateType) {
	var path string
	path = fmt.Sprintf("/api/qualitygates/get_by_project?project=%s", key)

	results = &QualityGateType{}
	c.doRequest("GET", path, nil, results)
	return results
}

//.../api/qualitygates/project_status?projectKey=
func (c *SonarQubeClient) GetProjectStatus(key string) (results *ProjectStatusType) {
	var path string
	path = fmt.Sprintf("/api/qualitygates/project_status?projectKey=%s", key)

	results = &ProjectStatusType{}
	c.doRequest("GET", path, nil, results)
	return results
}
