package connect

import (
	"connect/config"
	"connect/route"
)

func main() {
	config.SetEnv()                     // truyền biến môi trường vào chường trình để dùng
	app := route.NewService()           // Khởi tạo route sử dụng Gin
	if err := app.Start(); err != nil { // chạy chương trình
		panic(err)
	}
}
