// api/index.go
package handler

import (
	"douyin-backend/router" // 导入您项目中的路由包
	"net/http"
)

// Handler 是 Vercel 调用的唯一入口函数
// 它会将接收到的请求交给 Gin 引擎处理
func Handler(w http.ResponseWriter, r *http.Request) {
	// 初始化 Gin 路由
	engine := router.InitRouter() 

	// 使用 Gin 引擎处理所有传入的请求
	engine.ServeHTTP(w, r)
}
