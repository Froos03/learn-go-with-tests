package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.5, 20.0}
	actual := rectangle.Perimeter()
	expected := 61.0

	if actual != expected {
		t.Errorf("Expected %.2f but actual is %.2f", expected, actual)
	}

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{5.0, 5.0}, hasArea: 25.0},
		{name: "Circle", shape: Circle{1}, hasArea: 3.141592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Area()
			if actual != tt.hasArea {
				t.Errorf("%#v Expected %g but actual is %g", tt.name, tt.hasArea, actual)
			}
		})
	}

}
