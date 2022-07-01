module boolyeye.com/uc/uc-api/api-loginx

go 1.15

require (
    boolyeye.com/db/redis v0.0.0
    boolyeye.com/db/sqldb v0.0.0
	boolyeye.com/entity v0.0.0
	boolyeye.com/base v0.0.0
	boolyeye.com/uc/uc-svc v0.0.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/x-ream/sqlxb v0.4.3
)

replace (
    boolyeye.com/db/redis => ../../../../../db/boolyeye.com/db/redis
    boolyeye.com/db/sqldb => ../../../../../db/boolyeye.com/db/sqldb
    boolyeye.com/entity => ../../../../../entity/boolyeye.com/entity
    boolyeye.com/base => ../../../../../base/boolyeye.com/base
    boolyeye.com/uc/uc-svc => ../../uc-svc
)
