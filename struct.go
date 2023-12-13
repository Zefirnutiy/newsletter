package main


type Letter struct {
	Header 	string  `json:header`
	Body 	string  `json:body`
}

type Sender struct {
	Email 		string `json:email`
	Password 	string `json:password`
}

type SendData struct {
	Letter 		Letter	`json:letter`
	Sender 		Sender	`json:sender`
	Recipients  string	`json:recipients`
}