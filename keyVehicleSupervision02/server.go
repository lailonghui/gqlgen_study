package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"lai.com/gqlgen_study/keyVehicleSupervision02/graph/generated"
	"lai.com/gqlgen_study/keyVehicleSupervision02/graph/resolvers"
	"net/http"
	"time"
)

var mb int64 = 1 << 20

func main() {
	// Setting up Gin
	r := gin.Default()

	r.POST("/query", graphqlHandler())
	r.POST("/delete", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":8080")
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	//NewExecutableSchema and Config are in the generated.go file
	//Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))
	h.AddTransport(transport.POST{})
	//开启websocket功能
	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	//设置上传文件的大小
	h.AddTransport(transport.MultipartForm{
		//此选项指定用于将请求正文解析为内存中的multipart/form-data的最大字节数，其余部分存储在磁盘上的临时文件中
		MaxMemory: 32 * mb,
		//此选项指定用于将请求正文解析为multipart/form-data的最大字节数。
		MaxUploadSize: 50 * mb,
	})
	h.Use(extension.Introspection{})
	return func(c *gin.Context) {
		//dataloader.Middleware(h).ServeHTTP(c.Writer, c.Request)
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	fmt.Println(222)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
