package Processes

import (
	"encoding/json"
	"fmt"
	"net"

	"time"

	"sync"

	"../../models/user"
	"../Helper"
)

//type tCurrentUser struct {
//
//}

type tCommand struct {
	Command string          `json:"command"`
	Params  json.RawMessage `json:"params"`
}

//////////login
type tParamsLogin struct {
	Login    string `json:"username"`
	Password string `json:"password"`
}

func (obj tParamsLogin) process() error {
	reply := true
	empty := ""

	fmt.Println("login.process")
	fmt.Println(obj.Login)
	fmt.Println(obj.Password)

	user := user.User{
		Login:    newStr(obj.Login),
		Name:     &empty,
		LastName: &empty,
		Password: newStr(obj.Password),
	}

	_ = user.CheckPassword(user, &reply)

	//fmt.Println(user.Id)
	//fmt.Printf("%s", user.Password)
	//fmt.Printf("%s", user.Login)
	//fmt.Println(user.Password)
	fmt.Println(*user.Login)
	//fmt.Println(*user.Id)

	return nil
}

func newStr(s string) *string {
	return &s
}

//////////get_icon_by_user
type tParamsGetIconByUser struct {
	UserId int `json:"user_id"`
}

//////////get_icon_by_user
type tParamsFullInfo struct {
	UserId int `json:"id"`
}

func (obj tParamsGetIconByUser) process() error {

	fmt.Print("GetIconByUser.process")
	fmt.Println(obj.UserId)

	return nil
}

//////////get_chats_list
type tParamsGetChatsList struct {
}

func (obj tParamsGetChatsList) process() error {

	fmt.Print("GetChatsList.process")
	fmt.Println("none")

	return nil
}

type tChanList struct {
	UserList string
}

type tChanIcon struct {
	UserAvatars string
}

func getList(wg *sync.WaitGroup, ch chan tChanList) {
	//TODO: Брать данные из БД)))

	ch <- tChanList{"Katya\nPetya\nSasha\nVasya\nNastya\nMasha\nVanya"}

	time.Sleep(time.Second * 5) //долгий запрос в БД)

	wg.Done()

}

func getIcons(wg *sync.WaitGroup, ch chan tChanIcon) {
	//TODO: Брать данные из БД)))

	ch <- tChanIcon{" :) \n༼ つ ◕_◕ ༽つ\n(づ｡◕‿‿◕｡)づ\n(◕‿◕✿)\n◉_◉\n(｡◕‿‿◕｡)\n(｡◕‿◕｡)\n(ʘᗩʘ')"}

	time.Sleep(time.Second * 3) //долгий запрос в БД)

	wg.Done()
}

func (obj tParamsFullInfo) process() error {
	//#ASYNC

	//TODO: Брать данные из БД)))
	chanIcon := make(chan tChanIcon, 1)
	chanList := make(chan tChanList, 1)

	var wg sync.WaitGroup

	wg.Add(1)
	go getIcons(&wg, chanIcon)
	wg.Add(1)
	go getList(&wg, chanList)

	fmt.Println("===> Waiting for routines")
	wg.Wait()
	fmt.Println("Waiting done")

	list := <-chanList
	icon := <-chanIcon

	fmt.Print("====>>> tParamsFullInfo.process: ")
	fmt.Println(list)
	fmt.Println(icon)

	return nil
}

func processCommand(bufData tChanBufData, conn net.Conn) error {

	var commandData = &tCommand{}

	if err := Helper.ParseJsonIntoStruct(bufData.buf, commandData); err != nil {
		return fmt.Errorf("processCommand: %s: ", err)
	}

	switch commandData.Command {
	case "login": // {"Command": "login", "Params": {"name":"vasya", "password":"qwerty"}}

		var params = &tParamsLogin{}
		if err := Helper.ParseJsonIntoStruct(commandData.Params, params); err != nil {
			fmt.Printf("->>>> %s", commandData.Params)
			return cantParseParams("login: ", err)
		}
		params.process()

		break
	case "get_full_info": // {"Command": "get_full_info", "Params": {"user_id":1234}}

		var params = &tParamsFullInfo{}
		if err := Helper.ParseJsonIntoStruct(commandData.Params, params); err != nil {
			return cantParseParams("get_icon_by_user", err)
		}
		params.process()

		break
	case "get_chats_list": // {"Command": "get_chats_list", "Params": {"user_id":1234}}
		var params = &tParamsGetChatsList{}
		if err := Helper.ParseJsonIntoStruct(commandData.Params, params); err != nil {
			return cantParseParams("get_chats_list", err)
		}
		params.process()

		break
	case "get_chat": // {"Command": "get_chat", "Params": {"chat_id":1234}}
		var params = &tParamsGetChatsList{}
		if err := Helper.ParseJsonIntoStruct(commandData.Params, params); err != nil {
			return cantParseParams("get_chat", err)
		}
		params.process()

		break
	}

	transferData(conn, []byte("Your command: "+commandData.Command+"\n"))

	return nil
}

func cantParseParams(methodName string, err error) error {
	return fmt.Errorf("cant parse Params ("+methodName+") %s: ", err)
}
