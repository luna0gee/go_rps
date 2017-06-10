package main

import (
	"time"
	"strings"
)

var matches = make(chan *Game_Request, 100)

//TODO nested json
type Game_Request struct {
	Move      string
	Bet       string
	Opp_Move string
	Opp_Bet string
	UserData  string
	Result	string
}

func match_maker(){
	var first_request *Game_Request
	var second_request *Game_Request
	
	//TODO More elegant solution: Worker queue, dynamic generation of buffered channels
	//Loop until exit signal (quit boolean in go_rps)
	for !quit{
		for len(matches) > 1{
			first_request = <- matches
			second_request = <- matches
			run_game(first_request, second_request)
		}
		time.Sleep(1000)
	}
}

func run_game(player_one * Game_Request, player_two * Game_Request){
	var result = fight(player_one.Move, player_two.Move)
	
	switch result{
	case -1:
		player_one.Result = "lose"
		player_two.Result = "win"
		break
	case 0:
		player_one.Result = "draw"
		player_two.Result = "draw"
		break
	case 1:
		player_one.Result = "win"
		player_two.Result = "lose"
		break
	case -2:
		player_one.Result = "wat"
		player_two.Result = "wat"
		break
	}
	store_results()
}

func fight(p1_move string, p2_move string) int {
	//TODO refactor w/ ints
	if strings.Compare(p1_move,p2_move) == 0 {return 0}
	switch p1_move{
	case "rock":
		switch p2_move{
		case "paper":
			return -1
		case "scissors":
			return 1
		}
	case "scissors":
		switch p2_move{
		case "rock":
			return -1
		case "paper":
			return 1
		}
	case "paper":
		switch p2_move{
		case "scissors":
			return -1
		case "rock":
			return 1
		}
	default: break
	}
	return -2
}


