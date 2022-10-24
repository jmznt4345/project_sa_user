package main

import (
	"github.com/jmznt4345/sa-65-example/controller"
	"github.com/jmznt4345/sa-65-example/middlewares"
	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// User Routes 1
			r.GET("/users", controller.ListUsers)
			r.GET("/user/:id", controller.GetUser)
			r.POST("/users", controller.CreateUser)
			r.PATCH("/users", controller.UpdateUser)
			r.DELETE("/users/:id", controller.DeleteUser)

			// Request Routes 2
			r.GET("/requests", controller.ListRequests)
			r.GET("/request/:id", controller.GetRequest)
			r.POST("/requests", controller.CreateRequest)
			r.PATCH("/requests", controller.UpdateRequest)
			r.DELETE("/requests/:id", controller.DeleteRequest)

			// Cart Routes 3
			r.GET("/carts", controller.ListCarts)
			r.GET("/cart/:id", controller.GetCart)
			r.POST("/carts", controller.CreateCart)
			r.PATCH("/carts", controller.UpdateCart)
			r.DELETE("/carts/:id", controller.DeleteCart)

			// Room has Device Routes 4
			r.GET("/room_has_devices", controller.ListRoom_has_Devices)
			r.GET("/room_has_device/:id", controller.GetRoom_has_Device)
			r.GET("/room_has_device/room/:id", controller.GetRHD_Device)
			r.POST("/room_has_devices", controller.CreateRoom_has_Device)
			r.PATCH("/room_has_devices", controller.UpdateRoom_has_Device)
			r.DELETE("/room_has_devices/:id", controller.DeleteRoom_has_Device)

			// Device Routes 5
			r.GET("/devices", controller.ListDevices)
			r.GET("/device/:id", controller.GetDevice)
			r.POST("/devices", controller.CreateDevice)
			r.PATCH("/devices", controller.UpdateDevice)
			r.DELETE("/devices/:id", controller.DeleteDevice)

			// History Routes 6
			r.GET("/historys", controller.ListHistorys)
			r.GET("/history/:id", controller.GetHistory)
			r.POST("/historys", controller.CreateHistory)
			r.PATCH("/historys", controller.UpdateHistory)
			r.DELETE("/historys/:id", controller.DeleteHistory)

			// Building Routes 7
			r.GET("/buildings", controller.ListBuildings)
			r.GET("/building/:id", controller.GetBuilding)

			// Room Routes 8
			r.GET("/rooms", controller.ListRooms)
			r.GET("/room/:id", controller.GetRoom)
			r.GET("/rooms/building/:id", controller.GetRoomBuilding)

			// Role Routes 9
			r.GET("/roles", controller.ListRoles)
			r.GET("/role/:id", controller.GetRole)

			// JobType Routes 10
			r.GET("/jobtypes", controller.ListJobTypes)
			r.GET("/jobtype/:id", controller.GetJobType)

			// DMG Level Routes 11
			r.GET("/dmglevels", controller.ListDMGLevels)
			r.GET("/dmglevel/:id", controller.GetDMGLevel)

			// gender Routes 12
			r.GET("/genders", controller.ListGenders)
			r.GET("/gender/:id", controller.GetGender)

			// estimate Routes 13
			r.GET("/estimates", controller.ListEstimates)
			r.GET("/estimate/:id", controller.GetEstimate)

			// Educational_background Routes 14
			r.GET("/educational_backgrounds", controller.ListEducational_backgrounds)
			r.GET("/educational_background/:id", controller.GetEducational_background)

			// brand Routes 15
			r.GET("/brands", controller.ListBrands)
			r.GET("/brand/:id", controller.GetBrand)

			// distributor Routes 16
			r.GET("/distributors", controller.ListDistributors)
			r.GET("/distributor/:id", controller.GetDistributor)

			// type Routes 17
			r.GET("/types", controller.ListTypes)
			r.GET("/type/:id", controller.GetType)
		}

	}
	// Signup User Route
	r.POST("/signup", controller.CreateUser)
	// login User Route
	r.POST("/login", controller.Login)
	// Run the server go run main.go
	r.Run()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}