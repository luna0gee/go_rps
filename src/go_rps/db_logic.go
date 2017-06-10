package main

var match_records = make(chan *Match_Record, 100)

//TODO nested json
type Match_Record struct{
	player_one string
	player_two string
	game_results string
}

func note_taker(){
	//TODO ensure dealing with finishing commits in channel after quit signal
	for !quit{

	}
}

func store_results(){
	//
}

