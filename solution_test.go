package square

import (
	"fmt"
	"testing"
)

func TestCalcSquare(t *testing.T) {
	testCases := []struct {
		sideLen  float64
		sidesNum intCustomType
		result   float64
	}{
		{10.0, SidesTriangle, 43.30127018922193},
		{10.0, SidesSquare, 100},
		{10.0, SidesCircle, 314.1592653589793},
		{10.0, intCustomType(5), 0.0},
	}

	errorcount := 0
	for _, test := range testCases {
		got := CalcSquare(test.sideLen, test.sidesNum)
		if test.result != got {
			t.Errorf("with %v and %v wait for %v but got %v", test.sideLen, test.sidesNum, test.result, got)
			errorcount++
		}
	}
	if errorcount == 0 {
		fmt.Println("passed")
	}

}

//U+1F5FA
