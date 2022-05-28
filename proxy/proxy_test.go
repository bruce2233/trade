package proxy

import (
	"io/ioutil"
	"net"
	"testing"
)

func TestProxy(t *testing.T) {
	Start()
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:996")
	if err != nil {
		t.Log("dial error!")
	}
	p := make([]byte, 1024)
	conn.Write([]byte("from client"))
	conn.Read(p)
	t.Log(string(p))
}

func TestHttp(t *testing.T) {
	// l, err := net.Listen("tcp", ":996")
	// if err != nil {
	// 	t.Log(err)
	// }

	// conn, err := l.Accept()
	// bufReader := bufio.NewReader(conn)
	// req, err := http.ReadRequest(bufReader)
	// t.Log(req)

	// client := &http.Client{}
	// resp, _ := client.Get(req.URL)
	// resp.Write(conn)
	// t.Log(resp)
}

func TestFilePath(t *testing.T) {
	b, err := ioutil.ReadFile("server.key")
	if err != nil {
		t.Log("file error")
	}
	t.Log(string(b))
}
