/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   Kemal CAN SİLER
   19.07.2019
   13:56
*/
package gonesignal

import (
	"fmt"
	"net/http"
	"net/url"
)

// PlayersService handles communication with the player related
// methods of the OneSignal API.
type PlayersService struct {
	client *Client
}

/*
	- olanlar dökünmanda bulunamadı kontrol edilecek
	+ olanlar dökümanda var
	New olanlar yeni eklendi
*/

// Player represents a OneSignal player.
type Player struct {
	ID                string            `json:"id"`                 // -
	Playtime          int               `json:"playtime"`           // -
	SDK               string            `json:"sdk"`                // -
	Identifier        string            `json:"identifier"`         // +
	SessionCount      int               `json:"session_count"`      // +
	Language          string            `json:"language"`           // +
	Timezone          int               `json:"timezone"`           // +
	GameVersion       string            `json:"game_version"`       // +
	DeviceOS          string            `json:"device_os"`          // +
	DeviceType        int               `json:"device_type"`        // +
	DeviceModel       string            `json:"device_model"`       // +
	AdID              string            `json:"ad_id"`              // +
	Tags              map[string]string `json:"tags"`               // +
	LastActive        int               `json:"last_active"`        // +
	AmountSpent       float32           `json:"amount_spent"`       // +
	CreatedAt         int               `json:"created_at"`         // +
	InvalidIdentifier bool              `json:"invalid_identifier"` // +
	BadgeCount        int               `json:"badge_count"`        // +
}

// PlayerRequest represents a request to create/update a player.
type PlayerRequest struct {
	AppID             string            `json:"app_id"`                       // + REQUIRED
	DeviceType        int               `json:"device_type"`                  // -
	Identifier        string            `json:"identifier,omitempty"`         // +
	Language          string            `json:"language,omitempty"`           // +
	Timezone          int               `json:"timezone,omitempty"`           // +
	GameVersion       string            `json:"game_version,omitempty"`       // +
	DeviceOS          string            `json:"device_os,omitempty"`          // +
	DeviceModel       string            `json:"device_model,omitempty"`       // +
	AdID              string            `json:"ad_id,omitempty"`              // +
	SDK               string            `json:"sdk,omitempty"`                // +
	SessionCount      int               `json:"session_count,omitempty"`      // +
	Tags              map[string]string `json:"tags,omitempty"`               // +
	AmountSpent       string            `json:"amount_spent,omitempty"`       // + update float32 to string
	CreatedAt         int               `json:"created_at,omitempty"`         // +
	Playtime          int               `json:"playtime,omitempty"`           // +
	BadgeCount        int               `json:"badge_count,omitempty"`        // +
	LastActive        int               `json:"last_active,omitempty"`        // +
	TestType          int               `json:"test_type,omitempty"`          // +
	NotificationTypes string            `json:"notification_types,omitempty"` // +
	// New
	Long    float32 `json:"long,omitempty"`
	Lat     float32 `json:"lat,omitempty"`
	Country float32 `json:"country,omitempty"`
}

// PlayerListOptions specifies the parameters to the PlayersService.List method
type PlayerListOptions struct {
	AppID  string `json:"app_id"` // +
	Limit  string `json:"limit"`  // + update int to string
	Offset string `json:"offset"` // +update int to string
}

// PlayerListResponse wraps the standard http.Response for the
// PlayersService.List method
type PlayerListResponse struct {
	TotalCount int      `json:"total_count"` // +
	Offset     int      `json:"offset"`      // +
	Limit      int      `json:"limit"`       // +
	Players    []Player // +
}

// PlayerCreateResponse wraps the standard http.Response for the
// PlayersService.Create method
type PlayerCreateResponse struct {
	Success bool   `json:"success"` // +
	ID      string `json:"id"`      // +
}

// PlayerOnSessionOptions specifies the parameters to the
// PlayersService.OnSession method
type PlayerOnSessionOptions struct {
	Identifier  string            `json:"identifier,omitempty"`   // +
	Language    string            `json:"language,omitempty"`     // +
	Timezone    int               `json:"timezone,omitempty"`     // +
	GameVersion string            `json:"game_version,omitempty"` // +
	DeviceOS    string            `json:"device_os,omitempty"`    // +
	AdID        string            `json:"ad_id,omitempty"`        // +
	SDK         string            `json:"sdk,omitempty"`          // +
	Tags        map[string]string `json:"tags,omitempty"`         // +
}

// PlayerOnPurchaseOptions specifies the parameters to the
// PlayersService.OnPurchase method
type PlayerOnPurchaseOptions struct {
	Purchases []Purchase `json:"purchases"`          // + REQUIRED
	Existing  bool       `json:"existing,omitempty"` // +
}

// Purchase array
type Purchase struct {
	SKU    string  `json:"sku"`    // + REQUIRED
	Amount float32 `json:"amount"` // + REQUIRED
	ISO    string  `json:"iso"`    // + REQUIRED
}

// PlayerOnFocusOptions specifies the parameters to the
// PlayersService.OnFocus method
type PlayerOnFocusOptions struct {
	State      string `json:"state"`       // + REQUIRED
	ActiveTime int    `json:"active_time"` // + REQUIRED
}

// PlayerCSVExportOptions specifies the parameters to the
// PlayersService.CSVExport method
type PlayerCSVExportOptions struct {
	AppID string `json:"app_id"` // + REQUIRED
}

// PlayerCSVExportResponse Response
type PlayerCSVExportResponse struct {
	CSVFileURL string `json:"csv_file_url"` // +
}

// Get a single player.
//
// OneSignal	 API docs: https://documentation.onesignal.com/docs/playersid
// GET (https://onesignal.com/api/v1/players/:id)
func (s *PlayersService) Get(playerID string) (*Player, *http.Response, error) {
	// build the URL
	path := fmt.Sprintf("/players/%s", playerID)
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	plResp := new(Player)
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}
	plResp.ID = playerID

	return plResp, resp, err
}

// Update a player.
//
// OneSignal API docs: https://documentation.onesignal.com/docs/playersid-1
// PUT (https://onesignal.com/api/v1/players/:id)
func (s *PlayersService) Update(playerID string, player *PlayerRequest) (*SuccessResponse, *http.Response, error) {
	// build the URL
	path := fmt.Sprintf("/players/%s", playerID)
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	// create the request
	req, err := s.client.NewRequest("PUT", u.String(), player)
	if err != nil {
		return nil, nil, err
	}

	plResp := &SuccessResponse{}
	resp, err := s.client.Do(req, plResp)
	if err != nil {
		return nil, resp, err
	}

	return plResp, resp, err
}
