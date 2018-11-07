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
		if id, body, err := conn.Reserve(50 * time.Second); err != nil {
			fmt.Print(id, " : ", body)
		}else {
			fmt.Println(id,": ", string(body))
			conn.Delete(id)
		}
	}
}
