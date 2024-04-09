package main

import (
	"bufio"
	"fmt"
	"github.com/pion/webrtc/v4"
	"go-meeting/internal/helper"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// // 1. create peer connection
	connection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		return
	}
	defer func() {
		if e := connection.Close(); e != nil {
			log.Panicln(e.Error())
		}
	}()
	// 2. create data channel
	channel, err := connection.CreateDataChannel("foo", nil)
	channel.OnOpen(func() {
		log.Println("data channel has opened")
		i := 1000
		// 开启一个周期5秒的定时任务
		for range time.NewTicker(time.Second * 5).C {
			err := channel.SendText("offer: hello world " + strconv.Itoa(i))
			if err != nil {
				log.Println(err.Error())
			}
		}
	})
	channel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Println(string(msg.Data))
	})
	// 3. create offer
	offer, err := connection.CreateOffer(nil)
	if err != nil {
		return
	}
	// 4. set local description
	err = connection.SetLocalDescription(offer)
	if err != nil {
		return
	}
	// 5. print offer
	fmt.Println("OFFER: ")
	fmt.Println(helper.Encode(offer))
	// 6. input answer
	fmt.Println("请输入ANSWER: ")
	var answer webrtc.SessionDescription
	answerStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	helper.Decode(answerStr, &answer)
	// 7. set remote description
	err = connection.SetRemoteDescription(answer)
	if err != nil {
		panic(err)
	}
	// 阻塞1h
	<-time.After(time.Hour)
}
