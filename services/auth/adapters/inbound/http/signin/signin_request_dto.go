package signinhttpinboundadapter

type signInRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
