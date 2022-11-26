package figure

import (
	"testing"
)

func TestPerimeter(t *testing.T){
	perimeterTests := []struct { // 構造体のスライスを宣言
		shape Shape
		want float64
	}{
		{shape: Rectangle{Width: 10, Height:10}, want:40.0},
		{shape: Circle{Radius:10}, want:62.83185307179586},
		{shape: Triangle{Base: 10, Height:10}, want:34.14213562373095},
	}

	for _, tt := range perimeterTests{
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("%#v got %g but want %g", tt.shape, tt.want, got)
		}
	}
}

func TestArea(t *testing.T) { // テーブル駆動テスト
	areaTests := []struct { // 構造体のスライスを宣言
		shape Shape
		want float64
	}{
		{shape: Rectangle{Width: 10, Height:10}, want:100.0},
		{shape: Circle{Radius:10}, want:314.1592653589793},
		{shape: Triangle{Base: 10, Height:10}, want:50.0},
	}

	for _, tt := range areaTests{
		got := tt.shape.Area()
		if got != tt.want { // %#v フィールドを含む構造体を出力
			t.Errorf("%#v got %g but want %g", tt.shape, tt.want, got)
		}
	}
}
// func TestArea(t *testing.T) {
// 	checkArea := func(t *testing.T, shape Shape, want float64){
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("expecetd %g but got %g", want, got)
// 		}
// 	}

// 	t.Run("rectangle area", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		want := 100.0
// 		checkArea(t, rectangle, want)
// 	})

// 	t.Run("circle area", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		want := 314.1592653589793
// 		checkArea(t, circle, want)
// 	})
// }
