package main

import (
	"fmt"
	"math"
)

const (
	RECTANGLE = "rectangle"
	CIRCLE    = "circle"
	TRIANGLE  = "triangle"
)

type Shape struct {
	ShapeType string
	Length    int
	Area      float32
}

func calculateArea(shapeType string, lenght int, result chan Shape) {
	var area float32
	switch shapeType {
	case RECTANGLE:
		area = float32(lenght * lenght)
	case CIRCLE:
		area = float32(math.Pi) * float32(lenght) * float32(lenght)
	case TRIANGLE:
		area = 0.5 * float32(lenght) * float32(lenght)
	}

	shape := Shape{
		ShapeType: shapeType,
		Length:    lenght,
		Area:      area,
	}
	result <- shape
}

func main() {
	input := []Shape{
		{ShapeType: RECTANGLE, Length: 5},
		{ShapeType: CIRCLE, Length: 3},
		{ShapeType: TRIANGLE, Length: 5},
		{ShapeType: RECTANGLE, Length: 15},
		{ShapeType: CIRCLE, Length: 5},
	}

	rectangleChannel := make(chan Shape)
	circleChannel := make(chan Shape)
	triangleChannel := make(chan Shape)

	// Start goroutine ro calculate
	for _, shape := range input {
		switch shape.ShapeType {
		case RECTANGLE:
			go calculateArea(RECTANGLE, shape.Length, rectangleChannel)
		case CIRCLE:
			go calculateArea(CIRCLE, shape.Length, circleChannel)
		case TRIANGLE:
			go calculateArea(TRIANGLE, shape.Length, triangleChannel)
		}
	}

	// receive data from channels
	for i := 0; i < len(input); i++ {
		select {
		case rectangle := <-rectangleChannel:
			fmt.Printf("Rectangle (Length: %d) - Area: %.2f\n", rectangle.Length, rectangle.Area)
		case circle := <-circleChannel:
			fmt.Printf("Circle (Radius: %d) - Area: %.2f\n", circle.Length, circle.Area)
		case triangle := <-triangleChannel:
			fmt.Printf("Triangle (Base: %d) - Area: %.2f\n", triangle.Length, triangle.Area)
		}
	}
}
