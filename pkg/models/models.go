package models

type Event struct {
	eventId     int    `json:"eventId"`
	eventName   string `json:"eventName"`
	livehouseId int    `json:"livehouseId"`
	bandId      []int  `json:"bandId"`
	eventDate   string `json:"eventDate"`
	openTime    string `json:"openTime"`
	startTime   string `json:"startTime"`
	fee         int    `json:"fee"`
}

type Livehouse struct {
	livehouseId   int     `json:"livehouseId"`
	livehouseName string  `json:"livehouseName"`
	longitude     float64 `json:"longitude"`
	latitude      float64 `json:"latitude"`
	homePage      string  `json:"homePage"`
	mapLink       string  `json:"mapLink"`
}

type Band struct {
	bandId    int    `json:"bandId"`
	bandName  string `json:"bandName"`
	gerne     string `json:"gerne"`
	youtube   string `json:"youtube"`
	instagram string `json:"instagram"`
	twitter   string `json:"twitter"`
	tunecore  string `json:"tunecore"`
	homePage  string `json:"homePage"`
	image     string `json:"image"`
}
