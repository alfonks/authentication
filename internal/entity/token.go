package entity

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (t *Tokens) IsEmpty() bool {
	return t.AccessToken == "" || t.RefreshToken == ""
}
