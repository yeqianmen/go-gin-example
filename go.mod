module github.com/yeqianmen/go-gin-example

go 1.13

replace (
	github.com/yeqianmen/go-gin-example/conf => ~/Downloads/GO语言/go-gin-example/pkg/conf v1.0.0
	github.com/yeqianmen/go-gin-example/middleware => ~/Downloads/GO语言/go-gin-example/middleware v1.0.0
	github.com/yeqianmen/go-gin-example/models => ~/Downloads/GO语言/go-gin-example/models v1.0.0
	github.com/yeqianmen/go-gin-example/pkg/setting => ~/Downloads/GO语言/go-gin-example/pkg/setting v1.0.0
	github.com/yeqianmen/go-gin-example/routers => ~/Downloads/GO语言/go-gin-example/routers v1.0.0
)

require (
	github.com/bytedance/sonic v1.11.7 // indirect
	github.com/gabriel-vasile/mimetype v1.4.4 // indirect
	github.com/gin-gonic/gin v1.10.0 // indirect
	github.com/go-ini/ini v1.67.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/unknwon/com v1.0.1 // indirect
)
