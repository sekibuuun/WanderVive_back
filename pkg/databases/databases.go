package databases

import (
	"WanderVive_back/pkg/models"
	"database/sql"
	"log"
)

type Service struct {
	DB *sql.DB
}

func (svc *Service) GetBand() []models.Band {
	rows, err := svc.DB.Query("SELECT * FROM band;")
	if err != nil {
		log.Fatal(err)
	}
	bands := make([]models.Band, 0)
	for rows.Next() {
		var band models.Band
		if err := rows.Scan(&band.BandId, &band.BandName, &band.Genre, &band.Youtube, &band.Twitter, &band.Instagram, &band.Tunecore, &band.HomePage, &band.Image); err != nil {
			log.Fatal(err)
		}
		bands = append(bands, band)
	}
	return bands
}

func (svc *Service) GetLivehouse() []models.Livehouse {
	rows, err := svc.DB.Query("SELECT * FROM livehouse;")
	if err != nil {
		log.Fatal(err)
	}
	livehouses := make([]models.Livehouse, 0)
	for rows.Next() {
		var livehouse models.Livehouse
		if err := rows.Scan(&livehouse.LivehouseId, &livehouse.LivehouseName, &livehouse.Latitude, &livehouse.Longitude, &livehouse.HomePage, &livehouse.MapLink); err != nil {
			log.Fatal(err)
		}
		livehouses = append(livehouses, livehouse)
	}
	return livehouses
}

func (svc *Service) GetEvent() []models.Event {
	rows, err := svc.DB.Query("SELECT * FROM event;")
	if err != nil {
		log.Fatal(err)
	}
	events := make([]models.Event, 0)
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.EventId, &event.EventName, &event.LivehouseId, &event.EventDate, &event.OpenTime, &event.StartTime, &event.Fee); err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}
	paticipants := svc.getPaticipants(len(events))
	for i, e := range paticipants {
		events[i].BandIdList = e
	}
	return events
}

func (svc *Service) getPaticipants(eventNum int) [][]int {
	paticipants := make([][]int, eventNum)
	rows, err := svc.DB.Query("SELECT eventId, bandId FROM paticipant;")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var eventId int
		var bandId int
		if err := rows.Scan(&eventId, &bandId); err != nil {
			log.Fatal(err)
		}
		paticipants[eventId-1] = append(paticipants[eventId-1], bandId)
	}
	return paticipants
}

func (svc *Service) GetEventAndLivehouse() []models.EventAndLivehouse {
	rows, err := svc.DB.Query("SELECT * FROM event NATURAL INNER JOIN livehouse;")
	if err != nil {
		log.Fatal(err)
	}
	eventAndLivehouseList := make([]models.EventAndLivehouse, 0)
	for rows.Next() {
		var eventAndLivehouse models.EventAndLivehouse
		if err := rows.Scan(&eventAndLivehouse.LivehouseId, &eventAndLivehouse.EventId, &eventAndLivehouse.EventName, &eventAndLivehouse.EventDate, &eventAndLivehouse.StartTime, &eventAndLivehouse.OpenTime, &eventAndLivehouse.Fee, &eventAndLivehouse.LivehouseName, &eventAndLivehouse.Longitude, &eventAndLivehouse.Latitude, &eventAndLivehouse.HomePage, &eventAndLivehouse.MapLink); err != nil {
			log.Fatal(err)
		}
		eventAndLivehouseList = append(eventAndLivehouseList, eventAndLivehouse)
	}
	paticipants := svc.getPaticipants(len(eventAndLivehouseList))
	for i, eventPaticipant := range paticipants {
		var bandList []models.Band
		for _, paticipant := range eventPaticipant {
			bandList = append(bandList, svc.GetCertainBand(paticipant))
		}
		eventAndLivehouseList[i].BandList = bandList
	}

	return eventAndLivehouseList
}

func (svc *Service) GetCertainBand(id int) models.Band {
	log.Println(id)
	row := svc.DB.QueryRow("SELECT * FROM band WHERE bandId = $1", id)
	var band models.Band
	if err := row.Scan(&band.BandId, &band.BandName, &band.Genre, &band.Youtube, &band.Twitter, &band.Instagram, &band.Tunecore, &band.HomePage, &band.Image); err != nil {
		log.Fatal(err)
	}
	return band
}
