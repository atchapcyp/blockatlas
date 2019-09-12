package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/platform"
)

var routers = make(map[string]gin.IRouter)

func loadPlatforms(root gin.IRouter) {
	v1 := root.Group("/v1")
	v2 := root.Group("/v2")

	for _, txAPI := range platform.Platforms {
		router := getRouter(v1, txAPI.Coin().Handle)
		makeTxRouteV1(router, txAPI)

		routerV2 := getRouter(v2, txAPI.Coin().Handle)
		makeTxRouteV2(routerV2, txAPI)
	}

	for _, tokenAPI := range platform.Platforms {
		router := getRouter(v2, tokenAPI.Coin().Handle)
		makeTokenRoute(router, tokenAPI)
	}

	for _, stakeAPI := range platform.Platforms {
		router := getRouter(v2, stakeAPI.Coin().Handle)
		makeStakingRoute(router, stakeAPI)
	}

	for _, collectionAPI := range platform.Platforms {
		router := getRouter(v2, collectionAPI.Coin().Handle)
		makeCollectionRoute(router, collectionAPI)
	}

	for _, customAPI := range platform.CustomAPIs {
		router := getRouter(v1, customAPI.Coin().Handle)
		customAPI.RegisterRoutes(router)
	}

	logrus.WithField("routes", len(routers)).Info("Routes set up")

	v1.GET("/", getEnabledEndpoints)
}

// getRouter lazy loads routers
func getRouter(router *gin.RouterGroup, handle string) gin.IRouter {
	key := fmt.Sprintf("%s/%s", router.BasePath() , handle)
	if group, ok := routers[key]; ok {
		return group
	} else {
		path := fmt.Sprintf("/%s", handle)
		logrus.Debugf("Registering %s", path)
		group := router.Group(path)
		routers[key] = group
		return group
	}
}
