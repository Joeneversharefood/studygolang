package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")

	if nil != err {
		fmt.Printf("dial redis failed\n")
		return
	}

	fmt.Printf("dial redis success\n")

	_, err = conn.Do("set", "myage", 23)

	if nil != err {
		fmt.Printf("set redis failed\n")
		return
	}

	fmt.Printf("set to redis success\n")

	result, err := redis.Int(conn.Do("get", "myage"))

	if nil != err {
		fmt.Printf("get from redis failed\n")
		return
	}

	fmt.Println(result)

	defer conn.Close()
}
