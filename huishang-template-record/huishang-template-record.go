package main

import (
	"fmt"
	"github.com/beanstalkd/go-beanstalk"
	"time"
)

func main() {
	if conn, e := beanstalk.Dial("tcp", "127.0.0.1:14711"); e != nil {
		panic(e)

	} else {
		println(" body2")

		if id, body, err := conn.Reserve(12300*time.Second); err != nil {
			fmt.Print(id, " : ", body)
		}else {
			println(id," :  ",string(body))
			conn.Delete(id)
		}

	}
}

func execWork(){


}