package databases

import (
	"WanderVive_back/pkg/models"
)

func GetBand() models.Band {
	test_band := models.Band{
		BandId:    1,
		BandName:  "サンプルバンド",
		Genre:     "ロック",
		Youtube:   "https://www.youtube.com/channel/UCxxxxxxxx",
		Instagram: "https://www.instagram.com/sampleband",
		Twitter:   "https://twitter.com/sampleband",
		Tunecore:  "https://www.tunecore.com/artists/sampleband",
		HomePage:  "https://www.sampleband.com",
		Image:     "../../images/sampleband.png",
	}
	return test_band
}

func GetLivehouse() models.Livehouse {
	test_livehouse := models.Livehouse{
		LivehouseId:   1,
		LivehouseName: "サンプルライブハウス",
		Longitude:     123.456,
		Latitude:      78.90,
		HomePage:      "https://www.samplelivehouse.com",
		MapLink:       "https://www.google.com/maps/place/サンプルライブハウス",
	}
	return test_livehouse
}

func GetEvent() models.Event {
	test_event := models.Event{
		EventId:     1,
		EventName:   "サンプルイベント",
		LivehouseId: 1,
		BandId:      []int{1},
		EventDate:   "2022-01-01",
		OpenTime:    "12:00",
		StartTime:   "13:00",
		Fee:         2000,
	}
	return test_event
}
