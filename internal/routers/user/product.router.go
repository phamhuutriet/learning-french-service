package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
	// productController controller.ProductController
}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}

	// private router

}
