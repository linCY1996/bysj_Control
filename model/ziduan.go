package model

//轮播图
type LunImgs struct {
	ID    int    `json:"id" xml:"id"`       //id
	Imgs  string `json:"imgs" xml:"imgs"`   //图片连接
	Msgs  string `json:"msgs" xml:"msgs"`   //信息介绍
	Links string `json:"links" xml:"links"` //跳转连接
}

//用户信息
type User struct {
	ID        int    `json:"id" xml:"id"`
	TrueName  string `json:"truename" xml:"truename"`   //用户真实姓名
	Tel       string `json:"tel" xml:"tel"`             //电话
	Pass      string `json:"pass" xml:"pass"`           //密码
	Passagain string `json:"passagain" xml:"passagain"` //第二次密码
	NickName  string `json:"nickname" xml:"nickname"`   //昵称
	Sex       string `json:"sex" xml:"sex"`             //性别
	Email     string `json:"email" xml:"email"`         //邮箱
	HeadImgs  string `json:"head_imgs" xml:"head_imgs"` //头像
}

//课程信息
type Book struct {
	ID             int    `json:"id" xml:"id"`
	Bookname       string `json:"bookname" xml:"bookname"`             //课程名
	Bookimgs       string `json:"bookimgs" xml:"bookimgs"`             //课程首图
	Bookprice      string `json:"bookprice" xml:"bookprice"`           //课价格
	Bookmsgs       string `json:"bookmsgs" xml:"bookmsgs"`             //课程信息
	Booksalenum    int    `json:"booksalenum" xml:"booksalenum"`       //课程购买量
	Bookgz         int    `json:"bookgz" xml:"bookgz"`                 //用户关注量
	Bookdetailimgs string `json:"bookdetailimgs" xml:"bookdetailimgs"` //课程详情图
	Bookclass      int    `json:"bookclass" xml:"bookclass"`           //bookClass  课程种类  1：数学类，  2：文学类   3：杂志类   4：艺术类   5：心理学类     6：儿童类     7：漫画类     8：软件编程类    9：金融类
}

//购物车
type Cart struct {
	ID         int `json:"id" xml:"id"`
	Bookid     int `json:"bookid" xml:"bookid"`           //课程id
	Userid     int `json:"userid" xml:"userid"`           //用户id
	Booknum    int `json:"booknum" xml:"booknum"`         //购买量
	BookStatus int `json:"book_status" xml:"book_status"` //购买状态
}

//评论
type Command struct {
	ID          int    `json:"id" xml:"id"`
	Userid      int    `json:"userid" xml:"userid"`             //用户id
	CommandMsg  string `json:"command_msg" xml:"command_msg"`   //回复信息
	CommandTime string `json:"command_time" xml:"command_time"` //回复时间
	CommandImgs string `json:"command_imgs" xml:"command_imgs"` //回复图片
}

//商品收藏
type SaveGoods struct {
	ID      int    `json:"id" xml:"id"`
	Userid  int    `json:"userid" xml:"userid"`   //用户id
	Status  int    `json:"status" xml:"status"`   //status  1:  收藏   2：取消收藏
	Goodsid int    `json:"goodsid" xml:"goodsid"` //课程id
	Time    string `json:"time" xml:"time"`       //时间
}

//地址信息
type Address struct {
	ID            int    `json:"id" xml:"id"`
	Address       string `json:"address" xml:"address"`               //地址
	AddressDetail string `json:"address_detail" xml:"address_detail"` //详细地址
	Userid        int    `json:"userid" xml:"userid"`                 //用户id
	Phone         string `json:"phone" xml:"phone"`                   //电话号码
	UserName      string `json:"user_name" xml:"user_name"`           //用户姓名
}

//已购买商品
type HaveBuy struct {
	ID      int `json:"id" xml:"id"`
	Userid  int `json:"userid" xml:"userid"`   //用户id
	Goodsid int `json:"goodsid" xml:"goodsid"` //课程id
}
