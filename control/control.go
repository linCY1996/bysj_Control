package control

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"vue/model"
)

//用户登录
func Logins(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	var tel = r.FormValue(`number`)
	var pass = r.FormValue(`pass`)
	// fmt.Println("tel", tel)
	// fmt.Println("pass", pass)
	mods, _ := model.Login(tel)

	if tel != mods.Tel {
		w.Write([]byte(`0`))
		return
	}
	if mods.Pass != pass {
		w.Write([]byte(`0`))
		return
	}
	errs := model.UndateLoginStatus(1)
	if errs != nil {
		w.Write([]byte(`修改登录状态失败`))
		return
	}
	w.Header().Set("Content-Type", `contentTypeJSON`)
	md, _ := json.Marshal(mods)
	w.Write(md)
	// w.Write([]byte(`登录成功`))
}

//用户注册
func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	var truename = r.FormValue(`username`)
	var pass = r.FormValue(`pass`)
	var email = r.FormValue(`email`)
	var tel = r.FormValue(`tel`)
	fmt.Println("user=", r.FormValue(`username`))
	err := model.UserRegister(truename, pass, email, tel)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(`添加失败`))
	} else {
		w.Write([]byte(`添加成功`))
	}
}

//查询用户信息
func LookUserMsgs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var userid = r.FormValue(`uid`)
	uid, _ := strconv.Atoi(userid)
	mod, err := model.LookUserMsgs(uid)
	// fmt.Println("uid", uid)
	// fmt.Println("err", err)
	if err != nil {
		w.Write([]byte(`0`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//忘记密码
func Lostpass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	r.ParseForm()
	var email = r.FormValue(`email`)
	var pass = r.FormValue(`pass`)
	var passagain = r.FormValue(`passagain`)
	// fmt.Println("email", email)
	// fmt.Println("pass", pass)
	_, errs := model.EmailLook(email)
	// fmt.Println("errs", errs)
	if errs != nil {
		w.Write([]byte(`没有该用户`))
	} else {
		err := model.LostPass(pass, email)
		// fmt.Println("err", err)
		if err != nil {
			w.Write([]byte(`该邮箱不存在`))
		} else if pass != passagain {
			w.Write([]byte(`密码不一致`))
		} else {
			w.Write([]byte(`修改成功`))
		}
	}

}

//轮播
func ShowLun(w http.ResponseWriter, r *http.Request) {
	mod, err := model.Lun()
	if err != nil {
		w.Write([]byte("查询信息错误"))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//获取课程列表
func BookList(w http.ResponseWriter, r *http.Request) {
	mod, err := model.GetBookList()
	// fmt.Println(err)
	if err != nil {
		w.Write([]byte("查询信息错误"))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//获取首页子级入口得信息展示
func GetIndexChr(w http.ResponseWriter, r *http.Request) {
	mod, err := model.GetLabels()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if err != nil {
		w.Write([]byte(`查询信息失败`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//根据进入得首页入口不同。查询不一样种类课程
func LookDetailsMsgs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var class = r.FormValue(`class`)
	bclass, _ := strconv.Atoi(class)
	mod, err := model.LookdetailClass(bclass)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if err != nil {
		w.Write([]byte(`查询信息失败`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//添加商品进入购物车
func AddGoodsIntoCart(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var bookid = r.FormValue(`bookid`)
	var userid = r.FormValue(`userid`)
	bid, _ := strconv.Atoi(bookid)
	uid, _ := strconv.Atoi(userid)
	err := model.AddCart(bid, uid)
	if err != nil {
		w.Write([]byte(`0`))
		return
	}
	w.Write([]byte(`1`))
}

//查询商品详情信息
func LookGoodsDetailMsgs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var goodsid = r.FormValue(`gid`)
	gid, _ := strconv.Atoi(goodsid)
	mod, err := model.LookGoodsDetailMsgs(gid)
	// fmt.Println("uid", uid)
	// fmt.Println("err", err)
	if err != nil {
		w.Write([]byte(`0`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//查询商品详情底部多图展示
func LookGoodsDetailArrayImgs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var goodsid = r.FormValue(`gid`)
	gid, _ := strconv.Atoi(goodsid)
	mod, err := model.ShowDetailImgsArray(gid)
	// fmt.Println("uid", uid)
	// fmt.Println("err", err)
	if err != nil {
		w.Write([]byte(`0`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//收藏或删除商品
//status    0:删除     1：收藏
func ControlSaveGoods(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var userid = r.FormValue(`uid`)
	var goodsid = r.FormValue(`gid`)
	var time = r.FormValue(`time`)
	var Status = r.FormValue(`status`)
	uid, _ := strconv.Atoi(userid)
	gid, _ := strconv.Atoi(goodsid)
	status, _ := strconv.Atoi(Status)
	err := model.ControlGoods(uid, gid, status, time)
	fmt.Println("err", err)
	if err != nil {
		w.Write([]byte(`0`))
		return
	}
	w.Write([]byte(`1`))
}

//通过用户id去查是否收藏了该商品
func UserIdLookSaveGoods(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var userid = r.FormValue(`uid`)
	uid, _ := strconv.Atoi(userid)
	mod, err := model.LookGoodsIsSave(uid)
	if err != nil {
		w.Write([]byte(`0`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//查询指定商品得全部评价
func LookGoodsAllCommands(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var goodsid = r.FormValue(`gid`)
	gid, _ := strconv.Atoi(goodsid)
	mod, err := model.GetDetailCommand(gid)
	if err != nil {
		w.Write([]byte(`0`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//查询指定商品得全部评价
func GetSlideBar(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mod, err := model.SlideType()
	if err != nil {
		w.Write([]byte(`0`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Write(md)
}
