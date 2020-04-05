package models

//FormDados dados doformulário
type FormDados struct {
	Tipo     string `json:"tip"`
	Host     string `json:"host"`
	Database string `json:"database"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	VHost    string `json:"vhost"`
}
