package UserModels

type User struct {
	Id       *int64 `gorm:"primaryKey" json:"id,omitempty"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `gorm:"type:varchar(200)" json:"email"`
	Password string `gorm:"type:varchar(200)" json:"password"`
	FotoURL  string `gorm:"type:varchar(200)" json:"foto_url"`
}

type PaginatedResponse struct {
	List []User `json:"list"`
	Meta struct {
		Total int64 `json:"total"`
	} `json:"meta"`
}
