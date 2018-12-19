package jastMovAvg

import(

)

func getAvg(list []float64) (float64) {
	var acc = float64(0)

	for _, elem := range list {
		acc += elem
	}


	return acc / float64(len(list))
}

func MovingAverage(list []float64, window int) ([]float64) {
	ret := []float64{}
	helper := []float64{}


	for i, _ := range list {

		for j := 0 ; j < window; j++ {
			if i + j < len (list) {
				helper = append(helper, list[i + j])
			} else if i + j >= len(list) {
				helper = append(helper, list[i-j])
			}
		}

		a := getAvg(helper)
		ret = append(ret, a)

		helper = nil

	}

	return ret

}
