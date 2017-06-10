package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"time"
)
var quit bool = false

func start_game(resp http.ResponseWriter, req *http.Request){
	if req.Body == nil{
		http.Error(resp, "Bad Request", 400)
		return
	}

	//TODO nested json
	var new_game Game_Request
	parse_error := json.NewDecoder(req.Body).Decode(&new_game)

	if parse_error != nil {
		http.Error(resp, "Bad Request", 400)
		return
	}

	resp.Header().Set("Content'Type", "application/json")

	//queue game and wait for response
	go_play_rps(&new_game)

	results, err := json.Marshal(new_game)
	if err != nil{
		http.Error(resp, "Game Processing Error", 400)
	}
	fmt.Println(resp, results)
}

func go_play_rps(new_game * Game_Request){
	matches <- new_game
	for {
		if(new_game.Result != ""){
			return
		}
		time.Sleep(1000)
	}
}

func main(){
	//Start channels in independent process
	go match_maker()
	go note_taker()

	//Declare Routing
	serve_dir := http.FileServer(http.Dir("./client_rps"))
	http.Handle("/", serve_dir)
	http.HandleFunc("/post", start_game)

	//Start Listening
	log.Println("derp")
	http.ListenAndServe(":8080", nil)

	//Exit
	quit = true
}
