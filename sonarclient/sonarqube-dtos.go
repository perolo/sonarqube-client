package sonarclient

type SonarProjects struct {
	Paging struct {
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		Total     int `json:"total"`
	} `json:"paging"`
	Projects []struct {
		ID               string `json:"id"`
		Key              string `json:"key"`
		Name             string `json:"name"`
		Description      string `json:"description,omitempty"`
		LastAnalysisDate string `json:"lastAnalysisDate,omitempty"`
		QualityGate      string `json:"qualityGate,omitempty"`
		Links            []struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Href string `json:"href"`
		} `json:"links"`
	} `json:"projects"`
}


type ProjectStatusType struct {
	ProjectStatus struct {
		Status     string `json:"status"`
		Conditions []struct {
			Status           string `json:"status"`
			MetricKey        string `json:"metricKey"`
			Comparator       string `json:"comparator"`
			WarningThreshold string `json:"warningThreshold,omitempty"`
			ActualValue      string `json:"actualValue"`
			PeriodIndex      int    `json:"periodIndex,omitempty"`
			ErrorThreshold   string `json:"errorThreshold,omitempty"`
		} `json:"conditions"`
		Periods []struct {
			Index     int    `json:"index"`
			Mode      string `json:"mode"`
			Date      string `json:"date"`
			Parameter string `json:"parameter"`
		} `json:"periods"`
		IgnoredConditions bool `json:"ignoredConditions"`
	} `json:"projectStatus"`
}

type QualityGateType struct {
	QualityGate struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Default bool   `json:"default"`
	} `json:"qualityGate"`
}