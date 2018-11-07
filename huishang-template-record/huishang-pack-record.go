package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func task(ch <-chan []string) {
	for value := range ch {
		var f interface{}
		json.Unmarshal([]byte(value[1]), &f)
		fmt.Printf("%+v", f)
		time.Sleep(5 * time.Second)
	}
}
func main() {
	client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	defer client.Close()

	var tmp = `{
    "openid": "o-ypr5XTJ_2A2zF8M6Exena5sKK4",
    "formid": "9a00c6f7292c1758f30f9f9b45f6a419",
    "keyword": "收到回赏通知",
    "page": "pages/unpack/main?riddle=aa435491acae5568ced57b9d542d30a9",
    "data": {
        "keyword1": {
            "value": "Boss",
            "color": "#173177"
        },
        "keyword2": {
            "value": "0.01元",
            "color": "#173177"
        },
        "keyword3": {
            "value": "2018-11-06 13:42:44",
            "color": "#173177"
        },
        "keyword4": {
            "value": "恭喜发财，大吉大利。",
            "color": "#173177"
        }
    }
}`

	var f interface{}
	json.Unmarshal([]byte(tmp),&f)

	//fmt.Printf("%+v",f)



	for i := 0; i < 100; i++ {
		//client.LPush("test",fmt.Sprintf("test :  %d", i))
		client.LPush("test", tmp)
	}

	ch := make(chan []string)
	go task(ch)
	go task(ch)
	go task(ch)
	go task(ch)

	for {
		job := client.BLPop(365*24*time.Hour, "test")
		ch <- job.Val()
	}
}
