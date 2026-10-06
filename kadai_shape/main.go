package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Message string
	Field   string
	Value   float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: %s (Field: %s, Value: %f)", e.Message, e.Field, e.Value)
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func validShape(s Shape) error {
	switch shape := s.(type) {
	case Rectangle:
		if shape.Width <= 0 {
			return &ValidationError{
				Message: "幅が0以下です",
				Field:   "Width",
				Value:   shape.Width,
			}
		} else if shape.Height <= 0 {
			return &ValidationError{
				Message: "高さが0以下です",
				Field:   "Height",
				Value:   shape.Height,
			}
		}

	case Circle:
		if shape.Radius <= 0 {
			return &ValidationError{
				Message: "半径が0以下です",
				Field:   "Radius",
				Value:   shape.Radius,
			}
		}

	default:
		return fmt.Errorf("対応していない図形です: %T", s)
	}

	return nil
}

func main() {
	// 処理対象の図形スライスを作成
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0}, // 不正（幅が不正）
		Circle{Radius: -1.0},                // 不正（半径が不正）
		Rectangle{Width: 1.0, Height: 0.0},  // 不正（高さが不正）
	}
	calcShape := []Shape{} //バリデーションチェック用スライス

	for _, shape := range shapes {
		err := validShape(shape)
		if err != nil {
			var validationErr *ValidationError
			if errors.As(err, &validationErr) {
				fmt.Printf("Validation Error: %s (Field: %s, Value: %f)\n", validationErr.Message, validationErr.Field, validationErr.Value)
			} else {
				fmt.Printf("Unexpected error: %v\n", err)
			}
			continue // 不正な図形はスキップ
		} else {
			calcShape = append(calcShape, shape)
		}
	}

	for _, shape := range calcShape {
		fmt.Printf(
			"図形: %T, 面積: %.2f, 周囲長: %.2f\n",
			shape,
			shape.Area(),
			shape.Perimeter(),
		)
	}
}
