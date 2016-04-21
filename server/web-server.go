package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	//"strings"
	"github.com/gorilla/websocket"
	log "github.com/Sirupsen/logrus"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", 405)
        return
    }

    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.WithField("err", err).Println("Upgrading to websockets")
        http.Error(w, "Error Upgrading to websockets", 400)
        return
    } else {
		fmt.Println("new websocket connection")
	}

    for {
        mt, data, err := ws.ReadMessage()
        ctx := log.Fields{"mt": mt, "data": data, "err": err}
        if err != nil {
            if err == io.EOF {
                log.WithFields(ctx).Info("Websocket closed!")
            } else {
                log.WithFields(ctx).Error("Error reading websocket message")
            }
            break
        }
        switch mt {
        case websocket.TextMessage:
			var response string;
			fmt.Println("message received from client : ", string(data))
			response = "Hey client! I have received your \"" + string(data) + "\" message! Thank you! \n\n THE GO SERVER"
			ws.WriteMessage(mt, []byte(response))

            /*msg, err := validateMessage(data)
            if err != nil {
                ctx["msg"] = msg
                ctx["err"] = err
                log.WithFields(ctx).Error("Invalid Message")
                break
            }*/
        default:
            log.WithFields(ctx).Warning("Unknown Message!")
        }
    }

    ws.WriteMessage(websocket.CloseMessage, []byte{})
}

func loadIndex(w *http.ResponseWriter) {
	fmt.Println("loading index page")
	pwd, _ := os.Getwd()
	filepath := pwd + "/../client/index.html"
	index, _ := ioutil.ReadFile(filepath)

	io.WriteString(*w, string(index))
}

func serveFile(w *http.ResponseWriter, filePath string) {
	fmt.Println("serving page " + filePath)
	pwd, _ := os.Getwd()
	filePath = pwd + "/../client" + filePath
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		loadIndex(w)
	} else {
		io.WriteString(*w, string(file))
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("new client request")
	//fmt.Println(r.Header)
	if val, ok := r.Header["Sec-Websocket-Version"]; ok && val[0] == "13" {
		handleWebsocket(w, r)
	} else if r.URL.Path == "/" {
		loadIndex(&w)
	} else {
		serveFile(&w, r.URL.Path)
	}
}

func main() {
	fmt.Println("Initializing handler...")
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
