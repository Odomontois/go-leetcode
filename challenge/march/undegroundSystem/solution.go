package undegroundSystem

type Place struct {
	current string
	depart  int
}
type UndergroundSystem struct {
	current map[int]Place
	stat    map[[2]string][2]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{make(map[int]Place), make(map[[2]string][2]int)}
}
func (us *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	us.current[id] = Place{stationName, t}
}
func (us *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	place := us.current[id]
	delete(us.current, id)
	route := [2]string{place.current, stationName}
	state := us.stat[route]
	us.stat[route] = [2]int{state[0] + t - place.depart, state[1] + 1}
}
func (us *UndergroundSystem) GetAverageTime(start string, end string) float64 {
	state := us.stat[[2]string{start, end}]
	return float64(state[0]) / float64(state[1])
}
