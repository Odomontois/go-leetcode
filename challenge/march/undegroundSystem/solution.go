package undegroundSystem

type Route struct {
	from, to string
}

type Place struct {
	current string
	depart  int
}

type State struct {
	sum, count int
}

type UndergroundSystem struct {
	current map[int]Place
	stat    map[Route]State
}

func Constructor() (sys UndergroundSystem) {
	sys.current = make(map[int]Place)
	sys.stat = make(map[Route]State)
	return
}

func (us *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	us.current[id] = Place{stationName, t}
}

func (us *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	place := us.current[id]
	delete(us.current, id)
	route := Route{place.current, stationName}
	state := us.stat[route]
	state.sum += t - place.depart
	state.count++
	us.stat[route] = state
}

func (us *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	state := us.stat[Route{startStation, endStation}]
	return float64(state.sum) / float64(state.count)
}
