package hello

import (
	"template/internal/service/httpEngine"
)

func Init()  {
	hello := httpEngine.Group("/hello")
	hello.GET("",sayHello)
}