# 运行环境
`在go 1.24.5上运行
配置文件：config.yaml，其中填写myslq的dsn地址，需要先建立自己的数据库schema
jwtSecret秘钥可以根据自己的需要进行更改`


# 依赖安装步骤
>```
>go get -u github.com/gin-gonic/gin
>go get -u gorm.io/gorm
>go get -u gorm.io/driver/mysql
>go get -u github.com/golang-jwt/jwt/v5
>go get -u gopkg.in/yaml.v3
>go get -u golang.org/x/crypto/bcrypt
>```

# 启动方式
>```
>go run main.go
>```
