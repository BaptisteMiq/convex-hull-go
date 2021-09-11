package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Addr string = "localhost:8080"

// Define points structure
// We use struct tag json to specify what the field name should be when serialized into JSON.
// This way we can respect capital letters convention in Go and lowercase letters commonly seen in JSON
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// We use a point list to bind POST request body to it
// Excepted input body format:
// { "points": [{ "x": .., "y": .. }, ...] }
type PointList struct {
	Points []Point `json:"points"`
}

// Get a Convex Hull with given points using the Gift Wrapping algorithm
func getConvexHull(points []Point) ([]Point, error) {

	n := len(points)
	var result []Point

	// Initial condition (n >= 3)
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
	current := left           // The last point we found
	next := (current + 1) % n // The point we assume to be the next one (verified after checking orientation)

	// Find next point until we wrapped around
	for next != left {

		// Update vars and append the found point to the result
		next = (current + 1) % n
		result = append(result, points[current])

		// Get the next point by rotating counterclockwise
		for i, p := range points {
			c, n := points[current], points[next]
			// Check for counterclockwise orientation between p, c and n
			// This formula is from https://preparingforcodinginterview.wordpress.com/2016/08/30/orientation-of-3-ordered-points/
			// It allows to find the orientation of the 3 ordered points by looking at the slope of the two segments:
			// - p and c (σ)
			// - c and n (τ)
			// If σ < τ, then the orientation is counterclockwise
			sigma := (p.Y - c.Y) * (n.X - p.X)
			tau := (p.X - c.X) * (n.Y - p.Y)
			if sigma-tau < 0 {
				next = i // We found our next point
			}
		}
		current = next // Move current point
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

	// Retrieve points from the list
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
	router.Run(Addr)
}
