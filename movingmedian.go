package jastMovAvg

import(
	"sort"
)

func getMedian(list []float64) (float64) {

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	if len(list) % 2 == 0 {
		return (list[(len(list)/2)] + list[(len(list)/2) - 1]) / float64(2)
	} else {
		return list[(len(list)/2)]
	}
}

func MovingMedian(list []float64, window int) ([]float64) {
	ret := []float64{}
	helper := []float64{}


	for i, _ := range list {

		for j := 0 ; j < window; j++ {
			if i + j < len (list) {
				helper = append(helper, list[i + j])
			}
		}

		a := getMedian(helper)
		ret = append(ret, a)

		helper = nil

	}

	return ret

}
