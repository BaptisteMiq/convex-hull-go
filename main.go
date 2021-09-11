package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PointList struct {
	Points []Point `json:"points"`
}

func getConvexHull(points []Point) ([]Point, error) {

	n := len(points)
	var result []Point

	if n < 3 {
		return result, errors.New("not enough points")
	}

	// Find the point with minimum x
	left := 0
	for i, p := range points {
		if p.X < points[left].X {
			left = i
		}
	}

	// Initial variables
	result = append(result, points[left])
	current := left
	next := (current + 1) % n

	// Find next point until we wrapped around
	for next != left {

		next = (current + 1) % n
		result = append(result, points[current])

		for i, p := range points {
			c, n := points[current], points[next]
			// Check for counterclockwise orientation between p, c and n
			if (p.Y-c.Y)*(n.X-p.X)-(p.X-c.X)*(n.Y-p.Y) < 0 {
				next = i
			}
		}
		// Update vars and add the found point to the result
		current = next
	}

	return result, nil
}

func postConvexHullHTTP(c *gin.Context) {
	var points []Point
	var pointsList PointList

	// Bind received JSON to points
	if err := c.BindJSON(&pointsList); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	points = pointsList.Points

	// Get the convex hull of points and send the result
	if hull, err := getConvexHull(points); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, hull)
	}

}

func addRoutes(router *gin.Engine) {
	router.POST("/convex2d", postConvexHullHTTP) // The API route to get a convex hull
	router.Static("/", "./public")               // Serve HTML
}

func main() {
	// Init HTTP server
	router := gin.Default()
	addRoutes(router)
	router.Run("localhost:8080")
}
