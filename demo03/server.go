package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"lai.com/gqlgen_study/demo03/graph/generated"
	"lai.com/gqlgen_study/demo03/graph/model"
	"lai.com/gqlgen_study/demo03/graph/resolvers"
	"net/http"
	"time"
)

var mb int64 = 1 << 20

func main() {
	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":8080")
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	//NewExecutableSchema and Config are in the generated.go file
	//Resolver is in the resolver.go file
	//使用NewDefaultServer自动开启自省功能
	config := generated.Config{Resolvers: &resolvers.Resolver{}}

	config.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		// block calling the next resolver
		return nil, fmt.Errorf("Access denied")
		// or let it pass through
		//return next(ctx)
	}

	h := handler.NewDefaultServer(generated.NewExecutableSchema(config))
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

	//开启服务器的自省
	h.AddTransport(transport.Options{})
	h.Use(extension.Introspection{})

	//禁用身份验证的自省
	//h.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	//	graphql.GetOperationContext(ctx).DisableIntrospection = true
	//	return next(ctx)
	//})

	//设置查询复杂度
	//h.Use(extension.FixedComplexityLimit(5))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	fmt.Println(222)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
