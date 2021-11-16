package main

import (
	"CourseAPI/model"
	"fmt"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func main() {
	err := mgm.SetDefaultConfig(nil, "EDteam", options.Client().ApplyURI("mongodb+srv://user_cluster:Rykly2017@micluster.hwvlc.mongodb.net"))
	if err != nil {
		fmt.Println("Error on connect")
	}
	e := echo.New()
	e.GET("/courses", getAllCourses)
	e.GET("/course/:id", getOneCourse)
	e.DELETE("/course/:id", deleteCourse)
	e.PUT("/course/:id", updateCourse)
	e.POST("/courses", saveCourse)
	e.Logger.Fatal(e.Start(":8080"))
}

func getAllCourses(c echo.Context) error {
	courses := []model.Course{}
	err := mgm.Coll(&model.Course{}).SimpleFind(&courses, bson.M{})
	if err != nil {
		fmt.Println("Error on find all")
	}
	return c.JSON(http.StatusOK, courses)
}

func saveCourse(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	newCourse := model.NewCourse(name, description)

	err := mgm.Coll(newCourse).Create(newCourse)
	newCourse.SetID(newCourse.ID)

	_ = mgm.Coll(newCourse).Update(newCourse)
	if err != nil {
		fmt.Println("Error on insert")
	}
	return c.JSON(http.StatusCreated, newCourse)
}

func getOneCourse(c echo.Context) error {
	id := c.Param("id")
	obj, _ := primitive.ObjectIDFromHex(id)
	course := &model.Course{}
	coll := mgm.Coll(course)
	err := coll.FindByID(id, course)
	if err == nil {
		course.SetID(obj)
		return c.JSON(http.StatusOK, course)
	} else {
		return c.JSON(http.StatusNotFound, struct {
			Message string
		}{"Not found"})
	}
}

func deleteCourse(c echo.Context) error {
	id := c.Param("id")
	obj, _ := primitive.ObjectIDFromHex(id)
	course := &model.Course{}
	coll := mgm.Coll(&model.Course{})
	err1 := coll.FindByID(id, course)
	if err1 == nil {
		course.SetID(obj)
		err2 := coll.Delete(course)
		if err2 != nil {
			fmt.Println("Error on delete")
		}
		return c.JSON(http.StatusOK, course)
	} else {
		return c.JSON(http.StatusNotFound, struct {
			Message string
		}{"Not found"})
	}

}
func updateCourse(c echo.Context) error {
	id := c.Param("id")
	obj, _ := primitive.ObjectIDFromHex(id)
	name := c.FormValue("name")
	description := c.FormValue("description")

	course := &model.Course{}
	coll := mgm.Coll(&model.Course{})

	err1 := coll.FindByID(id, course)

	if err1 == nil {
		course.SetID(obj)
		if name != "" {
			course.Name = name
		}
		if description != "" {
			course.Description = description
		}
		err2 := coll.Update(course)
		if err2 != nil {
			fmt.Println("Error on update")
		}
		return c.JSON(http.StatusOK, course)
	} else {
		return c.JSON(http.StatusNotFound, struct {
			Message string
		}{"Not found"})
	}

}
