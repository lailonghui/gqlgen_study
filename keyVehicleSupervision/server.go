package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"lai.com/gqlgen_study/demo02/graph/dataloader"
	"lai.com/gqlgen_study/demo02/graph/generated"
	"lai.com/gqlgen_study/demo02/graph/model"
	"lai.com/gqlgen_study/demo02/graph/resolvers"
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

//设置解析结构体的函数
func getResolver() *resolvers.Resolver {
	resolver := &resolvers.Resolver{}

	resolver.MutationResolver.SingleUpload = func(ctx context.Context, file graphql.Upload) (*model.File, error) {
		content, err := ioutil.ReadAll(file.File)
		if err != nil {
			return nil, err
		}
		return &model.File{
			ID:          1,
			Name:        file.Filename,
			Content:     string(content),
			ContentType: file.ContentType,
		}, nil
	}
	resolver.MutationResolver.SingleUploadWithPayload = func(ctx context.Context, req model.UploadFile) (*model.File, error) {
		content, err := ioutil.ReadAll(req.File.File)
		if err != nil {
			return nil, err
		}
		return &model.File{
			ID:          1,
			Name:        req.File.Filename,
			Content:     string(content),
			ContentType: req.File.ContentType,
		}, nil
	}
	resolver.MutationResolver.MultipleUpload = func(ctx context.Context, files []*graphql.Upload) ([]*model.File, error) {
		if len(files) == 0 {
			return nil, errors.New("empty list")
		}
		var resp []*model.File
		for i := range files {
			content, err := ioutil.ReadAll(files[i].File)
			if err != nil {
				return []*model.File{}, err
			}
			resp = append(resp, &model.File{
				ID:          i + 1,
				Name:        files[i].Filename,
				Content:     string(content),
				ContentType: files[i].ContentType,
			})
		}
		return resp, nil
	}
	resolver.MutationResolver.MultipleUploadWithPayload = func(ctx context.Context, req []*model.UploadFile) ([]*model.File, error) {
		if len(req) == 0 {
			return nil, errors.New("empty list")
		}
		var resp []*model.File
		for i := range req {
			content, err := ioutil.ReadAll(req[i].File.File)
			if err != nil {
				return []*model.File{}, err
			}
			resp = append(resp, &model.File{
				ID:          i + 1,
				Name:        req[i].File.Filename,
				Content:     string(content),
				ContentType: req[i].File.ContentType,
			})
		}
		return resp, nil
	}
	return resolver
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	//NewExecutableSchema and Config are in the generated.go file
	//Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: getResolver()}))
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
	//h.Use(extension.Introspection{})
	return func(c *gin.Context) {
		//fmt.Println(c.PostForm("user"))
		//fmt.Println(c.PostForm("password"))
		//f, _ := c.FormFile("file")
		//fmt.Println(f.Filename)
		//fmt.Println(f.Size)
		dataloader.Middleware(h).ServeHTTP(c.Writer, c.Request)
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
