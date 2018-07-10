package main

import (
	"net/http"
	"log"
	"io"
	"path/filepath"
	"html/template"
	"fmt"
	"zhq/goweb/controller"
)

func main() {
	//指定当前文件所属目录为项目Web根目录
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/",http.StripPrefix("/",fs))
	http.HandleFunc("/temp.html",TempView)
	http.HandleFunc("/login.html", LoginView)
	http.HandleFunc("/login",Login)
	//设置监听端口
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("端口监听发生错误:", err)
	}

}
func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"hello world")
}
func TempView(w http.ResponseWriter,r *http.Request){
	page := filepath.Join("temp.html")
	//创建page转为一个模板实例
	result,err := template.ParseFiles(page)
	if err != nil {
		fmt.Println("创建模板实例错误:",err)
	}
	//融合模板数据
	err = result.Execute(w,nil)
	if err != nil {
		fmt.Println("融合模板数据时发生错误:",err)
	}
}
func LoginView(w http.ResponseWriter, r *http.Request)  {

	// 指定页面
	page := filepath.Join( "login.html")

	// 创建模板实例
	templateResult, err := template.ParseFiles(page)
	if err != nil {
		fmt.Println("创建模板实例错误: ", err)
	}

	// 融合模板数据
	err = templateResult.Execute(w, nil)
	if err != nil {
		fmt.Println("融合模板数据时发生错误: ", err)
	}

}
func Login(w http.ResponseWriter,r *http.Request)  {
	userName := r.FormValue("userName")
	password := r.FormValue("password")
	loginSuccess := false
	fmt.Println(userName," ====  ",password)
	for _,person := range controller.Persons {
		if person.Name == userName {
			if person.Password == password {
				loginSuccess = true
				break
			}else {
				loginSuccess = false
				break
			}
		}
	}
	//如果用户名，密码验证成功，跳转至登录成功页面
	if loginSuccess {
		// 指定页面
		page := filepath.Join( "loginSuccess.html")
		// 创建模板实例
		templateResult, err := template.ParseFiles(page)
		if err != nil {
			fmt.Println("创建模板实例错误: ", err)
		}
		//需要展示的数据
		data := &struct {
			Persons []controller.Person
		}{
			Persons:controller.Persons,
		}
		// 融合模板数据
		err = templateResult.Execute(w, data)
		if err != nil {
			fmt.Println("融合模板数据时发生错误: ", err)
		}
	}else {
		// 指定页面
		page := filepath.Join( "login.html")
		// 创建模板实例
		templateResult, err := template.ParseFiles(page)
		if err != nil {
			fmt.Println("创建模板实例错误: ", err)
		}
		//需要展示的数据
		data := &struct {
			Flag bool
		}{
			Flag: true,
		}
		// 融合模板数据
		err = templateResult.Execute(w, data)
		if err != nil {
			fmt.Println("融合模板数据时发生错误: ", err)
		}
	}

}