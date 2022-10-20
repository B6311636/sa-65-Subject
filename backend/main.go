package main

import (
	"github.com/B6311636/sa-G12-Subject/controller"
	"github.com/B6311636/sa-G12-Subject/entity"
	"github.com/B6311636/sa-G12-Subject/middlewares"
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
			// Officer Routes
			router.GET("/officers", controller.ListOfficers)
			router.GET("/officer/:id", controller.GetOfficer)
			router.POST("/officers/create", controller.CreateOfficer)
			router.PATCH("/officers", controller.UpdateOfficer)
			router.DELETE("/officers/:id", controller.DeleteOfficer)
			// Teacher Routes
			router.GET("/teachers", controller.ListTeachers) // ("path", cotroller)
			router.GET("/teacher/:id", controller.GetTeacher)
			router.POST("/teachers", controller.CreateTeacher)
			router.PATCH("/teachers", controller.UpdateTeacher)
			router.DELETE("/teachers/:id", controller.DeleteTeacher)
			// Faculty Routes
			router.GET("/facultys", controller.ListFaculties) // ("path", cotroller)
			router.GET("/faculty/:id", controller.GetFaculty)
			router.POST("/facultys", controller.CreateFaculty)
			router.PATCH("/facultys", controller.UpdateFaculty)
			router.DELETE("/facultys/:id", controller.DeleteFaculty)
			// Time Routes
			router.GET("/times", controller.ListTimes) // ("path", cotroller)
			router.GET("/time/:id", controller.GetTime)
			router.POST("/times", controller.CreateTime)
			router.PATCH("/times", controller.UpdateTime)
			router.DELETE("/times/:id", controller.DeleteTime)
			// Subject Routes
			router.GET("/subjects", controller.ListSubject) // ("path", cotroller)
			router.GET("/subject/:id", controller.GetSubject)
			router.POST("/subjects", controller.CreateSubject)
			router.PATCH("/subjects", controller.UpdateSubject)
			router.DELETE("/subjects/:id", controller.DeleteSubject)

		}
	}

	// Signup User Route
	r.POST("/signup", controller.CreateOfficer)
	// login User Route
	r.POST("/login", controller.Login)
	// Run the server go run main.go
	r.Run() // "localhost: " + PORT
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
