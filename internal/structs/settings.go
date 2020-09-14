package structs

type Settings struct {
	Twitter struct {
		BaseUrl      string `json:"baseUrl"`
		Key          string `json:"key"`
		Secret       string `json:"secret"`
		AccessToken  string `json:"accessToken"`
		AccessSecret string `json:"accessSecret"`
	} `json:"twitter"`
	Message struct {
		MaleMnemonic   string `json:"maleMnemonic"`
		FemaleMnemonic string `json:"femaleMnemonic"`
		Qty            int    `json:"qty"`
	} `json:"message"`
}
