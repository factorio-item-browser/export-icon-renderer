package transfer

type Color struct {
	Red   float64 `json:"red"`
	Green float64 `json:"green"`
	Blue  float64 `json:"blue"`
	Alpha float64 `json:"alpha"`
}

type Icon struct {
	Id     string  `json:"id"`
	Size   int     `json:"size"`
	Layers []Layer `json:"layers"`
}

type Layer struct {
	FileName string  `json:"fileName"`
	Offset   Offset  `json:"offset"`
	Scale    float64 `json:"scale"`
	Size     int     `json:"size"`
	Tint     Color   `json:"tint"`
}

type Offset struct {
	X int `json:"x"`
	Y int `json:"y"`
}
