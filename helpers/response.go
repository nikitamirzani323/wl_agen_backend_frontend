package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
	Time    string      `json:"time"`
}
type Responsemovie struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Perpage     int         `json:"perpage"`
	Totalrecord int         `json:"totalrecord"`
	Time        string      `json:"time"`
}
type Responsepaging struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Record      interface{} `json:"record"`
	Perpage     int         `json:"perpage"`
	Totalrecord int         `json:"totalrecord"`
	Time        string      `json:"time"`
}
type ResponseAdmin struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Record   interface{} `json:"record"`
	Listrule interface{} `json:"listruleadmin"`
	Time     string      `json:"time"`
}
type ErrorResponse struct {
	Field string
	Tag   string
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
