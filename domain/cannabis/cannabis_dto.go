package cannabis

type Cannabis struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Thc          float64 `json:"thc"`
	Cbd          float64 `json:"cbd"`
	Weight       float64 `json:"weight"`
	Height       float64 `json:"size"`
	Width        float64 `json:"width"`
	Volume       float64 `json:"volume"`
	Image        []byte  `json:"image"`
	RepositoryId int64   `json:"repository_id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
