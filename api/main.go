package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"
)

type Skill struct {
	Name        string    `json:"skill"`
	Complexity  float64   `json:"complexity"` 
	LastWorked  time.Time `json:"lastWorked"`
}


func handleData(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		decoder := json.NewDecoder(r.Body)
		skillStruct := Skill{}
		err := decoder.Decode(&skillStruct)
		if err != nil{
			fmt.Printf("Error Decoding Body: %s", err)
			w.WriteHeader(400)
			return 
		}

		//We are handling data here and passing it to our calculating function to get an interger type "score"
		regressorScore, err := regressorCalculator(skillStruct)
		if err != nil{
			fmt.Printf("Error calculating score %s: ", err)
			w.WriteHeader(400)
			return
		}

		//Now that we have the regressor score, the higher the score, the more information retained
		if regressorScore >= 8 {

		} else if regressorScore >= 4 && regressorScore <=6 {

		} else if regressorScore >= 1 && regressorScore <= 3{

		} else{

		}
		

		w.WriteHeader(200)

	} else{
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Printf("Method not allowed")
	}
}

//The calculations are based off of Ebbinghaus' forgetting curve
func regressorCalculator(s Skill) (int, error){
	skillSec:= time.Since(s.LastWorked)
	skillDays := skillSec.Hours()/24

	skillComplexity := s.Complexity

	baseScoreConstant := float64(10)
	decayRateConstant := 0.02


	retention := (baseScoreConstant * math.Exp(-decayRateConstant * skillDays * skillComplexity))

	return int(retention), nil
}

func main(){

	http.HandleFunc("/", handleData)

	err := http.ListenAndServe(":8080", nil)
	if err !=nil{
		fmt.Printf("Error starting server %s", err)
		return 
	}
}