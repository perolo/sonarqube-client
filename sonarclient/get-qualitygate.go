package sonarclient

import (
	"fmt"
)



//https://sonar.aa.st/sonarqube/api/qualitygates/get_by_project?project=com.assaabloy.stg:a3p

func (c *SonarQubeClient) GetQualityGateByProject(key string) (results *QualityGateType) {
	var path string
	path = fmt.Sprintf("/api/qualitygates/project_status?projectKey=%s", key)

	results = &QualityGateType{}
	c.doRequest("GET", path, nil, results)
	return results
}

//https://sonar.aa.st/sonarqube/api/qualitygates/project_status?projectKey=com.assaabloy.stg:a3p
func (c *SonarQubeClient) GetProjectStatus(key string) (results *ProjectStatusType) {
	var path string
	path = fmt.Sprintf("/api/qualitygates/project_status?projectKey=%s", key)

	results = &ProjectStatusType{}
	c.doRequest("GET", path, nil, results)
	return results
}
