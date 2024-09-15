package main

var feeSlice []int
var (
	MaxChargedWeight = 100
	StartingFee      = 18
	IncreFeePerKg    = 5
)

func init() {
	countCourierFee()
}
func countCourierFee() {
	feeSlice = make([]int, MaxChargedWeight+1)
	feeSlice[0] = StartingFee
	for i := 1; i <= MaxChargedWeight; i++ {
		fee := upOrDown(float64(feeSlice[i-1])*0.01) + StartingFee + IncreFeePerKg*i
		feeSlice[i] = fee
	}

}
func GetCourierFee(weight float64) (fee int) {
	if weight == 0 {
		return 0
	}
	if weight <= 1 {
		return feeSlice[0]
	} else if weight > float64(MaxChargedWeight) {
		return feeSlice[len(feeSlice)-1]
	} else if float64(int(weight)) == weight {
		return feeSlice[int(weight)-1]
	}
	return feeSlice[int(weight)]
}
func upOrDown(f float64) int {
	down := int(f)
	up := down + 1
	if f < float64(down)+0.5 {
		return down
	} else {
		return up
	}
}
