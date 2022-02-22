//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func weekday(day int) bool {
	return day <= 4
}

func main() {
	// The day and role. Change these to check your work.
	today, role := Sunday, Contractor

	//* Use the accessGranted() and accessDenied() functions to display
	//  informational messages
	if role == Admin || role == Manager {
		//* Access at any time: Admin, Manager
		accessGranted()
	} else if role == Contractor && !weekday(today) {
		//* Access weekends: Contractor
		accessGranted()
	} else if role == Member && weekday(today) {
		//* Access weekdays: Member
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		//* Access Mondays, Wednesdays, and Fridays: Guest
		accessGranted()
	} else {
		accessDenied()
	}

	// Second solution**

	// func accessGranted(day int, user int) {
	// 	if day < 7 && (user == 10 || user == 20) {
	// 		fmt.Println("Granted")
	// 	} else if day >= 5 && day <= 7 && user == 30 {
	// 		fmt.Println("Granted")
	// 	} else if day <= 4 && user == 40 {
	// 		fmt.Println("Granted")
	// 	} else if (day == 0 || day == 2 || day == 4) && user == 50 {
	// 		fmt.Println("Granted")
	// 	} else {
	// 		accessDenied()
	// 	}
	// }

	// func accessDenied() {
	// 	fmt.Println("Denied")
	// }

	// func main() {
	//// The day and role. Change these to check your work.
	// 	accessGranted(Saturday, Contractor)

	// First Solution**

	// //* Access at any time: Admin, Manager
	// if today, role := Saturday, Admin; today < 6 && role == 10 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }

	// if today, role := Monday, Manager; today < 6 && role == 20 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }
	// //------------------------------------------------------------------

	// //* Access weekends: Contractor
	// if today, role := Saturday, Contractor; today == 5 && role == 30 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }

	// if today, role := Sunday, Contractor; today == 6 && role == 30 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }
	// //-------------------------------------------------------------------

	// //* Access weekdays: Member
	// if today, role := Monday, Member; today == 0 && role == 40 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }

	// if today, role := Tuesday, Member; today == 1 && role == 40 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }

	// if today, role := Wednesday, Member; today == 2 && role == 40 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }

	// if today, role := Thursday, Member; today == 3 && role == 40 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }

	// if today, role := Friday, Member; today == 4 && role == 40 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }
	// //------------------------------------------------------------------------

	// //* Access Mondays, Wednesdays, and Fridays: Guest

	// if today, role := Monday, Guest; today == 0 && role == 50 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }

	// if today, role := Wednesday, Guest; today == 2 && role == 50 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }

	// if today, role := Friday, Guest; today == 4 && role == 50 {
	// 	accessGranted()
	// } else {
	// 	accessDenied()
	// }
}
