package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type HotDryNoodles struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Brand *Brand `json:"brand"`
}

var HotDryNoodlesList []HotDryNoodles

func initFakeData() {
	// 热干面 6元 蔡林记
	// 大碗热干面 8元 蔡林记
	// 牛肉热干面 10元 蔡林记
	HotDryNoodlesList = append(HotDryNoodlesList, HotDryNoodles{ID: 1, Name: "热干面", Price: 6, Brand: &Brand{ID: 1, Name: "蔡林记"}})
	HotDryNoodlesList = append(HotDryNoodlesList, HotDryNoodles{ID: 2, Name: "大碗热干面", Price: 8, Brand: &Brand{ID: 1, Name: "蔡林记"}})
	HotDryNoodlesList = append(HotDryNoodlesList, HotDryNoodles{ID: 3, Name: "牛肉热干面", Price: 10, Brand: &Brand{ID: 1, Name: "蔡林记"}})
}
func main() {
	r := mux.NewRouter()
	initFakeData()

	//api/noodles
	r.HandleFunc("/api/noodles", GetNoodles).Methods("GET")
	r.HandleFunc("/api/noodles/{id}", GetNoodle).Methods("GET")
	r.HandleFunc("/api/noodles", CreateNoodle).Methods("POST")
	r.HandleFunc("/api/noodles/{id}", UpdateNoodle).Methods("PUT")
	r.HandleFunc("/api/noodles/{id}", DeleteNoodle).Methods("DELETE")

	log.Println("Starting server on port 8000", "open http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func DeleteNoodle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range HotDryNoodlesList {
		if item.ID == toInt(params["id"]) {
			HotDryNoodlesList = append(HotDryNoodlesList[:index], HotDryNoodlesList[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(writer).Encode(HotDryNoodlesList)
	if err != nil {
		panic(err)
	}
}

func UpdateNoodle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range HotDryNoodlesList {
		if item.ID == toInt(params["id"]) {
			HotDryNoodlesList = append(HotDryNoodlesList[:index], HotDryNoodlesList[index+1:]...)
			var noodle HotDryNoodles
			_ = json.NewDecoder(request.Body).Decode(&noodle)
			noodle.ID = toInt(params["id"])
			HotDryNoodlesList = append(HotDryNoodlesList, noodle)
			err := json.NewEncoder(writer).Encode(noodle)
			if err != nil {
				panic(err)
			}
			return
		}
	}
	err := json.NewEncoder(writer).Encode(HotDryNoodlesList)
	if err != nil {
		panic(err)
	}

}

func CreateNoodle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var noodle HotDryNoodles
	_ = json.NewDecoder(request.Body).Decode(&noodle)
	HotDryNoodlesList = append(HotDryNoodlesList, noodle)
	err := json.NewEncoder(writer).Encode(noodle)
	if err != nil {
		panic(err)
	}

}

func GetNoodle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range HotDryNoodlesList {
		if item.ID == toInt(params["id"]) {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	err := json.NewEncoder(writer).Encode(&HotDryNoodles{})
	if err != nil {
		panic(err)
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func GetNoodles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(HotDryNoodlesList)
}
