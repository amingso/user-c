package main

import (
	baseInterceptor "boolyeye.com/base/interceptor"
	"boolyeye.com/base/server"
	"boolyeye.com/db/redis"
	sqldb "boolyeye.com/db/sqldb"
	"boolyeye.com/uc/uc-svc/context"
	"boolyeye.com/uc/uc-api/api-loginx/interceptor"
	"boolyeye.com/uc/uc-api/api-loginx/router"
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
