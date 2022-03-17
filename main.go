package main

import (
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) <= 2 {
		log.Println("终止运行，启动参数必须包含ip:port,以及文件路径")
		return
	}
	mime.AddExtensionType(".apk", "application/vnd.android.package-archive")

	err := startFileServer(os.Args[1])
	if err != nil {
		log.Println(err)
	}
}

func startFileServer(addr string) error {
	serverFilePath := os.Args[2]
	log.Println(serverFilePath)

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		name := strings.Replace(r.URL.Path, "/static", serverFilePath, 1)
		_, file := filepath.Split(name)
		log.Println("path", file)
		if file != "" {
			cd := mime.FormatMediaType("attachment", map[string]string{"filename": file})
			w.Header().Set("Content-Disposition", cd)
			if strings.HasSuffix(file, "apk") {
				w.Header().Set("Cache-Control", "no-cache")

				w.Header().Set("Content-Type", "application/vnd.android.package-archive")
			}
		}
		http.ServeFile(w, r, name)
	})

	err := http.ListenAndServe(addr, nil)
	return err
}
