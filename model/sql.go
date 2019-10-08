package model

//获取轮播图信息
func Lun() ([]LunImgs, error) {
	mod := make([]LunImgs, 0)
	err := DB.Select(&mod, `select * from lunImgs`)
	return mod, err
}

//用户注册
func UserRegister(truename, pass, email, tel string) error {
	_, err := DB.Exec(`insert into user(trueName, pass, email, tel) values(?,?,?,?)`, truename, pass, email, tel)
	return err
}

//忘记密码
func LostPass(pass, passagain, email string) error {
	_, err := DB.Exec(`update user set pass=?, passagain =? where email=?`, pass, passagain, email)
	return err
}

//用户登录
func Login(tel string) (User, error) {
	mod := User{}
	err := DB.Get(&mod, `select * from user where tel=?`)
	return mod, err
}

//获取所有课程
func GetBookList() ([]Book, error) {
	mod := make([]Book, 0)
	err := DB.Select(&mod, `select * from book`)
	return mod, err
}

//查询用户购物车信息
func UserCart(userid int) (Cart, error) {
	mod := Cart{}
	err := DB.Get(&mod, `select * from cart where userID=?`, userid)
	return mod, err
}

//查询用户地址信息
func UserAddress(userid int) (Address, error) {
	mod := Address{}
	err := DB.Get(&mod, `select * from address where =?`, userid)
	return mod, err
}

//查询用户关注信息
func UserGZ(userid int) (SaveGoods, error) {
	mod := SaveGoods{}
	err := DB.Get(&mod, `select * from saveGoods where userID=?`, userid)
	return mod, err
}

//查询用户已经购买信息
func UserHaveBuy(userid int) (HaveBuy, error) {
	mod := HaveBuy{}
	err := DB.Get(&mod, `select * from saveGoods where userID=?`, userid)
	return mod, err
}
