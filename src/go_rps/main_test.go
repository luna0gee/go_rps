package main

import (
	"testing"
	"net/http"
	"encoding/json"
	"strings"
	"go_rps/go_rps_logic"
	"io/ioutil"
)

const TEST_URL = "http://localhost:8080/post"

func init(){
	go main()
}

func Test_server_response(test *testing.T){
	//Build two game request objects
	player_one := go_rps_logic.Game_Request{
		Move		:	"rock",
		Bet       	:	"606",
		Opp_Move	:	"",
		Opp_Bet		:	"",
		UserData  	:	"stuff",
		Result		:	""}
	player_two := go_rps_logic.Game_Request{
		Move		:	"paper",
		Bet       	:	"616",
		Opp_Move	:	"",
		Opp_Bet		:	"",
		UserData  	:	"stuff",
		Result		:	""}

	p1_request_json, err := json.Marshal(player_one)
	if err!=nil {
		test.Fatal(err)
	}
	p2_request_json, err := json.Marshal(player_two)
	if err!=nil {
		test.Fatal(err)
	}

	p1_reader := strings.NewReader(string(p1_request_json))
	p2_reader := strings.NewReader(string(p2_request_json))

	//Queue two game requests, p1 asynchronously, p2 wait for response and use for assertions
	go http.Post(TEST_URL, "application/json", p1_reader)
	p2_post, err := http.Post(TEST_URL,"application/json", p2_reader)
	if err !=nil {
		test.Fatal(err)
	}

	//Get the response json, update player_two Game Request
	p2_response_json, err := ioutil.ReadAll(p2_post.Body)
	if err != nil{
		test.Fatal(err)
	}
	json.Unmarshal(p2_response_json, &player_two)

	//Assert win/loss conditions
	if player_two.Result != "win"{ test.Fatal("Player Two should have won") }

	//Assert correct betting logic
	if player_two.Opp_Bet != "606"{ test.Fatal("Player Two opponent bet value incorrect") }
	if player_two.Bet != "616"{ test.Fatal("Player Two's bet value incorrect") }
}

func Test_stress(test *testing.T){

	//Queue requests

	//Assert responses recieved

	//Record output & benchmarks
}
