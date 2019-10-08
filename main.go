package main

import (
	"net/http"
	"vue/control"
)

//中间件
// func Mid(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		return
// 	})
// }

func main() {
	http.HandleFunc(`/api/lunimgs`, control.ShowLun)       //轮播图信息
	http.HandleFunc(`/api/showMoreMsgs`, control.BookList) //课程列表信息
	http.HandleFunc(`/api/register`, control.Register)     //用户注册
	http.HandleFunc(`/api/updatePass`, control.Lostpass)   //用户修改密码
	http.HandleFunc(`/api/login`, control.Logins)          //用户登录

	http.ListenAndServe(`:6933`, nil)
}
