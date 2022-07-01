package main

import (
	baseInterceptor "github.com/amingso/user-c/base/interceptor"
	"github.com/amingso/user-c/base/server"
	"github.com/amingso/user-c/db/redis"
	sqldb "github.com/amingso/user-c/db/sqldb"
	"github.com/amingso/user-c/uc/uc-svc/context"
	"github.com/amingso/user-c/uc/uc-api/api-loginx/interceptor"
	"github.com/amingso/user-c/uc/uc-api/api-loginx/router"
)

func main()  {
	context.Init()
	sqldb.InitDB()
	redis.InitRedis("127.0.0.1:6379","")
	server.RegisterInterceptor(baseInterceptor.ExceptionInterceptor)
	server.RegisterInterceptor(interceptor.TxInterceptor)
	server.Listen(":80",router.Config())
	sqldb.CloseDB()
}
