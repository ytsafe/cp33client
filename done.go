package main

import (
	"fmt"
	"time"

	//"regexp"
	"strconv"

	"github.com/henrylee2cn/surfer"
)

var notice = make(chan string, 10000)

func loopDoneNotice() {
	var err error
	for {
		select {
		case temp := <-notice:
			_, err = surfer.Download(&surfer.Request{
				Url:          temp,
				DownloaderID: 1,
			})
			if err != nil {
				time.Sleep(time.Second)
				notice <- temp
			}
			//resp.Body.Close()
		}
	}
}

func (self *data) done() {
	var err error
	u := "http://127.0.0.1:8080/dataInNotice/" + strconv.Itoa(self.Type) + "/" + strconv.Itoa(self.Issue)
	_, err = surfer.Download(&surfer.Request{
		Url: u,
		//DownloaderID: 1,
	})
	if err != nil {
		fmt.Println("done() 22:", err)
		time.Sleep(time.Second)
		notice <- u
		return
	}
	//defer resp.Body.Close()
	//defer surfer.DestroyJsFiles()
	redisClient.Set("Client_"+strconv.Itoa(self.Type)+"_"+strconv.Itoa(self.Issue), self.Data, time.Minute*30)
	fmt.Println(time.Now(), " 成功入库：", "	", self.Type, self.Issue, "		", self.Data)
}
