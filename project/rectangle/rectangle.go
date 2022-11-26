package figure

import "math"

type Shape interface { // インターフェース→継承とかせず、暗黙的に解決されるらしい
	Perimeter() float64
	Area() float64
}

type Rectangle struct { // 構造体定義
	Width  float64
	Height float64
}

// メソッド
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
func (r Rectangle) Area() float64      { return r.Width * r.Height }

type Circle struct { // 構造体定義
	Radius float64
}
 //メソッド
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }


type Triangle struct {
	Base   float64
	Height float64
}
//メソッド
func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height + math.Sqrt(math.Pow(t.Base, 2) + math.Pow(t.Height, 2))
}
func (t Triangle) Area() float64 { return t.Base * t.Height / 2 }

func Perimeter(rectangle Rectangle) float64 {
	return 2.0 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
