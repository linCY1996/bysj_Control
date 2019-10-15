package model

//获取轮播图信息
func Lun() ([]LunImgs, error) {
	mod := make([]LunImgs, 0)
	err := DB.Select(&mod, `select * from lunImgs`)
	return mod, err
}

//用户注册
func UserRegister(truename, pass, email, tel string) error {
	_, err := DB.Exec(`insert into user(uname, pass, email, tel, status) values(?,?,?,?,0)`, truename, pass, email, tel)
	return err
}

//忘记密码
func LostPass(pass, email string) error {
	_, err := DB.Exec(`update user set pass=? where email=?`, pass, email)
	return err
}

//首先通过邮箱查找是否有此用户
func EmailLook(email string) (User, error) {
	mod := User{}
	err := DB.Get(&mod, `select * from user where email=?`, email)
	return mod, err
}

//修改登录状态
func UndateLoginStatus(status int) error {
	_, err := DB.Exec(`update user set status=?`, status)
	return err
}

//用户登录
func Login(tel string) (User, error) {
	mod := User{}
	err := DB.Get(&mod, `select * from user where tel=?`, tel)
	return mod, err
}

//通过id获取用户信息
func LookUserMsgs(uid int) (User, error) {
	mod := User{}
	err := DB.Get(&mod, `select * from user where id=?`, uid)
	return mod, err
}

//获取所有课程
func GetBookList() ([]Book, error) {
	mod := make([]Book, 0)
	err := DB.Select(&mod, `select * from book`)
	return mod, err
}

//加入购物车
func AddCart(bookid, uid int) error {
	_, err := DB.Exec(`insert into cart(bookid, userid, booknum, bookstatus) values(?,?,1,0)`, bookid, uid)
	return err
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

//获取首页子级入口展示信息
func GetLabels() ([]GetLabel, error) {
	mod := make([]GetLabel, 0)
	err := DB.Select(&mod, `select * from indexchr`)
	return mod, err
}

//首页进入子入口查询对应得课程
func LookdetailClass(class int) ([]Book, error) {
	mod := make([]Book, 0)
	err := DB.Select(&mod, `select * from book where jpclass=?`, class)
	return mod, err
}

//查询商品详情信息
func LookGoodsDetailMsgs(gid int) (Book, error) {
	mod := Book{}
	err := DB.Get(&mod, `select * from book where id=?`, gid)
	return mod, err
}

//查看详情页底部多图展示
func ShowDetailImgsArray(gid int) (Bookdetailimgs, error) {
	mod := Bookdetailimgs{}
	err := DB.Get(&mod, `select * from bookdetailimgs where bookid=?`, gid)
	return mod, err
}

//  收藏/删除商品
//接收status      0：取消收藏       1：收藏
func ControlGoods(userid, goodsid, status int, time string) error {
	if status == 0 {
		_, err := DB.Exec(`delete from savegoods where goodsid=? and userid = ?`, goodsid, userid)
		return err
	} else {
		_, err := DB.Exec(`insert into savegoods(goodsid, userid, time) values(?, ?, ?)`, goodsid, userid, time)
		return err
	}

}

//通过用户id去查询该商品是否收藏
func LookGoodsIsSave(uid int) ([]SaveGoods, error) {
	mod := make([]SaveGoods, 0)
	err := DB.Select(&mod, `select * from savegoods where userid=?`, uid)
	return mod, err
}

//获取指定得商品得全部评论
func GetDetailCommand(gid int) ([]Command, error) {
	mod := make([]Command, 0)
	err := DB.Select(&mod, `select * from command where goodsid=?`, gid)
	return mod, err
}

//获取分类栏目名称
func SlideType() ([]Slidebar, error) {
	mod := make([]Slidebar, 0)
	err := DB.Select(&mod, `select * from slidebar`)
	return mod, err
}
