package control

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vue/model"
)

//用户登录
func Logins(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var tel = r.FormValue(`tel`)
	var pass = r.FormValue(`pass`)
	mods, err := model.Login(tel)
	if err != nil {
		w.Write([]byte(`信息错误`))
		return
	}
	if tel != mods.Tel {
		w.Write([]byte(`电话号码不存在`))
		return
	}
	if mods.Pass != pass {
		w.Write([]byte(`密码不正确`))
		return
	}
	w.Write([]byte(`登录成功`))
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

//忘记密码
func Lostpass(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var email = r.FormValue(`email`)
	var pass = r.FormValue(`pass`)
	var passagain = r.FormValue(`passagain`)
	if pass != passagain {
		w.Write([]byte(`输入密码不一致`))
	} else {
		w.Write([]byte(`ok`))
	}
	err := model.LostPass(pass, passagain, email)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if err != nil {
		w.Write([]byte(`修改失败`))
	} else {
		w.Write([]byte(`添加成功`))
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
	fmt.Println(err)
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
