package main

//go mod init myapp   (type this for every project)
//https://www.youtube.com/watch?v=yyUHQIec83I&t=2s
//to run a go program
//go run main.go

//you uee the myapp/  which is the module you created with mod init command, so you can create and import your own modules
import (
	"fmt"
	"myapp/helper"
	"strconv"
	"sync"
	"time"
)

var conferenceName = "Rick James Comference"

const conferenceTickets int16 = 50

var remainingTickets int = 50

type UserData struct {
	userName    string
	userTickets int
}

func greetUsers(conferenceName string, conferenceTickets int16, remainingTickets int) {
	fmt.Println("Welcome to", conferenceName, "booking application")
	//below is when you want to use the printf function for formatting output
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

//https://www.youtube.com/watch?v=yyUHQIec83I&t=2s
//at 3:10 minutes explains breaking this out in it's own thread
func sendEmailConfForTicket(userTickets int, userName string) {

	time.Sleep(50 * time.Second) //stops the execution of the thread for 10 seconds
	//Sprintf allows you to save outpu to a variable
	var ticketInfo = fmt.Sprintf("%v tickets for %v", userTickets, userName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n%v \nto %v\n", ticketInfo, userName)
	fmt.Println("###################")
	//wg.Done()    //says that I am done, resume execution on the main thread
}

// a wait group allows you to add threads to a waiting container
var wg = sync.WaitGroup{}

func main() {

	//var attendees [50]string  --> this is a fixed size array
	var attendees []string                         //this is the syntax for a "slice" (slice is a dynamic data structure with no fixed size)
	var m_attendees = make([]map[string]string, 0) //create a map - initialize the size to 0
	var o_m_attendees = make([]UserData, 0)        //create structure of type UserData

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		var userName string
		var userTickets int

		fmt.Println("Who is buying a ticket ?..")
		fmt.Scan(&userName)

		//with fmt.Scan (you cannot declare your variables using userTickets:= 0, it will not pause for input)
		fmt.Println("How many tickets do yo want to purchase ?..")
		fmt.Scan(&userTickets)

		//DEMONSTRATES HOW A FUNCTION CAN RETURN MORE THAN ONE VALUE
		isValidName, isValidTicketNumber := helper.ValidateUserInput(userName, userTickets, remainingTickets)

		if isValidName && isValidTicketNumber {

			//https://www.youtube.com/watch?v=yyUHQIec83I&t=2s
			//at 3:10 minutes explains breaking this out in it's own thread
			//makes this a "go routine" - runs the below function in it's own seperate thread
			//wg.Add(1)   -->this means add a thread to the wait group
			go sendEmailConfForTicket(userTickets, userName)

			//.append can only be used with a slice
			attendees = append(attendees, userName)

			//below is how to use a map
			var userData = make(map[string]string)
			userData["userName"] = userName
			userData["userTickets"] = strconv.FormatInt(int64(userTickets), 10) //convert int to string
			m_attendees = append(m_attendees, userData)

			//below is how to use a structure
			var userData_2 = UserData{
				userName:    userName,
				userTickets: userTickets,
			}
			o_m_attendees = append(o_m_attendees, userData_2)

			remainingTickets = remainingTickets - userTickets

			/* NO NEED FOR THIS ANYMORE
			if userTickets > remainingTickets {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
				continue //allows the program not to stop, but continue to the next iteration
			}
			*/

			fmt.Printf("Thank you %v for booking %v tickets. Your purchase is appreciated \n", userName, userTickets)
			fmt.Printf("There are %d tickets remaining for the %v \n", remainingTickets, conferenceName)

			peopleComing := []string{}
			for _, attendee := range attendees {
				var names = attendee
				peopleComing = append(peopleComing, names)
			}

			fmt.Printf("These are your bookings: %v\n", peopleComing)

			if remainingTickets <= 0 {
				fmt.Printf("%v is Sold Out!!!!\n", conferenceName)
				break
			}

		} else {
			fmt.Println("Either your name is invalid or the number of tickets is invalid")
			continue //allows the program not to stop, but continue to the next iteration
		}

	} //for loop

	fmt.Printf("List of bookings from map slice %v\n", m_attendees)
	fmt.Printf("List of bookings from map a struct %v\n", o_m_attendees)
	//wg.Wait()   //tells the app to wait until all threads are finished

}
