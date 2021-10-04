module recruiting_system

go 1.16

require (
	github.com/gin-gonic/gin v1.7.2
	github.com/go-conflict/toolkit v0.0.0-20210408070932-623044c23489
	github.com/pkg/errors v0.8.1
	github.com/spf13/viper v1.8.1
	go.uber.org/zap v1.17.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.11
)

replace github.com/go-conflict/toolkit => ./toolkit
