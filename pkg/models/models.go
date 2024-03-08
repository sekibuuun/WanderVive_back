package models

type Event struct {
	EventId     int    `json:"eventId"`
	EventName   string `json:"eventName"`
	LivehouseId int    `json:"livehouseId"`
	BandIdList  []int  `json:"bandId"`
	EventDate   string `json:"eventDate"`
	OpenTime    string `json:"openTime"`
	StartTime   string `json:"startTime"`
	Fee         int    `json:"fee"`
}

type EventResponse struct {
	Contents []Event `json:"contents"`
}

type Livehouse struct {
	LivehouseId   int     `json:"livehouseId"`
	LivehouseName string  `json:"livehouseName"`
	Longitude     float64 `json:"longitude"`
	Latitude      float64 `json:"latitude"`
	HomePage      string  `json:"homePage"`
	MapLink       string  `json:"mapLink"`
}

type LivehouseResponse struct {
	Contents []Livehouse `json:"contents"`
}

type Band struct {
	BandId    int    `json:"bandId"`
	BandName  string `json:"bandName"`
	Genre     string `json:"gerne"`
	Youtube   string `json:"youtube"`
	Instagram string `json:"instagram"`
	Twitter   string `json:"twitter"`
	Tunecore  string `json:"tunecore"`
	HomePage  string `json:"homePage"`
	Image     string `json:"image"`
}

type BandResponse struct {
	Contents []Band `json:"contents"`
}
