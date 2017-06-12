package go_rps_logic

import (
	"net/http"
	"log"
	"encoding/json"
	"time"
	"fmt"
)
var quit bool = false

func Start_game(resp http.ResponseWriter, req *http.Request){
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

	resp.Header().Set("Content-Type", "application/json")

	//queue game and wait for response
	go_play_rps(&new_game)

	//results, err := json.Marshal(new_game)
	//if err != nil{
	//	http.Error(resp, "Game Processing Error", 400)
	//}

	json.NewEncoder(resp).Encode(new_game)
	fmt.Println("heyo")
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

func Go_rps(){
	//Start channels in independent process
	go match_maker()
	go note_taker()

	//TODO refactor routing with mux
	//Declare Routing
	serve_dir := http.FileServer(http.Dir("./client_rps"))
	http.Handle("/", serve_dir)
	http.HandleFunc("/post", Start_game)

	//Start Listening
	log.Println("derp")
	http.ListenAndServe(":8080", nil)

	//Exit
	quit = true
}
