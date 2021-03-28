package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
)

type Rq struct {
	Status string `json:"status"`
	Progress int `json:"progress"`
	Previews []Links `json:previews"`
}

type Links struct {
	Profile   string `json:"profile"`
	StreamUrl string `json:"stream_url""`
	wav       string `json:"waveform_url""`
}

func put(musicId string, autorization string ) {
	url := "https://api-v2.soundcloud.com/tracks/soundcloud:tracks:"+musicId+"/master-previews?client_id=doesntmatter"
	method := "PUT"
	payload := strings.NewReader(`{"profiles":["I","M","N","G","0","Z"],"preview_interval":{"start_seconds":0,"end_seconds":3600}}`)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:87.0) Gecko/20100101 Firefox/87.0")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", autorization)
	req.Header.Add("Origin", "https://mastering.soundcloud.com")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}


func get(musicId string, autorization string ) int {
	var Rq1 Rq
	url := "https://api-v2.soundcloud.com/tracks/soundcloud:tracks:"+musicId+"/master-previews?client_id=doesntmatter"
	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:87.0) Gecko/20100101 Firefox/87.0")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", autorization)
	req.Header.Add("Origin", "https://mastering.soundcloud.com")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		return 0
	}
	defer res.Body.Close()
	decoder :=json.NewDecoder(res.Body)
	err = decoder.Decode(&Rq1)
	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		return 0
	}
	if Rq1.Progress == 100 {
		links = Rq1.Previews
	}
	return Rq1.Progress
}



func downloadFile (obj Links){
	// Get the data
	resp, err := http.Get(obj.StreamUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(obj.Profile+"Master.mp3")
	if err != nil {
		return
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return
}