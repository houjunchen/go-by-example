// https://rongli.gitbooks.io/design-pattern/content/chapter3.html

package main

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing a circle")
}

type Square struct{}

func (s *Square) Draw() {
	fmt.Println("Drawing a square")
}

type Color interface {
	Fill()
}

type Red struct{}

func (r *Red) Fill() {
	fmt.Println("Filling red")
}

type White struct{}

func (w *White) Fill() {
	fmt.Println("Filling white")
}

type Factory interface {
	CreateShape(shapeType string) Shape
	CreateColor(colorType string) Color
}

type ShapeFactory struct{}

func (sf *ShapeFactory) CreateShape(shapeType string) Shape {
	switch shapeType {
	case "CIRCLE":
		return &Circle{}
	case "SQUARE":
		return &Square{}
	}
	return nil
}

func (sf *ShapeFactory) CreateColor(colorType string) Color {
	return nil
}

type ColorFactory struct{}

func (cf *ColorFactory) CreateShape(shapeType string) Shape {
	return nil
}

func (cf *ColorFactory) CreateColor(colorType string) Color {
	switch colorType {
	case "RED":
		return &Red{}
	case "WHITE":
		return &White{}
	}
	return nil
}

func CreateFactory(factoryType string) Factory {
	switch factoryType {
	case "SHAPE":
		return &ShapeFactory{}
	case "COLOR":
		return &ColorFactory{}
	}
	return nil
}

func main() {
	shapeFactory := CreateFactory("SHAPE")
	shape1 := shapeFactory.CreateShape("CIRCLE")
	shape1.Draw()
	shapeFactory.CreateShape("SQUARE").Draw()

	colorFactory := CreateFactory("COLOR")
	color1 := colorFactory.CreateColor("RED")
	color1.Fill()
	colorFactory.CreateColor("WHITE").Fill()
}
