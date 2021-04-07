package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var TEACHERS []*Teacher
var STUDENTS []*Student
var SUBJECTS []*Subject
var CLASSES []*Class

// Teacher Definement ( Name | ContactInfo | Wage | TeachingSubjects | Authority Level | TimeStamp )
type Teacher struct {
	FirstName       string
	LastName        string
	Email           string
	Number          string
	Wage            float32
	Subjects        *[]Subject
	isPrincipal     bool
	isVicePrincipal bool
	hireDate        string
}

// Student Definement ( Name | LearningSubjects | Grades | DateOfBirth | TimeStamp )
type Student struct {
	FirstName  string
	LastName   string
	Subjects   *[]Subject
	Grades     map[string]int
	DOB        string
	enrollDate string
}

// Subject Definement ( Name | Weekday : List of Times )
type Subject struct {
	Name      string
	TimeTable map[string][]string
}

type Class struct {
	ClassCode    string
	Teacher      *Teacher
	Students     []*Student
	creationDate string
}

// Get Currant Date and Time Function ( Returns TimeStamp String )
func getDateTime() string {
	today := time.Now()
	year := today.Year()
	month := today.Month()
	day := today.Day()
	hour := today.Hour()
	minute := today.Minute()
	second := today.Second()
	currentDateTime := fmt.Sprintf("%v:%v:%v - %v/%v/%v", hour, minute, second, day, month, year)
	return currentDateTime
}

// Main Menu
func mainMenu() {
	for {
		var menu string = "Menu"
		var createM string = "Go To Create Menu"
		var viewM string = "Go To View Menu"
		var editM string = "Go To Edit Menu"
		var assignM string = "Go To Assignment Menu"

		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("\n\n %#20q \n\n", menu)
		fmt.Printf("/createMenu %#30q \n", createM)
		fmt.Printf("/viewMenu %#30q \n", viewM)
		fmt.Printf("/editMenu %#30q \n", editM)
		fmt.Printf("/assignMenu %#34q \n", assignM)

		fmt.Println("\nChoice: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		menuOption := scanner.Text()
		if menuOption == "/createMenu" {
			createMenu()
		} else if menuOption == "/viewMenu" {
			viewMenu()
		} else if menuOption == "/editMenu" {
			editMenu()
		} else if menuOption == "/assignMenu" {
			assignMenu()
		}
	}
}

// Create Menu
func createMenu() {
	for {
		var menu string = "Create Menu"
		var createU string = "Creates a User"
		var createC string = "Creates a Class"
		var createS string = "Creates a Subject"
		var mm string = "Go back to Main Menu"

		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("\n\n %#25q \n\n", menu)
		fmt.Printf("/createUser %#30q \n", createU)
		fmt.Printf("/createClass %#30q \n", createC)
		fmt.Printf("/createSubject %#30q \n\n", createS)
		fmt.Printf("/mainMenu %#38q \n", mm)

		fmt.Println("\nChoice: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		menuOption := scanner.Text()
		if menuOption == "/createUser" {
			createUser()
		} else if menuOption == "/createClass" {
			createClass()
		} else if menuOption == "/createSubject" {
			createSubject()
		} else if menuOption == "/editUser" {
			fmt.Println("Edit User?")
		} else if menuOption == "/mainMenu" {
			mainMenu()
		}
	}
}

// View Menu
func viewMenu() {
	for {
		var menu string = "View Menu"
		var viewU string = "Views Users"
		var viewC string = "Views Classes"
		var viewS string = "Views a Subjects"
		var mm string = "Go back to Main Menu"

		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("\n\n %#25q \n\n", menu)
		fmt.Printf("/viewUser %#30q \n", viewU)
		fmt.Printf("/viewClass %#31q \n", viewC)
		fmt.Printf("/viewSubject %#32q \n\n", viewS)
		fmt.Printf("/mainMenu %#39q \n", mm)

		fmt.Println("\nChoice: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		menuOption := scanner.Text()
		if menuOption == "/viewUser" {
			viewUser()
		} else if menuOption == "/viewClass" {
			viewClass()
		} else if menuOption == "/viewSubject" {
			viewSubject()
		} else if menuOption == "/mainMenu" {
			mainMenu()
		}
	}
}

// Edit Menu
func editMenu() {
	for {
		var menu string = "Edit Menu"
		var editU string = "Edit Users"
		var editC string = "Edit Classes"
		var editS string = "Edit Subjects"
		var mm string = "Go back to Main Menu"

		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("\n\n %#25q \n\n", menu)
		fmt.Printf("/editUser %#30q \n", editU)
		fmt.Printf("/editClass %#31q \n", editC)
		fmt.Printf("/editSubject %#30q \n\n", editS)
		fmt.Printf("/mainMenu %#40q \n", mm)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		menuOption := scanner.Text()
		if menuOption == "/editUser" {
			editUser()
		} else if menuOption == "/editClass" {
			editClass()
		} else if menuOption == "/editSubject" {
			editSubject()
		} else if menuOption == "/mainMenu" {
			mainMenu()
		}

	}
}

// Assign Menu
func assignMenu() {
	for {
		var menu string = "Assignment Menu"
		var assignU string = "Assign Users"
		var assignS string = "Assign Subjects"
		var mm string = "Go back to Main Menu"

		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("\n\n %#25q \n\n", menu)
		fmt.Printf("/assignUser %#30q \n", assignU)
		fmt.Printf("/assignSubject %#30q \n\n", assignS)
		fmt.Printf("/mainMenu %#40q \n", mm)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		menuOption := scanner.Text()
		if menuOption == "/assignUser" {
			assignUser()
		} else if menuOption == "/assignSubject" {
			assignSubject()
		} else if menuOption == "/mainMenu" {
			mainMenu()
		}
	}
}

// Create User Menu
func createUser() {
	for {
		var userMenu string = "Create User Menu"
		var createT string = "Create a Teacher"
		var createS string = "Create a Student"
		var mm string = "Go back to Main Menu"

		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("\n\n %#25q \n\n", userMenu)
		fmt.Printf("/createTeacher %#30q \n", createT)
		fmt.Printf("/createStudent %#30q \n\n", createS)
		fmt.Printf("/mainMenu %#39q \n", mm)
		fmt.Printf("\n")
		fmt.Printf("Choice: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userOption := scanner.Text()
		if userOption == "/createTeacher" {
			createTeacher()
		} else if userOption == "/createStudent" {
			createStudent()
		} else if userOption == "/mainMenu" {
			mainMenu()
		}
	}
}

func createClass() {
	var code string
	var teacher *Teacher = &Teacher{}
	var students []*Student = []*Student{}
	var creation string = getDateTime()

	for {
		fmt.Printf("\n")
		fmt.Printf("Enter Code For The Class: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		code = scanner.Text()
		fmt.Printf("\n")
		fmt.Printf("Your Class Code Is: %v \n", code)
		break
	}
	classroom := &Class{code, teacher, students, creation}
	CLASSES = append(CLASSES, classroom)
	fmt.Println(*classroom)
}

// Create Teacher
func createTeacher() {
	var fname string
	var sname string
	var email string
	var number string
	var wage float32
	var mrPrincipal bool
	var mrVice bool
	var timeStamp string

	for {
		// Get FirstName
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("Please Enter Teachers First Name: ")
		scanner.Scan()
		fname = scanner.Text()
		break
	}
	for {
		// Get SecondName
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Please Enter Teachers Second Name: ")
		scanner.Scan()
		sname = scanner.Text()
		break
	}
	for {
		// Get Email
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Please Enter Teachers Email: ")
		scanner.Scan()
		email = scanner.Text()
		break
	}
	for {
		// Get Number
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Please Enter Teachers Phone Number: ")
		scanner.Scan()
		number = scanner.Text()
		break
	}
	for {
		// Get Wage
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Please Enter Teachers Wage: ")
		scanner.Scan()
		twage, err := strconv.ParseFloat(scanner.Text(), 10)
		if err != nil {
			fmt.Printf("Oops: %v", err)
		}
		wage = float32(twage)
		break
	}
	for {
		// Get Primary Authority Level
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Is This Teacher A Principal?: ")
		scanner.Scan()
		principal := scanner.Text()

		if strings.ToLower(principal) == "yes" {
			mrPrincipal = true
			break
		} else if strings.ToLower(principal) == "no" {
			mrPrincipal = false
			break
		}
	}
	for {
		// Get Secondary Level
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Is This Teacher A Vice Principal?: ")
		scanner.Scan()
		vice := scanner.Text()

		if strings.ToLower(vice) == "yes" {
			mrVice = true
			break
		} else if strings.ToLower(vice) == "no" {
			mrVice = false
			break
		}
	}
	for {
		// Get TimeStamp
		timeStamp = getDateTime()
		break
	}
	newTeacher := &Teacher{
		fname,
		sname,
		email,
		number,
		wage,
		&[]Subject{},
		mrPrincipal,
		mrVice,
		timeStamp,
	}
	TEACHERS = append(TEACHERS, newTeacher)
	fmt.Println(*newTeacher)

	mainMenu()
}

func createStudent() {
	var fname string
	var sname string
	var grades map[string]int = map[string]int{}
	var dob string
	var timeStamp string

	for {
		// Get FirstName
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("Please Enter Students First Name: ")
		scanner.Scan()
		fname = scanner.Text()
		break
	}
	for {
		// Get SecondName
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Please Enter Students Second Name: ")
		scanner.Scan()
		sname = scanner.Text()
		break
	}
	for {
		// Get DOB
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Please Enter Students Date of Birth: ")
		scanner.Scan()
		dob = scanner.Text()
		break
	}
	for {
		// Get TimeStamp
		timeStamp = getDateTime()
		break
	}
	newStudent := &Student{
		fname,
		sname,
		&[]Subject{},
		grades,
		dob,
		timeStamp,
	}
	STUDENTS = append(STUDENTS, newStudent)
	fmt.Println(*newStudent)

	mainMenu()
}

func createSubject() {
	var name string
	var times []string
	var timeTable map[string][]string

	// Subject Name
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("Name of this Subject: ")
		scanner.Scan()
		name = scanner.Text()
		break
	}
	// Subject TimeTable ( Monday )

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Monday's TimeTable")
		fmt.Printf("Enter Time (00:00) or x to finish: ")
		scanner.Scan()
		time := scanner.Text()
		if time == "x" {
			timeTable = map[string][]string{"Monday": times}
			fmt.Println(timeTable["Monday"])
			break
		}
		times = append(times, time)
		fmt.Println(times)
	}

	// Subject TimeTable ( Tuesday )
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Tuesday's TimeTable")
		fmt.Printf("Enter Time (00:00) or x to finish: ")
		scanner.Scan()
		time := scanner.Text()
		if time == "x" {
			timeTable["Tuesday"] = times
			break
		}
		times = append(times, time)
	}
	// Subject TimeTable ( Wednesday )
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Wednesday's TimeTable")
		fmt.Printf("Enter Time (00:00) or x to finish: ")
		scanner.Scan()
		time := scanner.Text()
		if time == "x" {
			timeTable["Wednesday"] = times
			break
		}
		times = append(times, time)
	}
	// Subject TimeTable ( Thursday )
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Thursday's TimeTable")
		fmt.Printf("Enter Time (00:00) or x to finish: ")
		scanner.Scan()
		time := scanner.Text()
		if time == "x" {
			timeTable["Thursday"] = times
			break
		}
		times = append(times, time)
	}
	// Subject TimeTable ( Friday )
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Friday's TimeTable")
		fmt.Printf("Enter Time (00:00) or x to finish: ")
		scanner.Scan()
		time := scanner.Text()
		if time == "x" {
			timeTable["Friday"] = times
			break
		}
		times = append(times, time)
	}
	newSubject := &Subject{name, timeTable}
	SUBJECTS = append(SUBJECTS, newSubject)
	fmt.Println(*newSubject)

	mainMenu()
}

// View User Menu
func viewUser() {
	for {
		var userMenu string = "View User Menu"
		var viewT string = "View a Teacher"
		var viewS string = "View a Student"
		var mm string = "Go back to Main Menu"

		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("\n\n %#25q \n\n", userMenu)
		fmt.Printf("/viewTeacher %#30q \n", viewT)
		fmt.Printf("/viewStudent %#30q \n\n", viewS)
		fmt.Printf("/mainMenu %#39q \n", mm)
		fmt.Println("\nChoice: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userOption := scanner.Text()
		if userOption == "/viewTeacher" {
			viewTeacher()
		} else if userOption == "/viewStudent" {
			viewStudent()
		} else if userOption == "/mainMenu" {
			mainMenu()
		}
	}
}

// View Teachers
func viewTeacher() {
	for _, v := range TEACHERS {
		name := v.FirstName + " " + v.LastName
		fmt.Println("\n\n\n\n\n ")
		fmt.Println("* Teacher *")
		fmt.Printf("Name                 : %v \n", name)
		fmt.Printf("Email                : %v \n", v.Email)
		fmt.Printf("Number               : %v \n", v.Number)
		fmt.Printf("Wage                 : %v \n", v.Wage)
		fmt.Printf("Subjects             : %v \n", v.Subjects)
		if v.isPrincipal {
			fmt.Println("This Person is the Principal! ")
		} else if v.isVicePrincipal {
			fmt.Println("This Person is the Vice Principal!")
		}
		fmt.Printf("Started Working here : %v \n", v.hireDate)
	}
}

// View Students
func viewStudent() {
	for _, v := range STUDENTS {
		name := v.FirstName + " " + v.LastName
		fmt.Println("\n\n\n\n\n ")
		fmt.Println("* Student *")
		fmt.Printf("Name                  : %v \n", name)
		fmt.Printf("Subjects              : %v \n", v.Subjects)
		fmt.Printf("Grades                : %v \n", v.Grades)
		fmt.Printf("Date of Birth         : %v \n", v.DOB)
		fmt.Printf("Subjects              : %v \n", v.Subjects)
		fmt.Printf("Started Studying here : %v \n", v.enrollDate)
	}
}

// View Classes
func viewClass() {
	for _, c := range CLASSES {
		fmt.Println("\n\n\n\n\n ")
		fmt.Println("* Class *")
		fmt.Printf("Code          : %v \n", c.ClassCode)
		fmt.Printf("Teacher       : %v \n", c.Teacher.FirstName)
		fmt.Println("Students      :")
		for _, s := range c.Students {
			fmt.Println(s.FirstName)
		}
		fmt.Printf("Class Created : %v \n", c.creationDate)
	}
}

// View Subjects
func viewSubject() {
	for _, v := range SUBJECTS {
		fmt.Println("\n\n\n\n\n ")
		fmt.Println("* Subject *")
		fmt.Printf("Subject Name : %v \n", v.Name)
		fmt.Printf("Time Table   : %v \n", v.TimeTable)
	}
}

// Edit User
func editUser() {
	for {
		var userMenu string = "Edit User Menu"
		var editT string = "Edit a Teacher"
		var editS string = "Edit a Student"
		var mm string = "Go back to Main Menu"

		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("\n\n %#25q \n\n", userMenu)
		fmt.Printf("/editTeacher %#30q \n", editT)
		fmt.Printf("/editStudent %#30q \n\n", editS)
		fmt.Printf("/mainMenu %#39q \n", mm)
		fmt.Println("\nChoice: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userOption := scanner.Text()
		if userOption == "/editTeacher" {
			editTeacher()
		} else if userOption == "/editStudent" {
			editStudent()
		} else if userOption == "/mainMenu" {
			mainMenu()
		}
	}
}

// Edit Teacher
func editTeacher() {
	var tChoice string
	var teacher *Teacher

	viewTeacher()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\n")
	fmt.Printf("Please Enter The Name Of The Teacher You Would Like To Edit: ")
	scanner.Scan()
	tChoice = scanner.Text()
	for _, v := range TEACHERS {
		if strings.ToLower(tChoice) == strings.ToLower(v.FirstName) {
			teacher = &*v
			name := teacher.FirstName + " " + teacher.LastName
			fmt.Println("\n\n\n\n\n ")
			fmt.Println("* Edit Teacher *")
			fmt.Printf("Name                 : %v \n", name)
			fmt.Printf("Email                : %v \n", teacher.Email)
			fmt.Printf("Number               : %v \n", teacher.Number)
			fmt.Printf("Wage                 : %v \n", teacher.Wage)
			fmt.Printf("Subjects             : %v \n", teacher.Subjects)
			if teacher.isPrincipal {
				fmt.Println("This Person is the Principal! ")
			} else if teacher.isVicePrincipal {
				fmt.Println("This Person is the Vice Principal!")
			}
			fmt.Printf("Started Working here: %v \n", teacher.hireDate)
		}
	}
	for {
		var fname string = "Change First Name"
		var lname string = "Change Surname"
		var email string = "Change Email"
		var number string = "Change Number"
		var wage string = "Change Wage"
		var alevel string = "Change Authority Level"
		var mm string = "Go To Main Menu"

		fmt.Println("\nPick Which Information You Would Like To Edit | Options Are:")
		fmt.Printf("/firstName %#30q \n", fname)
		fmt.Printf("/lastName %#28q \n", lname)
		fmt.Printf("/email %#29q \n", email)
		fmt.Printf("/number %#29q \n", number)
		fmt.Printf("/wage %#29q \n", wage)
		fmt.Printf("/authorityLevel %#30q \n\n", alevel)
		fmt.Printf("/mainMenu %#29q \n", mm)

		fmt.Printf("\nChoice: ")

		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		editOption := scanner.Text()
		if editOption == "/firstName" {
			fmt.Printf("\n")
			fmt.Printf("Please Enter New First Name: ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newFname := scanner.Text()
			oldFname := teacher.FirstName
			teacher.FirstName = newFname
			fmt.Printf("Teachers Name Changed From %v to %v", oldFname, newFname)
			break
		} else if editOption == "/lastName" {
			fmt.Printf("\n")
			fmt.Printf("Please Enter New Last Name: ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newLname := scanner.Text()
			oldLname := teacher.LastName
			teacher.LastName = newLname
			fmt.Printf("Teachers Name Changed From %v to %v", oldLname, newLname)
			break
		} else if editOption == "/email" {
			fmt.Printf("\n")
			fmt.Printf("Please Enter New Email: ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newEmail := scanner.Text()
			oldEmail := teacher.Email
			teacher.Email = newEmail
			fmt.Printf("Teachers Email Changed From %v to %v", oldEmail, newEmail)
			break
		} else if editOption == "/number" {
			fmt.Printf("\n")
			fmt.Printf("Please Enter New Number: ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newNumber := scanner.Text()
			oldNumber := teacher.Number
			teacher.Number = newNumber
			fmt.Printf("Teachers Email Changed From %v to %v", oldNumber, newNumber)
			break
		} else if editOption == "/wage" {
			fmt.Printf("\n")
			fmt.Printf("Please Enter New Wage: ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newWage, err := strconv.ParseFloat(scanner.Text(), 10)
			if err != nil {
				fmt.Printf("Oops: %v", err)
			}
			oldWage := teacher.Wage
			teacher.Wage = float32(newWage)
			fmt.Printf("\nTeachers Wage Changed From %v to %v", oldWage, newWage)
			break
		} else if editOption == "/authorityLevel" {
			var v string = "Change To Vice Principal"
			var p string = "Change To Principal"
			var t string = "Change To Teacher"
			var nevermind string = "Never Mind"

			if teacher.isPrincipal {
				for {
					fmt.Println("Teacher Is Principal, Would You Like To DownGrade to Vice Principal or Teacher?")
					fmt.Printf("/vicePrincipal %#30q \n", v)
					fmt.Printf("/teacher %#30q \n\n", t)
					fmt.Printf("/nevermind  %#30q \n", nevermind)

					scanner = bufio.NewScanner(os.Stdin)
					scanner.Scan()
					option := scanner.Text()
					if option == "/vicePrincipal" {
						teacher.isPrincipal = false
						teacher.isVicePrincipal = true
						break
					} else if option == "/teacher" {
						teacher.isPrincipal = false
						teacher.isVicePrincipal = false
						break
					} else if option == "/nevermind" {
						mainMenu()
					}
				}
			} else if teacher.isVicePrincipal {
				for {
					fmt.Println("Teacher Is Vice Principal, Would You Like To DownGrade Teacher or UpGrade To Principal?")
					fmt.Printf("/principal %#31q \n", p)
					fmt.Printf("/teacher %#31q \n\n", t)
					fmt.Printf("/nevermind  %#25q \n", nevermind)

					scanner = bufio.NewScanner(os.Stdin)
					scanner.Scan()
					option := scanner.Text()
					if option == "/principal" {
						teacher.isPrincipal = true
						teacher.isVicePrincipal = false
						break
					} else if option == "/teacher" {
						teacher.isPrincipal = false
						teacher.isVicePrincipal = false
						break
					} else if option == "/nevermind" {
						mainMenu()
					}
				}
			} else if !teacher.isPrincipal && !teacher.isVicePrincipal {
				for {
					fmt.Println("Would You Like To UpGrade Teacher To Principal or Vice Principal?")
					fmt.Printf("/principal %#31q \n", p)
					fmt.Printf("/vicePrincipal %#32q \n\n", v)
					fmt.Printf("/nevermind  %#22q \n", nevermind)

					scanner = bufio.NewScanner(os.Stdin)
					scanner.Scan()
					option := scanner.Text()
					if option == "/principal" {
						teacher.isPrincipal = true
						teacher.isVicePrincipal = false
						break
					} else if option == "/vicePrincipal" {
						teacher.isPrincipal = false
						teacher.isVicePrincipal = true
						break
					} else if option == "/nevermind" {
						mainMenu()
					}
				}
			}
		} else if editOption == "/mainMenu" {
			mainMenu()
		}
	}
}

// Edit Student
func editStudent() {
	var sChoice string
	var student *Student

	viewStudent()
	fmt.Printf("\n")
	fmt.Printf("Please Enter The Name Of The Student You Would Like To Edit: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sChoice = scanner.Text()
	for _, v := range STUDENTS {
		if strings.ToLower(sChoice) == strings.ToLower(v.FirstName) {
			student = &*v
			name := student.FirstName + " " + student.LastName
			fmt.Println("\n\n\n\n\n ")
			fmt.Println("* Edit Student *")
			fmt.Printf("Name                 : %v \n", name)
			fmt.Printf("Date Of Birth        : %v \n", student.DOB)
			fmt.Printf("Grades               : %v \n", student.Grades)
			fmt.Printf("Subjects             : %v \n", student.Subjects)
			fmt.Printf("Enrolled here        : %v \n", student.enrollDate)
		}
	}
	for {
		var fname string = "Change First Name"
		var lname string = "Change Surname"
		var dob string = "Change DOB"
		var grades string = "Change Grades"
		var mm string = "Go To Main Menu"

		fmt.Println("\nPick Which Information You Would Like To Edit | Options Are:")
		fmt.Printf("/firstName %#30q \n", fname)
		fmt.Printf("/lastName %#28q \n", lname)
		fmt.Printf("/dob %#29q \n", dob)
		fmt.Printf("/grades %#29q \n", grades)
		fmt.Printf("/mainMenu %#29q \n", mm)

		fmt.Println("\nChoice: ")

		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		editOption := scanner.Text()
		if editOption == "/firstName" {
			fmt.Printf("\n")
			fmt.Printf("Please Enter New First Name: ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newFname := scanner.Text()
			oldFname := student.FirstName
			student.FirstName = newFname
			fmt.Printf("Teachers Name Changed From %v to %v", oldFname, newFname)
			break
		} else if editOption == "/lastName" {
			fmt.Printf("\n")
			fmt.Printf("Please Enter New Last Name: ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newLname := scanner.Text()
			oldLname := student.LastName
			student.LastName = newLname
			fmt.Printf("Teachers Name Changed From %v to %v", oldLname, newLname)
			break
		} else if editOption == "/dob" {
			fmt.Printf("\n")
			fmt.Printf("Please Enter New Date Of Birth: ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newDOB := scanner.Text()
			oldDOB := student.DOB
			student.DOB = newDOB
			fmt.Printf("Teachers Name Changed From %v to %v", oldDOB, newDOB)
			break
		} else if editOption == "/grades" {
			fmt.Printf("\n")
			fmt.Printf("Enter Subject Name to Add Grades: ")
			scanner = bufio.NewScanner(os.Stdin)
			scanner.Scan()
			subName := scanner.Text()
			for k, _ := range student.Grades {
				if strings.ToLower(subName) == strings.ToLower(k) {
					fmt.Printf("Enter The Grade For Subject %v: ", k)
					scanner.Scan()
					grade, err := strconv.ParseInt(scanner.Text(), 10, 64)
					if err != nil {
						fmt.Printf("Oops: %v", err)
					}
					student.Grades[k] = int(grade)
				}

			}
			fmt.Println(student.Grades)
		} else if editOption == "/mainMenu" {
			mainMenu()
		}
	}
}

func editClass() {
	var classroom *Class
	var classCode string

	for {
		viewClass()
		fmt.Printf("\n")
		fmt.Printf("Enter Class Code of the Classroom You Want To Edit: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		classCode = scanner.Text()
		for _, c := range CLASSES {
			if classCode == c.ClassCode {
				classroom = &*c

			}
		}
		if classroom != nil {
			break
		}
	}
	fmt.Printf("\n")
	fmt.Printf("Enter New Code For Class: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	newCode := scanner.Text()
	oldCode := classroom.ClassCode
	classroom.ClassCode = newCode
	fmt.Printf("Classroom Code Changed From %v to %v \n", oldCode, newCode)
}

func editSubject() {
	var subject *Subject
	var sChoice string
	var name string = "Change Name"
	var tt string = "Change Time Table"
	var times []string
	var timeTable map[string][]string

	viewSubject()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\n")
	fmt.Printf("Please Enter The Name Of The Subject You Would Like To Edit: ")
	scanner.Scan()
	sChoice = scanner.Text()
	for _, v := range SUBJECTS {
		if strings.ToLower(v.Name) == strings.ToLower(sChoice) {
			subject = *&v
		}
	}
	for {
		fmt.Println("\n\n\n\n\n ")
		fmt.Println("* Edit Subject *")
		fmt.Printf("Name                 : %v \n", subject.Name)
		fmt.Printf("Time Table           : %v \n", subject.TimeTable)

		fmt.Printf("Would You Like To Edit Subject Name ( %v ) or Subject TimeTable ( %v ): \n\n", subject.Name, subject.TimeTable)
		fmt.Printf("/name %#30q \n", name)
		fmt.Printf("/timeTable %#31q \n", tt)

		scanner.Scan()
		sChoice = scanner.Text()
		if sChoice == "/name" {
			fmt.Printf("\n")
			fmt.Printf("Enter New Subject Name: ")
			scanner.Scan()
			sName := scanner.Text()
			oldsName := subject.Name
			subject.Name = sName
			fmt.Printf("\nSubject Name Changed From %v to %v!", oldsName, subject.Name)
			break
		} else if sChoice == "/timeTable" {
			for {
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Printf("\n")
				fmt.Println("Monday's TimeTable")
				fmt.Printf("Enter Time (00:00) or x to finish: ")
				scanner.Scan()
				time := scanner.Text()
				if time == "x" {
					timeTable = map[string][]string{"Monday": times}
					fmt.Println(timeTable["Monday"])
					break
				}
				times = append(times, time)
				fmt.Println(times)
			}

			// Subject TimeTable ( Tuesday )
			for {
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Printf("\n")
				fmt.Println("Tuesday's TimeTable")
				fmt.Printf("Enter Time (00:00) or x to finish: ")
				scanner.Scan()
				time := scanner.Text()
				if time == "x" {
					timeTable["Tuesday"] = times
					break
				}
				times = append(times, time)
			}
			// Subject TimeTable ( Wednesday )
			for {
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Printf("\n")
				fmt.Println("Wednesday's TimeTable")
				fmt.Printf("Enter Time (00:00) or x to finish: ")
				scanner.Scan()
				time := scanner.Text()
				if time == "x" {
					timeTable["Wednesday"] = times
					break
				}
				times = append(times, time)
			}
			// Subject TimeTable ( Thursday )
			for {
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Printf("\n")
				fmt.Println("Thursday's TimeTable")
				fmt.Printf("Enter Time (00:00) or x to finish: ")
				scanner.Scan()
				time := scanner.Text()
				if time == "x" {
					timeTable["Thursday"] = times
					break
				}
				times = append(times, time)
			}
			// Subject TimeTable ( Friday )
			for {
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Printf("\n")
				fmt.Println("Friday's TimeTable")
				fmt.Printf("Enter Time (00:00) or x to finish: ")
				scanner.Scan()
				time := scanner.Text()
				if time == "x" {
					timeTable["Friday"] = times
					break
				}
				times = append(times, time)
			}
			fmt.Printf("\n")
			subject.TimeTable = timeTable
			fmt.Printf("Subject Time Table Changed: %v", subject.TimeTable)
			break
		}
	}
}

// Assing User Menu
func assignUser() {
	for {
		var userMenu string = "Assign User Menu"
		var assignT string = "Assign Teacher"
		var assignS string = "Assign Student"
		var mm string = "Go back to Main Menu"

		fmt.Println("\n\n\n\n\n ")
		fmt.Printf("\n\n %#25q \n\n", userMenu)
		fmt.Printf("/assignTeacher %#30q \n", assignT)
		fmt.Printf("/assignStudent %#30q \n\n", assignS)
		fmt.Printf("/mainMenu %#41q \n", mm)
		fmt.Printf("\n")
		fmt.Printf("Choice: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userOption := scanner.Text()
		if userOption == "/assignTeacher" {
			assignTeacher()
		} else if userOption == "/assignStudent" {
			assignStudent()
		} else if userOption == "/mainMenu" {
			mainMenu()
		}
	}
}

// Assign Teacher
func assignTeacher() {
	var teacher *Teacher
	var class *Class

	viewTeacher()
	fmt.Printf("\n")
	fmt.Printf("Please Enter The Teacher You Would Like To Assign: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tChoice := scanner.Text()
	for _, t := range TEACHERS {
		if strings.ToLower(tChoice) == strings.ToLower(t.FirstName) {
			teacher = &*t
			break
		}

	}
	viewClass()
	fmt.Printf("\n")
	fmt.Printf("Please Enter The Class You Would Like To Assign %v To: ", teacher.FirstName)
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cChoice := scanner.Text()
	for _, c := range CLASSES {
		if cChoice == c.ClassCode {
			class = &*c
			break
		}
	}
	class.Teacher = teacher
}

// Assign Student
func assignStudent() {
	var student *Student
	var class *Class

	viewStudent()
	fmt.Printf("\n")
	fmt.Printf("Please Enter The Student You Would Like To Assign: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sChoice := scanner.Text()
	for _, s := range STUDENTS {
		if strings.ToLower(sChoice) == strings.ToLower(s.FirstName) {
			student = &*s
			break
		}
	}
	viewClass()
	fmt.Printf("\n")
	fmt.Printf("Please Enter The Class You Would Like To Assign %v To: ", student.FirstName)
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cChoice := scanner.Text()
	for _, c := range CLASSES {
		if cChoice == c.ClassCode {
			class = &*c
			break
		}
	}
	class.Students = append(class.Students, student)
}

// Assigning Subjects to Teachers and Students
func assignSubject() {
	var toBeAssigned Subject
	var tChoice string
	var sChoice string
	var student *Student

	scanner := bufio.NewScanner(os.Stdin)
	viewSubject()
	fmt.Printf("Please Enter The Subject You Would Like To Assign: ")
	scanner.Scan()
	subject := scanner.Text()
	for _, v := range SUBJECTS {
		if strings.ToLower(v.Name) == strings.ToLower(subject) {
			toBeAssigned = *v
		}
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Would you like to Assign %v to a Student or Teacher?: ", toBeAssigned.Name)
		scanner.Scan()
		choice := scanner.Text()
		if strings.ToLower(choice) == "teacher" {
			scanner := bufio.NewScanner(os.Stdin)
			viewTeacher()
			fmt.Printf("Please Enter The Name Of The Teacher You Would Like To Assign %v To: ", toBeAssigned.Name)
			scanner.Scan()
			tChoice = scanner.Text()
			for _, v := range TEACHERS {
				if strings.ToLower(tChoice) == strings.ToLower(v.FirstName) {
					sub := *v.Subjects
					sub = append(sub, toBeAssigned)
				}
			}

		} else if strings.ToLower(choice) == "student" {
			scanner := bufio.NewScanner(os.Stdin)
			viewStudent()
			fmt.Printf("Please Enter The Name Of The Student You Would Like To Assign %v To: ", toBeAssigned.Name)
			scanner.Scan()
			sChoice = scanner.Text()
			for _, v := range STUDENTS {
				if strings.ToLower(v.FirstName) == strings.ToLower(sChoice) {
					student = *&v
					student.Grades[toBeAssigned.Name] = 0
				}

			}
		}
	}

}

// Main Function
func main() {
	mainMenu()
}
