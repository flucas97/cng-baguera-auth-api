package cannabis

type Cannabis struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Thc          string `json:"thc"`
	Cbd          string `json:"cbd"`
	Weight       string `json:"weight"`
	Height       string `json:"size"`
	Width        string `json:"width"`
	Volume       string `json:"volume"`
	Image        string `json:"image"`
	RepositoryId int64  `json:"repository_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
