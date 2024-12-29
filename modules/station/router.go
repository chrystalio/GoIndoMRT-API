package station

import "github.com/gin-gonic/gin"

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/station")

	station.GET("", func(c *gin.Context) {
		GetAllStations(c, stationService)
	})
}

func GetAllStations(c *gin.Context, service Service) {
	datas, err := service.GetAllStations()

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, datas)
}
