package go_rps_logic

import "testing"

func init(){
	testing_env = true
}

func Test_game_logic(test *testing.T){
	player_one := Game_Request{
		Move		:	"rock",
		Bet       	:	"606",
		Opp_Move	:	"",
		Opp_Bet		:	"",
		UserData  	:	"stuff",
		Result		:	""}
	player_two := Game_Request{
		Move		:	"paper",
		Bet       	:	"616",
		Opp_Move	:	"",
		Opp_Bet		:	"",
		UserData  	:	"stuff",
		Result		:	""}

	run_game(&player_one, &player_two)
	if player_one.Result != "lose" || player_two.Result != "win"{
		test.Error("Rock should beat paper, instead p1 (rock) result was ", player_one.Result, "and p1 (rock) result was ", player_two.Result)
	}
	//TODO test for draw
	if player_one.Opp_Bet != "616" || player_two.Opp_Bet != "606"{
		test.Error("Betting amounts not traded properly p1.Opp_Bet exp: 666, actual:", player_one.Opp_Bet, "p2.Opp_Bet exp: 616, actual:", player_one.Opp_Bet)
	}
}
