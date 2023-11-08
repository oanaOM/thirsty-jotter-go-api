package models

// Plant represent the data structure of a plant record
type Plant struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	Type          string `json:"type,omitempty"`
	Image         string `json:"image,omitempty"`
	CreatedAtDate string `json:"created_at_date,omitempty"`
	Quantity      *int   `json:"quantity,omitempty"`
}

// ID function returns the id of the plant
func (p Plant) ID() string {
	return p.Id
}
