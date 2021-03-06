package main

import "fmt"
import "net"
import "bufio"
import "io/ioutil"
import "os"

func host(filename string) {
    ln, err := net.Listen("tcp", ":8080")
    defer ln.Close()

    if err != nil {
        fmt.Println("error server: " + err.Error())
        return
    }

    for {
        file, ferr := os.Open(filename)
        defer file.Close()

        if ferr != nil {
            fmt.Println("error opening file: " + ferr.Error())
            os.Exit(2)
        }

        conn, cerr := ln.Accept()

        if cerr != nil {
            continue
        }


        info, serr := file.Stat()

        if serr != nil {
            fmt.Println("error stating file: " + serr.Error())
            os.Exit(3)
        }

        arr := make([]byte, info.Size())
        file.Read(arr)
        conn.Write(arr)
        conn.Close()
        file.Close()
    }
}

func receive(ip string) {
    conn, err := net.Dial("tcp", ip + ":8080")

    if err != nil {
        fmt.Println("error client: " + err.Error())
        return
    }

    bytes, _ := ioutil.ReadAll(bufio.NewReader(conn))
    fmt.Println(string(bytes))
}

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("no action specified; exiting")
        os.Exit(0)
    }

    if os.Args[1] == "host" {
        fmt.Println("hosting file")
        if len(os.Args) <= 2 {
            fmt.Println("no file specified; exiting")
            os.Exit(1)
        }
        host(os.Args[2])
    } else if os.Args[1] == "receive" {
        fmt.Println("recieving file")
        ip := "127.0.0.1"
        if len(os.Args) > 2 {
            ip = os.Args[2]
        }
        receive(ip)
    } else {
        fmt.Println("unknown action '" + os.Args[1] + "'; exiting")
        os.Exit(-1)
    }
}
