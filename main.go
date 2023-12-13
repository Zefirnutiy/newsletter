package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func main(){
	router := gin.Default()

	router.Use(cors.Default())
	router.POST("api/v1/sendLetters", SendLetters)

	router.Run(":8080")
}



func SendLetters(ctx *gin.Context) {
	var sendData SendData  
	err := ctx.BindJSON(&sendData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":"отправь запрос нормально, блять",
		})
		log.Print("проблема с json фронта: ", err)
		return
	}

	if err := send(sendData); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "the server cannot process your request",
		})
		log.Print("произошла ошибка при обработке запроса отправки писем: ", err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func send(sendData SendData) error {
	for _, recipient := range strings.Split(sendData.Recipients, " ") {

		m := gomail.NewMessage()

		m.SetHeader("From", sendData.Sender.Email)
		
		m.SetHeader("To", recipient)
		
		m.SetHeader("Subject", sendData.Letter.Header)
		
		m.SetBody("text/plain", sendData.Letter.Body)
		d := gomail.NewDialer("smtp.gmail.com", 587, sendData.Sender.Email, sendData.Sender.Password)
		
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		
		err := d.DialAndSend(m);
		if err != nil {
			log.Fatal("какая-то error в функции отправки (send): ", err)
			return err
		}

	}
  
	return nil
}
