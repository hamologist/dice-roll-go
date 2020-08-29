package model

type (
	RollResponse struct {
		Step []Step `json:"step"`
	}

	Step struct {
		Rolls []Roll `json:"rolls"`
		Total int    `json:"total"`
	}

	Roll struct {
		Count    int   `json:"count"`
		Sides    int   `json:"sides"`
		Modifier int   `json:"modifier"`
		Rolls    []int `json:"rolls"`
		Total    int   `json:"total"`
	}
)
