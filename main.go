package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", ":80", "请输入端口")
	post := flag.String("post", "", "请输入端口")
	flag.Parse()
	if *post != "" {
		data := []byte("content=%3C%3Fxml+version%3D%221.0%22+encoding%3D%22GBK%22%3F%3E%3CDATA%3E%3CREQUEST%3E%3CNCCUPNTF%3E%3CMSGTYP%3ENCCUPNTF%3C%2FMSGTYP%3E%3CMSGNBR%3E201901100001110856%3C%2FMSGNBR%3E%3CPAYTYP%3E1%3C%2FPAYTYP%3E%3CREQNBR%3E%3C%2FREQNBR%3E%3CMCHNBR%3EM000004331%3C%2FMCHNBR%3E%3CSEQNBR%3E00000000000000006797%3C%2FSEQNBR%3E%3CSUBSEQ%3E0000000001%3C%2FSUBSEQ%3E%3CCCYNBR%3E10%3C%2FCCYNBR%3E%3CTRSAMT%3E000000000000100%3C%2FTRSAMT%3E%3CENDAMT%3E%3C%2FENDAMT%3E%3CBBKNBR%3E%3C%2FBBKNBR%3E%3CPAYACC%3E755915704110704%3C%2FPAYACC%3E%3CACCNAM%3E%C6%F3%D2%B5%CD%F8%D2%F8%D0%C220161342%3C%2FACCNAM%3E%3CRCVACC%3E755915708910605%3C%2FRCVACC%3E%3CRCVNAM%3E%C6%F3%D2%B5%CD%F8%D2%F8%D0%C220161371%3C%2FRCVNAM%3E%3CREFMCH%3EM000004331%3C%2FREFMCH%3E%3CREFORD%3E20190110010317069247%3C%2FREFORD%3E%3CSUBORD%3E%3C%2FSUBORD%3E%3CPAYNBR%3E20190110010317069247%3C%2FPAYNBR%3E%3CYURREF%3E%3C%2FYURREF%3E%3CENDDAT%3E%3C%2FENDDAT%3E%3CRTNFLG%3EP%3C%2FRTNFLG%3E%3CRTNDSP%3E%B6%A9%B5%A5%C7%EB%C7%F3%D2%D1%B7%A2%CB%CD%3C%2FRTNDSP%3E%3CRSV30Z%3EB000000090%3C%2FRSV30Z%3E%3C%2FNCCUPNTF%3E%3C%2FREQUEST%3E%3C%2FDATA%3E&signature=KuUnFREqdJVNk47q1Os0KVZOcIdG6cP4zSt2tNc3xluSu6HHmn9idgreH7Gi6XqiAO6YWo0FqWoLjGWX2QEQc%2FEq3MUrzE4M%2BEnLbFSWWI%2B67l66Bb7xjy5INMfZUq7W9qWbSD2tlX2TJYP1ZjIInM5HhBzEhy%2B1EgWmh0K6Rn0%3D")
		rsp, err := http.Post("http://"+*post, "application/x-www-form-urlencoded", bytes.NewReader(data))
		defer func() {
			if rsp != nil && rsp.Body != nil {
				rsp.Body.Close()
			}
		}()
		if err != nil {
			fmt.Printf("请求错误：%s", err.Error())
		}
		d, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			fmt.Printf("读取返回流错误：%s", err.Error())
			return
		}
		fmt.Printf("返回内容：%s", string(d))
		return
	}
	fmt.Printf("监听端口：%s", *port)
	err := http.ListenAndServe(*port, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("url:%+v\n", request.URL.Path)
		fmt.Printf("head:%+v\n", request.Header)
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte("读取body错误"))
			return
		}

		fmt.Printf("body:%s\n", string(data))
		fmt.Printf("body:%s\n", base64.StdEncoding.EncodeToString(data))
		writer.Write([]byte("ok"))
	}))
	log.Fatal(err)

}
