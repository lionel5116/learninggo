package helper

//DEMONSTRATES HOW A FUNCTION CAN RETURN MORE THAN ONE VALUE
//a function is exported by making the function name befin with a capital letter (no need to use an exports keyword)
func ValidateUserInput(userName string, userTickets int, remainingTickets int) (bool, bool) {
	isValidName := len(userName) >= 2
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidTicketNumber
}
