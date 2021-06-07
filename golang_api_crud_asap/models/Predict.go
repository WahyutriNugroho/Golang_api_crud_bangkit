package models

type Predict struct {
	Id          int     `json:id`
	Class_name  string  `json:"class_name"`
	ImgUrl      string  `json:"imgUrl"`
	Date        string  `json:"date"`
	Latitude    string  `json:"latitude"`
	Longtitude  string  `json:"longtitude"`
	Probability float64 `json:"probability"`
}
