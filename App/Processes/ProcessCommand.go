package Processes

import (
	"encoding/json"
	"fmt"
	"net"

	"../Helper"
)

type tCommand struct {
	Command string          `json:"command"`
	Data    json.RawMessage `json:"data"`
}

//////////login
type tParamsLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (obj tParamsLogin) process() error {

	fmt.Println("login.process")
	fmt.Println(obj.Name)
	fmt.Println(obj.Password)

	return nil
}

//////////get_icon_by_user
type tParamsGetIconByUser struct {
	UserId int `json:"user_id"`
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

func processCommand(bufData tChanBufData, conn net.Conn) error {

	var commandData = &tCommand{}

	if err := Helper.ParseJsonIntoStruct(bufData.buf, commandData); err != nil {
		return fmt.Errorf("processCommand: %s: ", err)
	}

	switch commandData.Command {
	case "login": // {"Command": "login", "Data": {"name":"vasya", "password":"qwerty"}}

		var params = &tParamsLogin{}
		if err := Helper.ParseJsonIntoStruct(commandData.Data, params); err != nil {
			return cantParseParams("login", err)
		}
		params.process()

		break
	case "get_icon_by_user": // {"Command": "get_icon_by_user", "Data": {"user_id":1234}}
		var params = &tParamsGetIconByUser{}
		if err := Helper.ParseJsonIntoStruct(commandData.Data, params); err != nil {
			return cantParseParams("login", err)
		}
		params.process()

		break
	case "get_chats_list": // {"Command": "get_chats_list", "Data": {"user_id":1234}}
		var params = &tParamsGetChatsList{}
		if err := Helper.ParseJsonIntoStruct(commandData.Data, params); err != nil {
			return cantParseParams("login", err)
		}
		params.process()

		break
	case "get_chat": // {"Command": "get_chat", "Data": {"chat_id":1234}}
		var params = &tParamsGetChatsList{}
		if err := Helper.ParseJsonIntoStruct(commandData.Data, params); err != nil {
			return cantParseParams("login", err)
		}
		params.process()

		break
	}

	transferData(conn, []byte("Your command: "+commandData.Command+"\n"))

	return nil
}

func cantParseParams(methodName string, err error) error {
	return fmt.Errorf("cant parse Data ("+methodName+") %s: ", err)
}
