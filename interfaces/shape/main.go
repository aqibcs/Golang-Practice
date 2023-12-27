package main

import "fmt"

// Define an interface named Shape
type Shape interface {
    Area() float64
}

// Implement the Shape interface for a Rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Implement the Shape interface for a Circle
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    // Create instances of Rectangle and Circle
    rectangle := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 2}

    // Use the interface to calculate the area of different shapes
    shapes := []Shape{rectangle, circle}
    for _, shape := range shapes {
        fmt.Printf("Area: %f\n", shape.Area())
    }
}
