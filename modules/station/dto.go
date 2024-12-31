package station

type Station struct {
	Id       string `json:"nid"`
	Name     string `json:"title"`
	Order    string `json:"urutan"`
	LocalMap string `json:"peta_lokalitas"`
	Banner   string `json:"banner"`
}

type StationResponse struct {
	Id       string `json:"nid"`
	Name     string `json:"title"`
	Order    string `json:"urutan"`
	LocalMap string `json:"local_map"`
	Banner   string `json:"banner"`
}

type Schedule struct {
	StationId          string `json:"nid"`
	StationName        string `json:"title"`
	ScheduleBundaranHI string `json:"jadwal_hi_biasa"`
	SechduleLebakBulus string `json:"jadwal_lb_biasa"`
}

type ScheduleResponse struct {
	StationName string `json:"station"`
	Time        string `json:"time"`
}
