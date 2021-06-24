package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
type loginResponse struct {
	Token string `json:"accessToken"`
}
*/

type SecurityRes struct {
	Answer AnswerRes `json:"answer"`
}

type AnswerRes struct {
	ID int `json:"id"`
}

func mojangLogin1() string {

	conn, _ := tls.Dial("tcp", "authserver.mojang.com:443", nil)

	e, err := ioutil.ReadFile("accounts.txt")
	if err != nil {
		panic(err)
	}

	test := strings.Split(string(e), ":")

	payload := "{\"username\": \"" + test[0] + "\", \"password\": \"" + test[1] + "\", \"agent\": {\"name\": \"Minecraft\", \"version\": 1}}"

	data := "POST /authenticate HTTP/1.1\r\nContent-Type: application/json\r\nHost: authserver.mojang.com\r\nUser-Agent: AthenaGO/1/1/unknown\r\nContent-Length: " + strconv.Itoa(len(payload)) + "\r\n\r\n" + payload

	var authbytes []byte
	authbytes = make([]byte, 4096)

	auth := make(map[string]interface{})

	var security []SecurityRes
	conn.Write([]byte(data))
	conn.Read(authbytes)
	conn.Close()

	authbytes = []byte(strings.Split(strings.Split(string(authbytes), "\x00")[0], "\r\n\r\n")[1])
	err = json.Unmarshal(authbytes, &auth)

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.mojang.com/user/security/challenges", nil)
	if err != nil {
		fmt.Println("[ERR] An error accured:", err)
	}

	if auth["accessToken"] == nil {
		fmt.Println("[ERR] Bearer token is empty!")
		os.Exit(0)
	}

	req.Header.Set("Authorization", "Bearer "+auth["accessToken"].(string))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	securitybytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(securitybytes, &security)
	if err != nil {
		fmt.Println(err)
	}

	data = `[{"id": ` + strconv.Itoa(security[0].Answer.ID) + `, "answer": "` + test[2] + `"}, {"id": ` + strconv.Itoa(security[1].Answer.ID) + `, "answer": "` + test[3] + `"}, {"id": ` + strconv.Itoa(security[2].Answer.ID) + `, "answer": "` + test[4] + `"}]`

	b := bytes.NewReader([]byte(data))
	req, err = http.NewRequest("GET", "https://api.mojang.com/user/security/location", b)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", "Bearer "+auth["accessToken"].(string))
	_, _ = client.Do(req)

	return (auth["accessToken"].(string))

}

type loginResponse struct {
	Token string `json:"accessToken"`
}

func mojangLogin2(email, password string) string {
	postBody, _ := json.Marshal(map[string]string{
		"username": email,
		"password": password,
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://authserver.mojang.com/authenticate", "application/json", responseBody)
	if err != nil {
		log.Fatalf("[ERR] An error has occured: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("[ERR] an error has occured", err)
	}

	var f loginResponse
	json.Unmarshal(body, &f)

	return f.Token
}
