package commons

type DataBase struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	DBName       string `json:"dbName"`
	Password     string `json:"password"`
	MaxOpenConns int    `json:"maxOpenConns"`
}
