package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"
	"strconv"
)

type Skill struct {
	Name        string    `json:"skill"`
	Complexity  string `json:"complexity"` 
	LastWorked  string `json:"lastpracticed"`
}

type Response struct{
	RegressorScore int `json:"regressorscore"`
	Interpretation string `json:"interpretation"`
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

		//Now that we have the regressor score, we will be sending it back to our front end along with a interpretatino
		if regressorScore >= 8 {
			response := Response{
				RegressorScore: regressorScore,
				Interpretation: "Fresh",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		} else if regressorScore >= 4 && regressorScore <=6 {
			response := Response{
				RegressorScore: regressorScore,
				Interpretation: "Rusty",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		} else if regressorScore >= 1 && regressorScore <= 3{

			response := Response{
				RegressorScore: regressorScore,
				Interpretation: "Weak",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		} else{
			responseStruct := Response{
				RegressorScore: regressorScore,
				Interpretation: "Forgotten",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseStruct)
		}
		

		w.WriteHeader(200)

	} else{
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Printf("Method not allowed")
	}
}

//The calculations are based off of Ebbinghaus' forgetting curve
//Struct Strings need to be converted to the respective date as well as integers
func regressorCalculator(s Skill) (int, error){
	
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, s.LastWorked)
	if err != nil{
		fmt.Printf("Error parsing date: %s", err)
	}
	skillTime:= time.Since(parsedDate)
	skillDays := skillTime.Hours()/24

	skillComplexity, err := strconv.Atoi(s.Complexity)
	if err != nil{
		fmt.Printf("Error converting string to intefer: %s: ", err)
	}

	baseScoreConstant := float64(10)
	decayRateConstant := 0.002

	fmt.Printf("Skill day %v\n", skillDays)
	retention := (baseScoreConstant * math.Exp(-decayRateConstant * skillDays * float64(skillComplexity)))

	return int(retention), nil
}

func main(){
	mux := http.FileServer(http.Dir("static/"))
	http.HandleFunc("/api", handleData)
	http.Handle("/", mux)

	err := http.ListenAndServe(":8080", nil)
	if err !=nil{
		fmt.Printf("Error starting server %s", err)
		return 
	}
}