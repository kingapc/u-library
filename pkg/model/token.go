package model

type TokenDetails struct {
	AccessToken string
	AccessUuid  string
	AtExpires   int64
}

type AccessDetails struct {
	AccessUuid string `json:"accessuuid"`
	UserId     string `json:"userid"`
	MyId       string `json:"myid"`
	RoleId     string `json:"role_id"`
}
