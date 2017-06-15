package route

type Route struct {
	UID      string      `json:"uid"`
	Name     string      `json:"name"`
	Color    string      `json:"color"`
	Duration int         `json:"duration"`
	Points   [][]float32 `json:"points"`
}
