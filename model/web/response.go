package web

type CatResp struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Phone    int64  `json:"phone"`
	Role     string `json:"role"`
}
