package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"golang.org/x/exp/slices"
)




func gameLogic(user_input string, ran_num int, smallest int, greatest int,loop_iterator int, guess_max int) (string,bool) {
	

	con_guess, err := strconv.Atoi(user_input)
	

	if err != nil && loop_iterator != 4{
		fmt.Println("\nThat was an invalid input. Intergers only")
	}
	// || is or operator && is and ! is not operator
	if con_guess < smallest || con_guess > greatest {
		return fmt.Sprintf("\nThat is an Invalid Input. Cannot be less than %[1]d nor greater than %[2]d\n", smallest, greatest),false
	}

	if con_guess > ran_num && loop_iterator != 4{
		return "\nYou need to go lower\n",false
	}

	if con_guess < ran_num && loop_iterator != 4{
		return "\nYou need to go higher\n",false
	}

	if con_guess == ran_num {
		return fmt.Sprintf("\nYou Won!\nGuesses: %d",loop_iterator+1),true
	}
	if loop_iterator == 4 {
		return fmt.Sprintf("\nYou Loss!\nThe number was %d",ran_num),true
	}
	return "",false
}


func main() {
	x := 0
	var small int = 0
	var great int = 100
	

	for {
		ran_num := rand.Intn(great)
		
		x+=1
		for i :=0; i < 5; i++{
			var guess string
			fmt.Printf("you have %v guess(es) left\n", 5-i)
			fmt.Printf("Guess the number from %[1]d to %[2]d: ",small,great)
			fmt.Scanln(&guess)
			fmt.Println(guess)

			return_message, game_state := gameLogic(guess,ran_num,small,great,i,5)

			if game_state{
				fmt.Println(return_message)
				break
			} else {
				fmt.Println(return_message)	
			}
		}
		var again string 
		fmt.Print("Do you want to play again?: ")
		fmt.Scanln(&again)
		if slices.Contains([]string{"yes","y","si","yep"},again) {
			fmt.Print("\033c")
			continue
			
		}else if slices.Contains([]string{"no","n","nope"},again) {
			fmt.Println("Goodbye")
			break
			
		}else {
			fmt.Println("Invalid Input. Quiting Program")
			break
		}

	




	}
	
	
	
	
}