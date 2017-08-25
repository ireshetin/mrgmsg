package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./models/user"
)

//var user = person{id: 0, name: "", password: ""}

type authData struct {
	userName     string
	userPassowrd string
}

type person struct {
	id       int
	name     string
	password string
}

func main() {
	command := ""
	result := true
	code := 0
	for result {
		command = getCommand("Введите команду (help)")
		result, code = processCommand(command)

		if code > 0 {
			consoleLog("some error")
		}
	}

	consoleLog("Fin")
}

func getCommand(msg string) string {
	consoleLog(msg)

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func processCommand(command string) (bool, int) {

	switch command {
	case "help":
		consoleLog("you can type: quit | help | login | register | whoami | getmylist | logout")
		break

	case "login":
		user, _ = processLogin()
		//if (code > 0) {
		//	//do something
		//}
		break

	case "logout":
		user = getEmptyUser()
		break

	case "whoami":
		consoleLog(strconv.Itoa(user.id) + " : " + user.name)
		break

	case "quit":
		consoleLog("userId: " + strconv.Itoa(user.id))
		return false, 0
		break

	case "register":
		var authdata authData
		authdata, code := processRegister()
		if code > 0 {
			consoleLog("Ошибка")
		}
		user, _ = auth(authdata)

		consoleLog(user.name)
		consoleLog(user.password)
		consoleLog(string(code))
		break

	case "getmylist":
		processMyList(user)
		break

	default:
		consoleLog("command not found")

	}

	return true, 0
}

func processMyList(user person) {
	if user.id > 0 {
		getMyList(user)
	} else {
		consoleLog("Forbidden =(")
	}
}

func getMyList(user person) {
	consoleLog("This is " + user.name + "'s list ")
}

func processRegister() (authData, int) {
	userName := getCommand("Введите user name")
	// TODO: проверка на занятость логина пользователя
	// TODO: createUser

	userPassowrd := generatePassword()

	return authData{userName: userName, userPassowrd: userPassowrd}, 0
}

func generatePassword() string {
	return "qwe"
}

func processLogin() (person, int) {
	userName := getCommand("Введите user name")
	userPassword := getCommand("Введите user passowrd")

	return auth(authData{userName: userName, userPassowrd: userPassword})
}

func auth(authdata authData) (person, int) {
	// TODO: logAuth
	return getUserByAuth(authdata)
}

func getUserByAuth(authdata authData) (person, int) {
	// TODO
	if "vasya" == authdata.userName && "qwe" == authdata.userPassowrd {
		consoleLog("correct")
		return person{id: 555, name: authdata.userName, password: authdata.userPassowrd}, 0
	}

	return getEmptyUser(), 404
}

func consoleLog(msg string) {
	fmt.Println(msg)
}

func getEmptyUser() person {
	return person{id: 0, name: "", password: ""}
}
