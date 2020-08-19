package account

type Account struct {
	ID                   int64    `json:"id"`
	Name                 string   `json:"name"`
	Email                string   `json:"email"`
	Password             string   `json:"password"`
	City                 string   `json:"city"`
	Country              string   `json:"country"`
	State                string   `json:"state"`
	AvaliableFeatures    []string `json:"avaliable_features"`
	CannabisRepositoryID string   `json:"cannabis_repository_id"`
	Description          string   `json:"description"`
	Language             string   `json:"language"`
	Status               string   `json:"status"`
	CreatedAt            string   `json:"created_at"`
	UpdatedAt            string   `json:"updated_at"`
}
