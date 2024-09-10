package structinterfaces

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{Width: 10.0, Height: 10.0}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Expected %f, Got %f", want, got)
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()

// 		got := shape.Area()

// 		if got != want {
// 			t.Errorf("Expected %f, Got %f", want, got)
// 		}
// 	}

// 	t.Run("Rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{Width: 12.0, Height: 6.0}

// 		checkArea(t, rectangle, 72.0)

// 	})

// 	t.Run("Circles", func(t *testing.T) {
// 		circle := Circle{Radius: 10.0}

// 		checkArea(t, circle, 314.1592653589793)

// 	})
// }

func TestArea(t *testing.T) {

	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	// for _, test := range areaTest {
	// 	got := test.shape.Area()
	// 	if got != test.want {
	// 		t.Errorf("Expected %f, Got %f", test.want, got)
	// 	}
	// }

	for _, test := range areaTest {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("Expected %f, Got %f", test.hasArea, got)
			}
		})
	}
}
