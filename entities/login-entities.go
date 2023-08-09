package entities

type Login struct {
	Username  string `json:"username" `
	Password  string `json:"password" `
	Ipaddress string `json:"ipaddress" `
	Timezone  string `json:"timezone" `
}
type Home struct {
	Page string `json:"page"`
}
