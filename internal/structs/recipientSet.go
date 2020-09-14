package structs

type RecipientSet struct {
	Congressmen      string `json:"congressmen"`
	HasCongressmen   bool   `json:"hasCongressmen"`
	Congresswomen    string `json:"congresswomen"`
	HasCongresswomen bool   `json:"hasCongresswomen"`
}
