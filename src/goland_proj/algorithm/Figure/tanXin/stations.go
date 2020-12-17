package tanXin


import (
	"fmt"
	"github.com/adam-hanna/arrayOperations"

)



//《算法图解》passage 123

func SetCover(station_list map[string][]string, states_needed []string)([]string ,error) {

	finalStations := []string{}
	state_covered := []string{} //
	best_station := ""

	for len(states_needed) > 0 {

		state_covered = []string{}

		for station, state_for_station := range station_list {
			covered, ok := arrayOperations.Intersect(states_needed, state_for_station)
			if !ok {
				continue
			}
			covered_sa := covered.Interface().([]string)
			if len(covered_sa) > len(state_covered) {
				best_station = station
				state_covered = covered_sa
			}
		}

		out, ok := arrayOperations.Difference(states_needed, state_covered)

		if !ok {
			return nil, fmt.Errorf("[SetCover] Failed: difference uknown error")
		}
		states_needed = out.Interface().([]string)
		//fmt.Printf("[SetCover] (debug) states_needed:%v\n", states_needed)
		finalStations = append(finalStations, best_station)

	}
	return finalStations, nil

}

