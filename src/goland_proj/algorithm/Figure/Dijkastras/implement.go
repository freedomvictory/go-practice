package Dijkastras

import "log"

/*describe a node name , and the distance from One to it*/
type NodesAndDistances struct{
	Value map[string]int
}


//<SuanFaTuJie passage 108>
// Faster path to destination
// parameters:
// 1.input -> description of graph
// 2.start -> starting position
// 3.end -> ending position
// return value:
// 1.int -> distance
// 2. []string -> path map
func FasterestToDest(input map[string]NodesAndDistances) (int, []string){
	costs := make(map[string]int, 32)
	parents := make(map[string]string, 32)
	processed := []string{}

	for node_name, distance := range input["start"].Value{
		costs[node_name] = distance
		parents[node_name] = "start"
	}
	costs["fin"] = 255
	parents["fin"] = ""

	node_name, _ := GotLowCostNode(costs, processed)


	for{
		cost := costs[node_name]
		log.Printf("[FasterestToDest] (log) node_name:%v\n", node_name)
		for neighbor_nodeName, value := range input[node_name].Value{
			new_cost := cost + value
			if costs[neighbor_nodeName] > new_cost{
				log.Printf("[FasterestToDest] (log) neighbor_nodeNmae:%v update, value:%v\n", neighbor_nodeName, new_cost)
				costs[neighbor_nodeName] = new_cost
				parents[neighbor_nodeName] = node_name
			}
		}
		processed = append(processed, node_name)
		node_name, _ = GotLowCostNode(costs, processed)

		if node_name == "" || node_name == "fin"{
			break
		}

	}
	log.Printf("[FasterestToDest] (log) costs:%v, parents:%v\n", costs, parents)
	return 0, nil
}

func GotLowCostNode(input map[string]int, proced []string)(string, int){
	min := 255
	node_name := ""
	for key, val := range input{
		if key != "fin"{
			cost := val
			if cost < min && IsInProcessed(key, proced) == false{
				min = cost
				node_name = key
			}
		}
	}
	return node_name, -1
}

func IsInProcessed(node_name string, processed []string) bool{


	cnt_i := 0

	if len(processed) == 0{
		return false
	}

	for _, val := range processed{
		if node_name != val{
			cnt_i++
		}else{
			cnt_i = 0
			break
		}
	}
	if cnt_i != 0{
		return false
	}

	return true
}

