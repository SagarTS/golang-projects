package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name 	 string `json:"course"`
	Price 	 int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags 	 []string`json:"tags,omitempty"`
}

func main(){
	fmt.Println("Welcome to JSON.")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	courses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in","abc123",[]string{"web-dev","js"}},
		{"MERN Bootcamp", 499, "LearnCodeOnline.in","xyz123",[]string{"web-dev","js"}},
		{"Angular Bootcamp", 199, "LearnCodeOnline.in","mno123",[]string{}},
	}

	//package this data as json data

	finalJson, err := json.MarshalIndent(courses,"","\t")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
		{
		"course": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
		}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid.")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}else {
		fmt.Println("JSON WAS NOT VALID.")
	}

	// json to key-value pair

	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	for k,v := range onlineData{
		fmt.Printf("Key is %v and value is %v and Type is %T\n",k,v,v)
	}
}