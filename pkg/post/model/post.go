package model

type Post struct {
	//Id          primitive.ObjectID `json:"id,omitempty"`
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
