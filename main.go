package main

import (
	"net/http"
	"os"
	"log"
)

func main() {
	if len(os.Args) <= 2 {
		log.Println("终止运行，启动参数必须包含ip:port,以及文件路径")
		return
	}
	err := startFileServer(os.Args[1])
	if err != nil {
		log.Println(err)
	}
}

func startFileServer(addr string) error {
	serverFilePath := os.Args[2]
	log.Println(serverFilePath)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(serverFilePath))))
	err := http.ListenAndServe(addr, nil)
	return err
}
