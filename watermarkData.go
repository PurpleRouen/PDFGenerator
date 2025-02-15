package main

// You can customize the payloads to match your needs

// RunnerPayload represents the data of the runner
type RunnerPayload struct {
	Gender        string `json:"gender"`
	LastName      string `json:"lastName"`
	FirstName     string `json:"firstName"`
	BirthDate     string `json:"birthDate"`
	RaceLabel     string `json:"raceLabel"`
	RegisteredAt  string `json:"registeredAt"`
	Email         string `json:"email"`
	PPSIdentifier string `json:"ppsIdentifier"`
	CheckoutID    string `json:"checkoutId"`
}

// BuyerPayload represents the data of the buyer
type BuyerPayload struct {
	Gender    string `json:"gender"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}

// InscriptionPayload represents the data of the inscription
type InscriptionPayload struct {
	Runner RunnerPayload `json:"runner"`
	Buyer  BuyerPayload  `json:"buyer"`
}

// GetWatermarksData You can customize the watermarks to match your needs
// GetWatermarksData returns the data to be used as watermarks
func GetWatermarksData(payload *InscriptionPayload) []WatermarkData {
	return []WatermarkData{
		{Text: payload.Runner.Gender + " " + payload.Runner.LastName, X: 23, Y: 86, IsDark: false},
		{Text: payload.Runner.FirstName, X: 23, Y: 93, IsDark: false},
		{Text: payload.Runner.BirthDate, X: 23, Y: 106, IsDark: false},
		{Text: payload.Runner.RaceLabel, X: 44, Y: 128, IsDark: false},
		{Text: payload.Runner.RegisteredAt, X: 123, Y: 30, IsDark: true},
		{Text: payload.Runner.Email, X: 123, Y: 42, IsDark: true},
		{Text: payload.Runner.PPSIdentifier, X: 123, Y: 53, IsDark: true},
		{Text: payload.Runner.CheckoutID, X: 123, Y: 64, IsDark: true},
		{Text: payload.Buyer.Gender + " " + payload.Buyer.LastName + " " + payload.Buyer.FirstName, X: 123, Y: 75, IsDark: true},
		{Text: payload.Buyer.Email, X: 123, Y: 81, IsDark: true},
	}
}
