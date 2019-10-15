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
	http.HandleFunc(`/api/lunimgs`, control.ShowLun)                                    //轮播图信息
	http.HandleFunc(`/api/showMoreMsgs`, control.BookList)                              //课程列表信息
	http.HandleFunc(`/api/register`, control.Register)                                  //用户注册
	http.HandleFunc(`/api/updatePass`, control.Lostpass)                                //用户修改密码
	http.HandleFunc(`/api/login`, control.Logins)                                       //用户登录
	http.HandleFunc(`/api/get/indexchr`, control.GetIndexChr)                           //获取首页入口子级栏目展示
	http.HandleFunc(`/api/find/detailclass`, control.LookDetailsMsgs)                   //根据指定数据查询对应得课程
	http.HandleFunc(`/api/insertGoods/cart`, control.AddGoodsIntoCart)                  //加入购物车
	http.HandleFunc(`/api/look/userMsgs`, control.LookUserMsgs)                         //查询用户信息
	http.HandleFunc(`/api/look/goodsDetailMsgs`, control.LookGoodsDetailMsgs)           //查询商品信息
	http.HandleFunc(`/api/look/goodsDetailArrayImgs`, control.LookGoodsDetailArrayImgs) //查询详情页底部多图展示
	http.HandleFunc(`/api/control/savegoods`, control.ControlSaveGoods)                 //操作收藏课程操作
	http.HandleFunc(`/api/userid/looksavegoods`, control.UserIdLookSaveGoods)           //通过用户id去查是否收藏了对应得商品
	http.HandleFunc(`/api/look/goodscommand`, control.LookGoodsAllCommands)             //查询指定商品得全部评价
	http.HandleFunc(`/api/get/slidebar`, control.GetSlideBar)                           //获取侧边栏信息

	http.ListenAndServe(`:6933`, nil)
}
