package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func httpSpeed(resp *http.Response, sentTime time.Time) {
	fmt.Println("[+]", resp.StatusCode, "Sent Request @:", formatTime(sentTime))
}

func sendMojangRequestsGC(name, bearerGC string) {
	var js = []byte(`{"profileName":"` + name + `"}`)
	req, _ := http.NewRequest("POST", "https://api.minecraftservices.com/minecraft/profile",
		bytes.NewBuffer(js))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+bearerGC)
	for i := 0; i != 6; i++ {

		resp, _ := http.DefaultClient.Do(req)

		sentTime := time.Now()

		go httpSpeed(resp, sentTime)

	}
}

func socketSendingMS(url1 string, bearerMS string, name string) {
	oo, _ := url.Parse(url1)
	conn, _ := tls.Dial("tcp", oo.Hostname()+":443", nil)

	for i := 0; i != 2; i++ {
		time2 := time.Now()
		fmt.Fprintln(conn, "PUT /minecraft/profile/name/"+name+" HTTP/1.1\r\nHost: api.minecraftservices.com\r\nUser-Agent: Medusa/1.0\r\nAuthorization: bearer "+bearerMS+"\r\n\r\n")
		time1 := time.Now()
		go Speed(conn, bearerMS)

		fmt.Println("[INFO] Sent:", formatTime(time1), "MS", "\n[INFO] Received", formatTime(time2), "MS")

		if i == 2 {
			break
		}
	}
	fmt.Print("\n")

}

func socketSending(bearer string, name string) {
	conn, _ := tls.Dial("tcp", "api.minecraftservices.com"+":443", nil)

	for i := 0; i != 2; i++ {
		time2 := time.Now()
		fmt.Fprintln(conn, "PUT /minecraft/profile/name/"+name+" HTTP/1.1\r\nHost: api.minecraftservices.com\r\nUser-Agent: Medusa/1.0\r\nAuthorization: bearer "+bearer+"\r\n\r\n")
		time1 := time.Now()
		fmt.Println("[INFO] Sent:", formatTime(time1), "MS", "\n[INFO] Received", formatTime(time2), "MS")
		go Speed(conn, bearer)

		if i == 2 {
			break
		}

	}
	fmt.Print("\n")

}

func testingGC(name string, bearerGC string) {

	conn, _ := tls.Dial("tcp", "api.minecraftservices.com"+":443", nil)
	var js = []byte(`{"profileName":"` + name + `"}`)
	length := strconv.Itoa(len(string(js)))

	for i := 0; i != 6; i++ {
		time2 := time.Now()
		payload := "POST /minecraft/profile HTTP/1.1\r\nHost: api.minecraftservices.com\r\nConnection: close\r\nContent-Length:" + length + "\r\nContent-Type: application/json\r\nAccept: application/json\r\nAuthorization: Bearer " + bearerGC + "\r\n\r\n" + string(js) + "\r\n"
		fmt.Fprint(conn, payload)
		time1 := time.Now()
		go Speed(conn, bearerGC)

		fmt.Println("[INFO] Sent:", formatTime(time1), "MS", "\n[INFO] Received", formatTime(time2), "MS")
		if i == 6 {
			break

		}

	}

	fmt.Print("\n")

}
