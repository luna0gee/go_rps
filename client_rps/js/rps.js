(function rps(){
    var move_selection;

    //****************************************
    // Game logic: provides an API for accessing rps game logic and submitting request to server
    //****************************************

    var rps_logic = (function(){
        //Game logic variables
        var balance = 1000;

        //"Public"
        function submit_move(new_move){
            //Get the info
            new_move.User_Data = get_user_info()

            //Build JSON
            var new_move_json = JSON.stringify(new_move)
            //Send Request
            $.ajax({
                type        :   "POST",
                url         :   "/post",
                dataType    :   "json",
                success     :   function(data, status, jqxhr){game_results(data, status, jqxhr)},
                complete    :   function(){view_stuff.game_requested()},
                error       :   function(){alert("Game Request Failed")},
                data        :   new_move_json
            })
        }

        function get_balance(){
            return balance;
        }

        //"Private"
        function game_results(data, status, jqxhr){
            console.log("Received Results")
            console.log(data)
            console.log(status)

            var results = JSON.parse(data)
            console.log(results)

            update_balance(results.Bet, results.Opp_Bet, results.Result)
            view_stuff.show_results()
        }

        function update_balance(bet, opp_bet, result){
            if (bet > balance){}
            switch (result) {
                case "win":
                    balance += opp_bet
                    balance += bet
                case "lose":
                    balance -= bet
            }

        }

        function get_user_info(){
            return {
                user_agent  :   navigator.userAgent,
                platform    :   navigator.platform,
                cookies     :   navigator.cookieEnabled,
            }
        }

        //Return references to "public" functions
        return{
            submit_move : function(new_move){return submit_move(new_move)},
            get_balance : function(){return get_balance()}
        }
    }())

    //****************************************
    //Elements Interface for jQuery objects / HTML elements
    //****************************************

    var elements = {
        name_text       :   $("#name_text"),
        name_button     :   $("#name_button"),
        name_form       :   $("#name_form   "),
        name_form_input :   $("#name_form_input"),
        balance_text    :   $("#balance_text"),
        bet_form        :   $("#bet_form"),
        bet_warning     :   $("#bet_warning"),
        rps_div         :   $("#rps_div"),
        rock_button     :   $("#rock_button"),
        paper_button    :   $("#paper_button"),
        scissors_button :   $("#scissors_button"),
        make_move       :   $("#make_move"),
        submit_msg      :   $("#submit_msg"),
        results_div     :   $("#results_div"),
        play_again      :   $("#play_again")
    }

    //****************************************
    //View Stuff: convenient methods to manipulate elements on page: show/hides & listeners
    //****************************************

    var view_stuff = {
        initialize      :   function(){
            view_stuff.change_balance()
            elements.make_move.click(view_stuff.submit_move)
            elements.name_button.click(view_stuff.show_name_change)
            elements.rock_button.click(function(){view_stuff.make_selection(0)})
            elements.paper_button.click(function(){view_stuff.make_selection(1)})
            elements.scissors_button.click(function(){view_stuff.make_selection(2)})
            elements.play_again.click(view_stuff.new_game)
            elements.name_form_input.hide()
            elements.bet_warning.hide()
            elements.results_div.hide()
            elements.submit_msg.hide()
        },

        submit_move     :   function(){
            console.log("Submitting Move")
            var new_bet = elements.bet_form.val()
            if( new_bet == "" || parseInt(new_bet) > rps_logic.get_balance() || parseInt(new_bet) < 1){view_stuff.bad_bet();}
            else{view_stuff.ok_bet()}
            var new_move = {
                Move        : move_selection,
                Bet         : new_bet,
                User_Data   : "",
                Match_Data  : null
            }
            elements.make_move.off()
            rps_logic.submit_move(new_move)
        },

        bad_bet         :   function(){
            elements.bet_form.addClass('bad_bet')
            elements.bet_warning.show()
        },

        ok_bet          :   function(){
            elements.bet_form.removeClass('bad_bet')
            elements.bet_warning.hide()
        },

        make_selection  :   function(selection){
            elements.rock_button.removeClass('selection')
            elements.paper_button.removeClass('selection')
            elements.scissors_button.removeClass('selection')
            switch(selection){
                case 0:
                    elements.rock_button.addClass('selection')
                    move_selection = "rock"
                    break
                case 1:
                    elements.paper_button.addClass('selection')
                    move_selection = "paper"
                    break
                case 2:
                    elements.scissors_button.addClass('selection')
                    move_selection = "scissors"
                    break
                default:
                    break
            }
        },

        show_name_change:   function(){
            elements.name_form_input.show()
            elements.name_button.off()
            elements.name_button.click(view_stuff.change_name)
        },

        change_name     :   function(){
            console.log(elements.name_form_input.text())
            elements.name_text.text(elements.name_form_input.val())
            elements.name_form_input.val("")
            elements.name_form_input.hide()
            elements.name_button.off()
            elements.name_button.click(view_stuff.show_name_change)
        },

        change_balance  :   function(){
            console.log(rps_logic.get_balance())
            elements.balance_text.text(rps_logic.get_balance())
        },

        game_requested  :   function(){
            elements.submit_msg.show()
        },

        show_results    :   function(){
            elements.submit_msg.hide()
            elements.results_div.show()
        },

        new_game        :   function(){
            elements.play_again.hide()
            elements.results_div.hide()
            view_stuff.make_selection(-1)
            elements.make_move.click(view_stuff.submit_move)
        }
    }

    view_stuff.initialize()
})();

