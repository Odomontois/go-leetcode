package golph

type UndergroundSystem [2]map[interface{}]interface{}

func Constructor() UndergroundSystem {
	return [2]map[interface{}]interface{}{{}, {}}
}
func (us *UndergroundSystem) CheckIn(i, s, t interface{}) {
	us[0][i] = [2]interface{}{s, t}
}
func (us *UndergroundSystem) CheckOut(i, s, t interface{}) {
	place := us[0][i].([2]interface{})
	route := [2]string{place[0].(string), s.(string)}
	state, _ := us[1][route].([2]int)
	us[1][route] = [2]int{state[0] + t.(int) - place[1].(int), state[1] + 1}
}
func (us *UndergroundSystem) GetAverageTime(s, e string) float64 {
	state := us[1][[2]string{s, e}].([2]int)
	return float64(state[0]) / float64(state[1])
}
