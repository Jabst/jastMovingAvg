package jastMovAvg

import (
	// "testing"
)

func main() {
	list := []float64{1,4,2,3,5,3,2,1,2,4,5,2,4,1,5,6,2,7,3,2,1,2,5,2,3,5,4,2,65,2,1,3,2,3,5,3,6,1}

	preCalc := []float64{3.2,3.2,3.1,3.25,3.2,3.1,3.2,3.3,3.35,6.5,6.4,6.2,6.25,6.15,6.25,6.25,6.1,6.3,6,6.157894736842105,6.388888888888889,6.705882352941177,7,7.133333333333334,7.5,7.846153846153846,8.083333333333334,8.454545454545455,9.1,2.888888888888889,3,3.2857142857142856,3.3333333333333335,3.6,3.75,3.3333333333333335,3.5,1}

	a :=  MovingAverage(list, 20)

	for idx, elem := range a {
		if elem != preCalc[idx]{
			// testing.Fail()
		}

	}

}
