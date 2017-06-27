package main

import (

	alexa "github.com/mikeflynn/go-alexa/skillserver"
	//"github.com/cloudfoundry-community/go-cfenv"
	//"github.com/cloudnativego/cf-tools"
	//"net/http"
	//"github.com/gorilla/context"
	//"strings"
	//"log"
	//"fmt"
	//"fmt"
	//"gopkg.in/mgo.v2"
	"log"
	//"strconv"
	//"encoding/json"
	"os"
	"fmt"
)

var Applications = map[string]interface{}{
	"/echo/helloworld": alexa.EchoApplication{ // Route
		AppID:    "amzn1.ask.skill.3731f8ac-19e2-4530-b740-746a460a84c1", // Echo App ID from Amazon Dashboard
		OnIntent: EchoIntentHandler,
		OnLaunch: EchoIntentHandler,
	},
}

func EchoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	//echoResp.OutputSpeech("Hello world from my new Echo test app!").Card("Hello World", "This is a test card.")


	if echoReq.GetRequestType() == "LaunchRequest" {

		echoResp.OutputSpeech("You have reached, Launch Request").EndSession(false)


	} else if echoReq.GetRequestType() == "IntentRequest" {
		log.Println(echoReq.GetIntentName())

		fmt.Println("Intent Name === " + echoReq.GetIntentName())

		switch echoReq.GetIntentName() {
		case "HelloWorld":

			echoResp.OutputSpeech("Hello, You have reached Hello World Intent").EndSession(true)
		        break;
		case "Test":

			echoResp.OutputSpeech("Hello, You have reached Test Intent").EndSession(true)
			break;

		default:
			echoResp.OutputSpeech("I'm sorry, I didn't get that. Can you say that again?").EndSession(false)
		}


	} else if echoReq.GetRequestType() == "SessionEndedRequest" {
		//session.Delete(col)
		echoResp.OutputSpeech("Hello, You have Ended Session").EndSession(true)
	}

}

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	alexa.Run(Applications, port)
}



/*
type MySession struct {
	AWSID           string
	Dollars         int
	NumQuestions    int
	//CurrentQuestion JeopardyQuestion
	UpdatedAt       int64
}


type User struct{
	Id      string
	Balance uint64
}



func MyHelloWorld (w http.ResponseWriter, r *http.Request) {

	echoReq := context.Get(r, "echoRequest").(*alexa.EchoRequest)


	log.Println(echoReq.GetRequestType())
	log.Println(echoReq.GetSessionID())

	if echoReq.GetRequestType() == "LaunchRequest" {

		echoResp := alexa.NewEchoResponse().OutputSpeech("You have reached, Launch Request").EndSession(true)
		json, _ := echoResp.String()
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write(json)

	} else if echoReq.GetRequestType() == "IntentRequest" {
		log.Println(echoReq.GetIntentName())




		switch echoReq.GetIntentName() {
		case "HelloWorld":
			echoResp = alexa.NewEchoResponse().OutputSpeech("Hello, You have reached Hello World Intent").EndSession(true)
		default:
			echoResp = alexa.NewEchoResponse().OutputSpeech("I'm sorry, I didn't get that. Can you say that again?").EndSession(false)
		}

		json, _ := echoResp.String()
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write(json)
	} else if echoReq.GetRequestType() == "SessionEndedRequest" {
		//session.Delete(col)
	}
}



*/