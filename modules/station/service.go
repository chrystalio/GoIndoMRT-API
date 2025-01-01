package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/chrystalio/GoIndoMRT-API/common/client"
)

const baseUrl = "https://www.jakartamrt.co.id/id/val/stasiuns"

type Service interface {
	GetAllStations() (response []StationResponse, err error)
	CheckSchedulesByStation(id string) (response []ScheduleResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStations() (response []StationResponse, err error) {

	byteResponse, err := client.DoRequest(s.client, baseUrl)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		return
	}

	for _, item := range stations {
		response = append(response, StationResponse(item))
	}

	return
}

func (s *service) CheckSchedulesByStation(id string) (response []ScheduleResponse, err error) {

	byteResponse, err := client.DoRequest(s.client, baseUrl)

	if err != nil {
		return
	}

	var schedules []Schedule
	err = json.Unmarshal(byteResponse, &schedules)

	if err != nil {
		return
	}

	var scheduleSelected Schedule

	for _, item := range schedules {
		if item.StationId == id {
			scheduleSelected = item
			break
		}
	}

	if scheduleSelected.StationId == "" {
		err = errors.New("Station not found")
		return
	}

	response, err = ConvertDataToResponse(scheduleSelected)

	if err != nil {
		return
	}

	return
}

func ConvertDataToResponse(schedule Schedule) (response []ScheduleResponse, err error) {

	var (
		LebakBulusTripName = "Stasiun Lebak Bulus Grab"
		BundaranHITripName = "Stasiun Bundaran HI BANK DKI"
	)

	scheduleLebakBulus := schedule.SechduleLebakBulus
	scheduleBundaranHI := schedule.ScheduleBundaranHI

	scheduleLebakBulusParsed, err := ConvertScheduleToTimeFormat(scheduleLebakBulus)

	if err != nil {
		return
	}

	scheduleBundaranHIParsed, err := ConvertScheduleToTimeFormat(scheduleBundaranHI)

	if err != nil {
		return
	}

	for _, item := range scheduleLebakBulusParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, ScheduleResponse{
				StationName: LebakBulusTripName,
				Time:        item.Format("15:04"),
			})
		}
	}

	for _, item := range scheduleBundaranHIParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, ScheduleResponse{
				StationName: BundaranHITripName,
				Time:        item.Format("15:04"),
			})
		}
	}

	return
}

func ConvertScheduleToTimeFormat(schedule string) (response []time.Time, err error) {
	var schedules = strings.Split(schedule, ",")

	for _, item := range schedules {
		trimmedTime := strings.TrimSpace(item)

		if trimmedTime == "" {
			continue
		}

		parsedTime, parseErr := time.Parse("15:04", trimmedTime)
		if parseErr != nil {
			return nil, errors.New("Invalid time format: " + trimmedTime)
		}

		response = append(response, parsedTime)
	}

	return
}
