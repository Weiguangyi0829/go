package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
	"time"
)

type Adapter struct {
	pool sync.Pool
}

func (A *Adapter) find(filedata multipart.File,filename string) []byte  {
	fmt.Println("456")
	buffer := A.pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer func() {
		if buffer != nil {
			A.pool.Put(buffer)
			buffer = nil
		}
	}()

	//fileBytes, err := ioutil.ReadAll(filedata)
	data, err := io.Copy(buffer,filedata)
	fmt.Println("data",data)
	//fmt.Println("buffer",buffer)
	fileBytes, _  := json.Marshal(data)
	if err != nil{
		panic(err)
	}
	destFile, err := os.Create("./upload/" + filename)
	if err != nil {
		log.Printf("Create failed: %s\n", err)
		return nil
	}
	defer destFile.Close()

	// 读取表单文件，写入保存文件
	data1, err := io.Copy(destFile, filedata)
	if data == data1{
		fmt.Println("true")
	}
	if err != nil {
		log.Printf("Write file failed: %s\n", err)
		return nil
	}

	A.pool.Put(buffer)
	buffer = nil
	return fileBytes
}

func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1")
	response := map[string]interface{}{
		"ref" : time.Now().UnixNano(),
	}
	A := Adapter{pool:sync.Pool{
		New: func() interface{}{
			return bytes.NewBuffer(make([]byte,4098))
		},
	}}
	//---------------从网页那边读取数据---------------//
	filedata,header, err := r.FormFile("file")
	defer filedata.Close()
	if err != nil{
		response["status"] = "The upload is fail,Please try aginst!"
		if err := json.NewEncoder(w).Encode(response) ; err != nil{
			http.Error(w, err.Error(), 500)
			return
		}
	}
	fmt.Println(2)
	//session := db.Init()
	//defer session.Close()
	fileBytes := A.find(filedata,header.Filename)
	fileType := http.DetectContentType(fileBytes)
	if fileType != "image/jpeg"  &&  fileType != "image/gif" {
		response["status"] = "The file'sType is wrong,Please try aginst!"
		if err := json.NewEncoder(w).Encode(response) ; err != nil{
			http.Error(w, err.Error(), 500)
			return
		}
	}
	response["status"] = "upload successfully"
	if err := json.NewEncoder(w).Encode(response) ; err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
}

func Upload2(w http.ResponseWriter, r *http.Request)  {

}

func Download(w http.ResponseWriter, r *http.Request) {

}
