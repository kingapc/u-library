package model

type TokenDetails struct {
	AccessToken string
	//RefreshToken string
	AccessUuid string
	//RefreshUuid  string
	AtExpires int64
	//RtExpires    int64
}

type AccessDetails struct {
	AccessUuid string `json:"accessuuid"`
	UserId     string `json:"userid"`
	MyId       string `json:"myid"`
	RoleId     string `json:"role_id"`
}
