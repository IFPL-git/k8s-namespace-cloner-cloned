package router

import (
	"github.com/IFPL-git/k8s-namespace-cloner/controllers"
	"github.com/IFPL-git/k8s-namespace-cloner/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func InitializeRoutes(clientset *kubernetes.Clientset, dynamicClientSet *dynamic.DynamicClient) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	v1 := r.Group("/api/v1")

	v1.Use(middlewares.K8sClientSetMiddleware(clientset, dynamicClientSet))
	{
		v1.GET("/namespaces", controllers.GetNS)
		v1.GET("/namespaces/:namespace/deployments", controllers.GetDeployments)
		v1.GET("/namespaces/:namespace/deployments/display", controllers.DisplayDeployments)
		v1.GET("/namespaces/:namespace/secrets/display", controllers.DisplaySecrets)
		v1.GET("/namespaces/:namespace/configmaps/display", controllers.DisplayConfigMap)

		v1.POST("/namespaces/:namespace/cloneNamespace", controllers.CloneNamespace)
		v1.POST("/deployments/:deployment", controllers.UpdateDeploymentImage)
		v1.POST("/secrets/:secret", controllers.UpdateSecret)
		v1.POST("/configmaps/:configmap", controllers.UpdateConfigMap)

		v1.GET("/cohorts", controllers.GetCohorts)
		v1.POST("/cohorts/deployment", controllers.DeployCohort)
	}
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
