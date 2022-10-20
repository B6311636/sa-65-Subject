package main

import (
	"github.com/B6311636/sa-G12-Subject/controller"
	"github.com/B6311636/sa-G12-Subject/entity"
	"github.com/gin-gonic/gin"
)

const PORT = "8888"

func main() {
	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// router := r.Group("/")
	// {
	// 	router.Use(middlewares.Authorizes())
	// 	{
	// 		// Officer Routes
	// 		router.GET("/users", controller.ListOfficers)
	// 		router.GET("/user/:id", controller.GetOfficer)
	// 		router.PATCH("/users", controller.UpdateOfficer)
	// 		router.DELETE("/users/:id", controller.DeleteOfficer)
	// 		// Teacher Routes
	// 		router.GET("/teachers", controller.ListTeachers) // ("path", cotroller)
	// 		router.GET("/teacher/:id", controller.GetTeacher)
	// 		router.POST("/teachers", controller.CreateTeacher)
	// 		router.PATCH("/teachers", controller.UpdateTeacher)
	// 		router.DELETE("/teachers/:id", controller.DeleteTeacher)
	// 		// Faculty Routes
	// 		router.GET("/facultys", controller.ListFaculties) // ("path", cotroller)
	// 		router.GET("/faculty/:id", controller.GetFaculty)
	// 		router.POST("/facultys", controller.CreateFaculty)
	// 		router.PATCH("/facultys", controller.UpdateFaculty)
	// 		router.DELETE("/facultys/:id", controller.DeleteFaculty)
	// 		// Time Routes
	// 		router.GET("/times", controller.ListTimes) // ("path", cotroller)
	// 		router.GET("/time/:id", controller.GetTime)
	// 		router.POST("/times", controller.CreateTime)
	// 		router.PATCH("/times", controller.UpdateTime)
	// 		router.DELETE("/times/:id", controller.DeleteTime)
	// 		// Subject Routes
	// 		router.GET("/subjects", controller.ListSubject) // ("path", cotroller)
	// 		router.GET("/subject/:id", controller.GetSubject)
	// 		router.POST("/subjects", controller.CreateSubject)
	// 		router.PATCH("/subjects", controller.UpdateSubject)
	// 		router.DELETE("/subjects/:id", controller.DeleteSubject)

	// 	}
	// }

	// Officer Routes
	r.GET("/users", controller.ListOfficers)
	r.GET("/user/:id", controller.GetOfficer)
	r.PATCH("/users", controller.UpdateOfficer)
	r.DELETE("/users/:id", controller.DeleteOfficer)
	// Teacher Routes
	r.GET("/teachers", controller.ListTeachers) // ("path", cotroller)
	r.GET("/teacher/:id", controller.GetTeacher)
	r.POST("/teachers", controller.CreateTeacher)
	r.PATCH("/teachers", controller.UpdateTeacher)
	r.DELETE("/teachers/:id", controller.DeleteTeacher)
	// Faculty Routes
	r.GET("/facultys", controller.ListFaculties) // ("path", cotroller)
	r.GET("/faculty/:id", controller.GetFaculty)
	r.POST("/facultys", controller.CreateFaculty)
	r.PATCH("/facultys", controller.UpdateFaculty)
	r.DELETE("/facultys/:id", controller.DeleteFaculty)
	// Time Routes
	r.GET("/times", controller.ListTimes) // ("path", cotroller)
	r.GET("/time/:id", controller.GetTime)
	r.POST("/times", controller.CreateTime)
	r.PATCH("/times", controller.UpdateTime)
	r.DELETE("/times/:id", controller.DeleteTime)
	// Subject Routes
	r.GET("/subjects", controller.ListSubject) // ("path", cotroller)
	r.GET("/subject/:id", controller.GetSubject)
	r.POST("/subjects", controller.CreateSubject)
	r.PATCH("/subjects", controller.UpdateSubject)
	r.DELETE("/subjects/:id", controller.DeleteSubject)

	// Signup User Route
	r.POST("/signup", controller.CreateOfficer)
	// login User Route
	r.POST("/login", controller.Login)
	// Run the server go run main.go
	r.Run("localhost: " + PORT) // "localhost: " + PORT
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
