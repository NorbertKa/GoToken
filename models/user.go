package db

type Userprofile struct {
	Id          int    `json:"id"`
	Created     string `json:"created"`
	LastChanges string `json:"lastChanges"`
	LastLogin   string `json:"lastLogin"`
	Identifier  string `json:"identifier"`
	Password    string `json:"password"`
}

