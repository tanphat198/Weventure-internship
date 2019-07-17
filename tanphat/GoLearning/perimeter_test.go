package GoLearning

import "testing"


func TestPerimeter(t *testing.T){
	got := Perimeter(Rectangle{10.0,10.0})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f",got ,want)
	}
}

func TestArea(t *testing.T){
	areaTests := []struct{
		shape Shape
		want  float64
	}{
		{Rectangle{12,6},72},
		{Circle{10},314.1592653589793},
		{Triangle{12,6},36},
	}

	for _, tt := range areaTests{
		got := Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f",got ,tt.want)
		}
	}
}

