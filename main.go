package main

import (
	"fmt"
	"net/http"

	"github.com/ihsan/pragmatic-review/controller"
	router "github.com/ihsan/pragmatic-review/http"
	"github.com/ihsan/pragmatic-review/repository"
	"github.com/ihsan/pragmatic-review/service"
)

var (
	postRepo       repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepo)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "up and running...")
	})
	httpRouter.GET("/post", postController.GetPosts)
	httpRouter.POST("/post", postController.AddPost)

	httpRouter.SERVE(":8080")
}
