package main

import "fmt"

func Receiver(v interface{}) {
		switch v.(type) {
		case bool:
			fmt.Printf("这个是bool\n")
		case float64:
			fmt.Printf("这个是float64\n")
		case int,int64:
			fmt.Printf("这个是int\n")
		case nil:
			fmt.Printf("这个是nil\n")
		case string:
			fmt.Printf("这个是string\n")
		default:
			fmt.Printf("这个属于非常规类型\n")
		}
}

func main(){

	Receiver("你好吗")
	Receiver(32)
	Receiver(true)

}