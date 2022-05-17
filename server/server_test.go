package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
	"testing"
	"trade/order"
)

func TestStart(t *testing.T) {
	// 1、与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("conn server failed, err:%v\n", err)
		return
	}
	// 2、使用 conn 连接进行数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}

		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed, err:%v\n", err)
			return
		}
		// 从服务端接收回复消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed:%v\n", err)
			return
		}
		fmt.Printf("收到服务端回复:%v\n", string(buf[:n]))
	}
}

func TestMyNet(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("connection error!")
	}
	s := []byte("Internet by Go is cool!")
	conn.Write(s)
}

func TestMarshal(t *testing.T) {

	orderBook := order.OrderBook{}
	var orderBookJson []byte
	orderBookJson, _ = json.Marshal(orderBook)
	t.Log(string(orderBookJson))

	orderBook.SellOrders = append(orderBook.SellOrders, order.SellOrder{Order: order.GenerateOrder()})
	orderBook.SellOrders = append(orderBook.SellOrders, order.SellOrder{Order: order.GenerateOrder()})
	orderBookJson, _ = json.Marshal(orderBook)
	t.Log(string(orderBookJson))

	orderBook.BuyOrders = append(orderBook.SellOrders, order.BuyOrder{Order: order.GenerateOrder()})
	orderBook.BuyOrders = append(orderBook.SellOrders, order.BuyOrder{Order: order.GenerateOrder()})

	order1 := order.GenerateOrder()
	order1JSON, _ := json.Marshal(order1)
	t.Log("order1JSON: ", order1JSON)
	t.Log("order1JSON: ", string(order1JSON))

	var ip IParent
	var parent Parent
	var child Child
	var b []byte

	child = Child{Str: "I am a child node...", Num: 996}
	parent = Parent{Child: child}
	ip = parent

	b, _ = json.Marshal(ip)
	t.Log("Marshal IParent", string(b))

	c, _ := json.Marshal(parent)
	t.Log("Marshal Parent", string(c))

	d, _ := json.Marshal(child)
	t.Log("Marshal Child", string(d))
	t.Log("log", child)

	t.Log("type of ip.(Parent)", reflect.TypeOf(ip.(Parent)))
	b, _ = json.Marshal(ip.(Parent))
	t.Log("Marshal IParent type change", string(b))

}

type IParent interface {
}

type Parent struct {
	Child
}

type Child struct {
	Str string
	Num int `json:"NUMBER"`
}
