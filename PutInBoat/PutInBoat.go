package PutInBoat

var state = "boat is empty"

func ViewState() string {

	return state
}

func InsertHS() {
	state = "HS is in boat"
}

func InsertFox() {
	state = "Fox is in boat"
}

func InsertChicken() {
	state = "Chicken is in boat"
}

func InsertCorn() {
	state = "Corn is in boat"
}
