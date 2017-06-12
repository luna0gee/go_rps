package go_rps_logic

import (
	"database/sql"
	"fmt"
	"time"
	"encoding/json"
	_"github.com/lib/pq"
)

var match_records = make(chan *Game_Record, 100)
var testing_env bool = false

const PGDB_ADMIN = "go_rps_admin"
const PGDB_PASS = "g0_rp5_g0"
const PGDB = "go_rps"

//TODO nested json
type Game_Record struct{
	player_one string
	player_two string
}

func note_taker(){
	//TODO ensure dealing with finishing commits in channel after quit signal
	var game_record * Game_Record;
	for !quit{
		for len(match_records) > 0{
			game_record = <- match_records
			store_results(game_record)
		}
		time.Sleep(1000)
	}
}

func store_results(game_record * Game_Record){
	//open connection
	results := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		PGDB_ADMIN, PGDB_PASS, PGDB)
	pg_db, err := sql.Open("postgres", results)

	//close if error
	if err != nil{
		panic(err)
	}
	defer pg_db.Close()


	//close the connection
	pg_db.Close()
}

func queue_game_record(player_one * Game_Request, player_two * Game_Request){
	//convert player one and two to json
	p1_record, err := json.Marshal(player_one)
	if err != nil{
		panic(err)
	}
	fmt.Println(len(string(p1_record)))

	p2_record, err := json.Marshal(player_two)
	if err != nil{
		panic(err)
	}
	fmt.Println(len(string(p2_record)))

	//build Game Record
	var game_record Game_Record
	game_record.player_one = string(p1_record)
	game_record.player_two = string(p2_record)

	//push to channel
	if !testing_env {
		match_records <- &game_record
	}
}