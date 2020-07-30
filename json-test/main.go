package main

import (
	"encoding/json"
	"log"
)

type Power_t struct {
	Value float32 `json:"value"`
	Time  string  `json:"time"`
}

type Params_t struct {
	Power Power_t `json:"power"`
}

type TempSerias struct {
	Id      string   `json:"id"`
	Version string   `json:"version"`
	Params  Params_t `json:"params"`
	Method  string   `json:"method"`
}

func main() {

	var Power_inst = Power_t{
		Value: 23.6,
		Time:  "1234234235",
	}

	var Params_inst = Params_t{
		Power: Power_inst,
	}

	var temp_data = TempSerias{
		Id:      "123",
		Version: "1.0",
		Params:  Params_inst,
		Method:  "thing.event.property.post",
	}

	json_data, err := json.MarshalIndent(temp_data, "", "  ")
	if err != nil {
		log.Fatalf("[main] (error)Json Marshal failed : %v", err)
	}
	log.Printf("[main] (log) Json Marshal data:\n%s\n", json_data)
}
