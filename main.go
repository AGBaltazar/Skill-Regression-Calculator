package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
	"net/url"

	"github.com/joho/godotenv"
)

type Skill struct {
	Name        string    `json:"skill"`
	Complexity  string `json:"complexity"` 
	LastWorked  string `json:"lastpracticed"`
}

type Response struct{
	RegressorScore int `json:"regressorscore"`
	Interpretation string `json:"interpretation"`
	YoutubeQuery []string `json:"youtubelinks"`
}
type YoutubeSearchResult struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
	} `json:"items"`
}


func handleData(w http.ResponseWriter, r *http.Request){
	apiKey := os.Getenv("API_KEY")
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

		//Now that we have the regressor score, we will be sending it back to our front end along with a interpretation and recommendations
		if regressorScore >= 7 {
			query := "?part=snippet&type=video&maxResults=5&q=" + url.QueryEscape(skillStruct.Name+" advanced concepts") + "&key=" + apiKey

			youtubeResponse, err := callYoutube(query)
			if err != nil{
				fmt.Printf("Error calling YouTube: %s ", err)
			}

			response := Response{
				RegressorScore: regressorScore,
				Interpretation: "Fresh",
				YoutubeQuery: youtubeResponse,
			}
			fmt.Printf("Youtube response List: %v", youtubeResponse)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		} else if regressorScore >= 4 && regressorScore <=6 {
			query := "?part=snippet&type=video&maxResults=5&q=" + url.QueryEscape(skillStruct.Name+" intermediate concepts") + "&key=" + apiKey

			youtubeResponse, err := callYoutube(query)
			if err != nil{
				fmt.Printf("Error calling YouTube: %s ", err)
			}

			response := Response{
				RegressorScore: regressorScore,
				Interpretation: "Rusty",
				YoutubeQuery: youtubeResponse,
			}
			fmt.Printf("Youtube response List: %v", youtubeResponse)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		} else if regressorScore >= 1 && regressorScore <= 3{
			query := "?part=snippet&type=video&maxResults=5&q=" + url.QueryEscape(skillStruct.Name+" basic concepts") + "&key=" + apiKey

			youtubeResponse, err := callYoutube(query)
			if err != nil{
				fmt.Printf("Error calling YouTube: %s ", err)
			}

			response := Response{
				RegressorScore: regressorScore,
				Interpretation: "Weak",
				YoutubeQuery: youtubeResponse,
			}
			fmt.Printf("Youtube response List: %v", youtubeResponse)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)

		} else{

			query := "?part=snippet&type=video&maxResults=5&q=" + url.QueryEscape(skillStruct.Name+" beginner tutorial") + "&key=" + apiKey

			youtubeResponse, err := callYoutube(query)
			if err != nil{
				fmt.Printf("Error calling YouTube: %s ", err)
			}

			responseStruct := Response{
				RegressorScore: regressorScore,
				Interpretation: "Forgotten",
				YoutubeQuery: youtubeResponse,
			}
			fmt.Printf("Youtube response List: %v", youtubeResponse)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseStruct)
		}

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

	retention := (baseScoreConstant * math.Exp(-decayRateConstant * skillDays * float64(skillComplexity)))

	return int(retention), nil
}

///Function takes in the URL ready with parameters aand will rcall to Youtube to get the links of the recommended videos
func callYoutube(queryURL string) (youtubeResponse []string, err error){
	
	apiLink := "https://www.googleapis.com/youtube/v3/search" + queryURL
	response, err := http.Get(apiLink)
	if err != nil{
		fmt.Printf("Error fetching YouTube content: %s", err)
		emptySlice := make([]string, 3)
		return emptySlice, err
	}

	decoder := json.NewDecoder(response.Body)
	responseStruct := YoutubeSearchResult{}
	err = decoder.Decode(&responseStruct)
	if err != nil{
		fmt.Printf("Error decoding %s: ", err)
		return nil, err
	}

	//We will look through the returning data and only take 5 or less 
	videos := []string{}
	for i, item := range responseStruct.Items {
		if i >= 5 {
        	break
    }
    if item.ID.VideoID != "" {
	videoUrl := "https://www.youtube.com/watch?v=" + item.ID.VideoID
	videos = append(videos, videoUrl)
}

}

	return videos, nil
}

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	mux := http.FileServer(http.Dir("static/"))
	http.HandleFunc("/api", handleData)
	http.Handle("/", mux)

	

	err = http.ListenAndServe(":8080", nil)
	if err !=nil{
		fmt.Printf("Error starting server %s", err)
		return 
	}
}