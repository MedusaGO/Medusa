package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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

/*
func askForInput(input string) string {
	fmt.Println(input)

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
*/

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

type nameChangeCheck struct {
	NamechangeAll string `json:"nameChangeAllowed"`
}

func checkChange(bearer string) string {

	check, _ := http.Get("https://api.minecraftservices.com/minecraft/profile/namechange")

	check.Header.Set(
		"Authorization: bearer", bearer)
	defer check.Body.Close()
	fmt.Println("HTTP Response Status:", check.StatusCode, http.StatusText(check.StatusCode))

	body, err := ioutil.ReadAll(check.Body)
	if err != nil {
		log.Fatalln("[ERR] an error has occured", err)
	}

	var g nameChangeCheck
	json.Unmarshal(body, &g)

	return g.NamechangeAll

}

func checkBearer(bearer string) string {
	req, _ := http.NewRequest("GET", "https://api.minecraftservices.com/minecraft/profile", nil)
	req.Header.Set("Authorization", "Bearer "+bearer)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("[ERR] Retrying...")
		checkBearer(bearer)
	}
	resp.Body.Close()

	switch resp.StatusCode {
	case 401:
		panic("bearer is unauthorized")
	case 200:
		return "NC"
	case 404:
		return "GC"
	default:
		panic("unknown bearer type")
	}
}

func sendWebHook(wh string, id string, name string, dropDelay float64) {
	if wh == "" {
		fmt.Println("You do not have any webhooks!")
		os.Exit(0)
	}

	webhookINFO := fmt.Sprintf(`{"username": "AthenaGO", "avatar_url": "https://cdn.discordapp.com/attachments/834840617901096990/855873132577423390/a.png", "embeds": [{"title": ":trident: **__New Smite!__** :trident:", "color": "14177041", "image": {"url": "https://cdn.discordapp.com/attachments/834840617901096990/855922075629518869/concours-discord-cartes-voeux-fortnite-france-6.png"}, "fields": [{"name": "User:", "value": "<@%v>", "inline": "false"},{"name": "Name Sniped: :smoking:", "value": "%v", "inline": "false"},{"name": "Delay used: :cloud_lightning:", "value": "%v", "inline": "false"},{"name": "Discord", "value": "https://discord.gg/WtQ2d7NNQ4", "inline": "false"}]}]}`, id, name, dropDelay)
	newRequest, _ := http.NewRequest("POST", wh, bytes.NewReader([]byte(webhookINFO)))
	newRequest.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(newRequest)
	if err != nil {
		fmt.Println("[ERR] Failure sending webhook!!!")
	}

	time.Sleep(2 * time.Second)
	fmt.Println("[INFO] Sent Webhook!", resp.StatusCode)

}
