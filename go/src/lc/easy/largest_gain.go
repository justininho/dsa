package easy

/*
(1732 | Easy) largestAltitude
*/
func largestAltitude(gain []int) int {
	altitude := 0
	peak := 0
	for _, g := range gain {
		altitude += g
		peak = max(peak, altitude)
	}
	return peak
}
