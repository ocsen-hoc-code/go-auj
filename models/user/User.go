package user

type User struct {
	ID       string `form:"id" json:"id" xml:"id"`
	UserName string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
