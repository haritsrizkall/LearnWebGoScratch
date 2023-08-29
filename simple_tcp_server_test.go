package learnwebgoscratch

import (
   	"bufio"
	"fmt"
	"net"
    "testing"
	"strings"
	"time"
)

func TestSimpleTCP(t *testing.T) {
    fmt.Println("Hello World")
    li, err := net.Listen("tcp", ":8090")
    if err != nil {
        panic(err)
    }
    defer li.Close()
    for {
        conn, err := li.Accept()
        conn.SetDeadline(time.Now().Add(10 * time.Second))
        if err != nil {
            panic(err)
        }
        defer conn.Close()
        go handler(conn)
    }
}

func handler(conn net.Conn) {
    method, url := request(conn)
    mux(conn, method, url)
}

func mux(conn net.Conn, method, url string) {
    switch {
    case method == "GET" && url == "/":
        fmt.Println("Method Get")
        index(conn)
    case method == "GET" && url == "/about":
        fmt.Println("Another Method")
        about(conn)
    case method == "GET" && url == "/support":
        support(conn)
    default:
        notfound(conn) 
    }
}

func request(conn net.Conn) (string, string) {
    scanner := bufio.NewScanner(conn)
    method := ""
    url := ""
    idx := 0
    for scanner.Scan() {
        text := scanner.Text()
        if text == "" {
            break
        }
        switch idx {
        case 0:
            textSplit := strings.Split(text, " ")
            method = textSplit[0]
            url = textSplit[1]
            fmt.Println(method, url)
        default:
            fmt.Println(text)
        }
        idx++
    }  
    return method, url
}

func index(conn net.Conn) {
    const body = "Hello This Reponse From Get Request and Index URL\n"
    fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprintf(conn, "\r\n")
    fmt.Fprint(conn, body)
}        

func about(conn net.Conn) {
    const body = "Hello This Reponse From Get Request and About URL\n"
    fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprintf(conn, "\r\n")
    fmt.Fprint(conn, body)
} 

func support(conn net.Conn) {
    const body = "Hello This Reponse From Get Request and Support URL\n"
    fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprintf(conn, "\r\n")
    fmt.Fprint(conn, body)
} 

func notfound(conn net.Conn) {
    const body = "Hello This Reponse From Get Request and Support URL\n"
    fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprintf(conn, "\r\n")
    fmt.Fprint(conn, body)
} 
 

