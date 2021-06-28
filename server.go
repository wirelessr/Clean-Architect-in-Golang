package main

import (
	"fmt"
	"net/http"

	"example.com/clean-arch/controller"
	router "example.com/clean-arch/http"
	"example.com/clean-arch/repository"
	"example.com/clean-arch/service"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)

}
