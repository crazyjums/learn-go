package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/download", downloadHandler2)
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func downloadHandler2(w http.ResponseWriter, r *http.Request) {
	// 设置要下载的文件路径
	filePath := "./output.csv"

	// 获取数据（假设数据为一个二维字符串切片）
	data := [][]string{
		{"Name", "Age", "Country"},
		{"John", "25", "USA"},
		{"Emily", "27", "Canada"},
		{"Tom", "30", "UK"},
	}

	// 创建 CSV 文件
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Failed to create file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 创建 CSV writer
	writer := csv.NewWriter(file)

	// 写入数据
	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
			log.Println("Error writing row:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	writer.Flush()

	// 设置响应头，告诉浏览器该文件应该被下载
	w.Header().Set("Content-Disposition", "attachment; filename="+filePath)
	w.Header().Set("Content-Type", "text/csv")

	// 将文件内容写入响应
	_, err = file.Seek(0, 0) // 将文件指针移动到文件开头
	if err != nil {
		http.Error(w, "Failed to seek file", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to write file content", http.StatusInternalServerError)
		return
	}
	go func() {
		time.Sleep(time.Second)
		err = os.Remove(filePath)
		if err != nil {
			return
		}
	}()
}
