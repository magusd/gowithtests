package structs

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{20.0, 40.0}, want: 120.0},
		{shape: Circle{10}, want: 3.141592653589793 * 2 * 10},
		//{Triangle{12, 6}, ???},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("failed to calculate the perimeter or %#v, got %g, want %g", tt.shape, got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{20.0, 40.0}, hasArea: 800.0},
		{shape: Circle{10.0}, hasArea: 314.1592653589793},
		{shape: Triangle{12, 6}, hasArea: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("failed to calculate the area of %#v, got %g, want %g", tt.shape, got, tt.hasArea)
		}
	}
}
