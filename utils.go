package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	//"net"
)

//var client *http.Client

type nameDropBody struct {
	UNIX  int64  `json:"UNIX"`
	Error string `json:"error"`
}

type nameChangeCheck struct {
	NamechangeAll string `json:"nameChangeAllowed"`
}

func getDropTime(name string) (nameDropBody, error) {
	resp, err := http.Get("https://mojang-api.teun.lol/droptime/" + name)
	if err != nil {
		return nameDropBody{}, errors.New("\n[ERR] failed to send request to teun api")
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nameDropBody{}, errors.New("\n[ERR] cannot read from response body")
	}

	var f nameDropBody
	json.Unmarshal(bodyBytes, &f)

	if f.Error != "" {
		return nameDropBody{}, errors.New(f.Error)
	}

	return f, nil
}

func formatTime(t time.Time) string {
	return t.Format("02:05.99999")
}

func skinChange(bearer string) string {
	postBody, _ := json.Marshal(map[string]string{
		"Content-Type": "application/json",
		"url":          "https://textures.minecraft.net/texture/ddc848a475d34db4bb2b39c5ca6ed585fc4295e6b994170c0d8dddea6bf282e2",
		"variant":      "slim",
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, _ := http.NewRequest("POST", "https://api.minecraftservices.com/minecraft/profile/skins", responseBody)

	resp.Header.Set("Authorization", "bearer "+bearer)

	skin, _ := http.DefaultClient.Do(resp)

	if skin.StatusCode == 200 {
		fmt.Println("\n[+] Succesfully Changed your skin:", skin.StatusCode)
	} else if skin.StatusCode != 200 {
		fmt.Println("\n[INFO] failed...")
	}
	return "a"
}

func Speed(conn *tls.Conn, bearer string) []byte {
	//heh := make([]byte, 4028)

	e := make([]byte, 4028)
	n, _ := conn.Read(e)

	//e = append(e[9:12], heh[0])

	fmt.Println("[INFO] Status Code:", string(e[9:12]))

	if string(e[9:12]) == `200` {
		go sendWebHook(config["webhook_url"].(string), config["discord_ID"].(string), name, dropDelay)
		go skinChange(bearer)
	}

	for i := 0; i == -1; i++ {
		if i == 10 {
			fmt.Println(n)
		}
	}

	return e
}

func checkChange(bearer string) string {

	check, _ := http.Get("https://api.minecraftservices.com/minecraft/profile/namechange")

	check.Header.Add("Authorization", "Bearer "+bearer)
	check.Header.Add("Content-Type", "application/json")

	defer check.Body.Close()

	fmt.Println("HTTP Response Status:", check.StatusCode, http.StatusText(check.StatusCode))

	body, _ := ioutil.ReadAll(check.Body)

	var g nameChangeCheck
	json.Unmarshal(body, &g)

	return g.NamechangeAll

}

func sendWebHook(wh string, id string, name string, dropDelay float64) {
	if wh == "" {
		fmt.Println("You do not have any webhooks!")
		os.Exit(0)
	}

	webhookINFO := fmt.Sprintf(`{"username": "Medusa", "avatar_url": "https://cdn.discordapp.com/attachments/834840617901096990/855873132577423390/a.png", "embeds": [{"title": ":snake: **__New Smite!__** :snake:", "color": "14177041", "fields": [{"name": "User:", "value": "<@%v>", "inline": "false"},{"name": "Name Sniped: :smoking:", "value": "%v", "inline": "false"},{"name": "Delay used: :cloud_lightning:", "value": "%v", "inline": "false"},{"name": "Discord", "value": "https://discord.gg/WtQ2d7NNQ4", "inline": "false"}]}]}`, id, name, dropDelay)
	newRequest, _ := http.NewRequest("POST", wh, bytes.NewReader([]byte(webhookINFO)))
	newRequest.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(newRequest)
	if err != nil {
		fmt.Println("[ERR] Failure sending webhook!!!")
	}

	time.Sleep(2 * time.Second)
	fmt.Println("[INFO] Sent Webhook!", resp.StatusCode)

}
