package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RubberDuck struct {
	ID       string `json:"id"`
	Color    string `json:"color"`
	Material string `json:"material"`
	Size     string `json:"size"`
}

var rubberDucks = []RubberDuck{}

func main() {
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/rubberducks", getRubberDucks)
	e.POST("/rubberducks", createRubberDuck)
	e.PUT("/rubberducks/:id", updateRubberDuck)
	e.DELETE("/rubberducks/:id", deleteRubberDuck)

	e.Logger.Fatal(e.Start(":8080"))
}

func getRubberDucks(c echo.Context) error {
	return c.JSON(http.StatusOK, rubberDucks)
}

func createRubberDuck(c echo.Context) error {
	rubberDuck := new(RubberDuck)
	if err := c.Bind(rubberDuck); err != nil {
		return err
	}
	rubberDuck.ID = uuid.New().String()
	rubberDucks = append(rubberDucks, *rubberDuck)
	return c.JSON(http.StatusCreated, rubberDuck)
}

func updateRubberDuck(c echo.Context) error {
	id := c.Param("id")
	for i, duck := range rubberDucks {
		if id == duck.ID {
			if err := c.Bind(&rubberDucks[i]); err != nil {
				return err
			}
			return c.JSON(http.StatusOK, rubberDucks[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Rubber duck not found"})
}

func deleteRubberDuck(c echo.Context) error {
	id := c.Param("id")
	for i, duck := range rubberDucks {
		if id == duck.ID {
			rubberDucks = append(rubberDucks[:i], rubberDucks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Rubber duck not found"})
}
