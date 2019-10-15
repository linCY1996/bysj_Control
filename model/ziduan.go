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
	ID     int    `json:"id" xml:"id"`
	Tel    string `json:"tel" xml:"tel"`       //电话
	Pass   string `json:"pass" xml:"pass"`     //密码
	Email  string `json:"email" xml:"email"`   //邮箱
	Status int    `json:"status" xml:"status"` //0：未登录    1：已登录
	Uname  string `json:"uname" xml:"uname"`
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
	Jpclass        int    `json:"jpclass" xml:"jpclass"`
}

//购物车
type Cart struct {
	ID         int `json:"id" xml:"id"`
	Bookid     int `json:"bookid" xml:"bookid"`         //课程id
	Userid     int `json:"userid" xml:"userid"`         //用户id
	Booknum    int `json:"booknum" xml:"booknum"`       //购买量
	Bookstatus int `json:"bookstatus" xml:"bookstatus"` //购买状态
}

//评论
type Command struct {
	ID          int    `json:"id" xml:"id"`
	Goodsid     int    `json:"goodsid" xml:"goodsid"`         //课程id
	Commandmsg  string `json:"commandmsg" xml:"commandmsg"`   //回复信息
	Commandtime string `json:"commandtime" xml:"commandtime"` //回复时间
	Uid         int    `json:"uid" xml:"uid"`
	Uhead       string `json:"uhead" xml:"uhead"`
	Uname       string `json:"uname" xml:"uname"`
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

//首页轮播图下边得子级入口展示
type GetLabel struct {
	ID    int    `json:"id" xml:"id"`
	Imgs  string `json:"imgs" xml:"imgs"`
	Names string `json:"names" xml:"names"`
}

//详情页底部多图展示
type Bookdetailimgs struct {
	ID             int    `json:"id" xml:"id"`
	Bookid         int    `json:"bookid" xml:"bookid"`
	Bookdetailimgs string `json:"bookdetailimgs" xml:"bookdetailimgs"`
}

//分类侧边栏得显示名称
type Slidebar struct {
	ID      int    `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`       //名字
	Classid int    `json:"classid" xml:"classid"` //类别
}
