package main 
import "fmt"
import "time"
import "github.com/gen2brain/beeep"
import "github.com/charmbracelet/lipgloss"


func playBeep(){
	// play beep
		err := beeep.Beep(440, 1000)
    	if err != nil {
    		panic(err)
		}
}

func showNotification(title string, message string){

	//show notificaation
	noti := beeep.Notify(title, message, "")
	if noti != nil {
		panic(noti)
	}

}

func pomodoro(pomoTime int){
	for i:=pomoTime;i>=0;i-- {
		time.Sleep(1*time.Second)


	}

	fmt.Println("time is over..")
	
}

func main(){

	//styling with lipgloss lib
	var menuStyle=lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#7D56F4")).
	PaddingTop(0).
    PaddingLeft(0).
    Width(22)



	var inp string
	var workTime int = 25
	var breakTime int = 5
	var sessionSize int = 1
	for inp!="e" {
	fmt.Println(menuStyle.Render("\n This is a pomodoro tracker \n  Enter 0 to start pomodoro timer\n Enter c to change work time and break time \n Enter e to exit \n "))
	fmt.Scan(&inp)
		
	if inp=="0"{

		fmt.Println("How many sessions do you want to do? 1 session means 25 minute work 5 minute break \n")
		fmt.Println("Enter how many sessions you want to do: ")
		fmt.Scan(&sessionSize)
		//sessionSize input validation
		if sessionSize<0 {
			fmt.Println("minimum session number is 1..\n")
			fmt.Println("Enter how many sessions you want to do: ")
			fmt.Scan(&sessionSize)
		}else {
			sessionSize=sessionSize-1
			for j:=sessionSize;j>=0;j-- {
				fmt.Println("\n Pomodoro timer started.. get to work \n")
				pomodoro(5)
				playBeep()
				showNotification("Time Over","time to take a break, timer started")
				// break timer
				pomodoro(10)
				playBeep()
				showNotification("Break over","time to get back to work")
			}

			fmt.Println("\n \n Sessions ended.. Well done.. Treat yourself")
		}
		
		

	
	

	}else if inp=="c"{
		fmt.Println("Enter new work time: ")
		fmt.Scan(&workTime)
		fmt.Println("Enter new Break time: ")
		fmt.Scan(&breakTime)
		fmt.Println("\n\n New work time and break time is set..\n\n")
		fmt.Println("\n This is a pomodoro tracker \n  Enter 0 to start pomodoro timer\n Enter c to change work time and break time \n Enter e to exit \n ")
		fmt.Scan(&inp)
	

	}

}


}