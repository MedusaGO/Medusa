package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func securityQuestionsMFA(bearer string) string {

	//e, err := ioutil.ReadFile("accounts.txt")
	//if err != nil {
	//	panic(err)
	//}

	//test := strings.Split(string(e), ":")

	//sec1 := test[2]
	//sec2 := test[3]
	//sec3 := test[4]

	req, err := http.NewRequest("GET", "https://api.mojang.com/user/security/challenges", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", "bearer "+bearer)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	return string(body)
}
