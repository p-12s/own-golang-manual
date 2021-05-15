package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	data := []byte(`
	{
	  "ok": true,
	  "total": 3,
	  "documents": [
	    {"id":11, "title": "c++"},
	    {"id":2,  "title": "suddenly perl"},
	    {"id":5,  "title": "go"}
	  ]
	}
	`)
	var resp struct {
		Ok        bool `json:"ok"`
		Total     int  `json:"total"`
		Documents []struct{
			Id    int    `json:"id"`
			Title string `json:"title"`
		} `json:"documents"`
	}
	err := json.Unmarshal(data, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Documents[2].Title)
	fmt.Printf("%v\n", resp)
	fmt.Printf("%T\n", resp)
}
