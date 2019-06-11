package main

import (
	"encoding/json"
	"os"
)

type config struct {
	SourceImageDirectory string //directory containing wallpaper images
	// SourceImageName  string  //name of source file
	OutputImageDirectory string
	MarginRight          float64 //margin from right to start appointments text
	MarginTop            float64 //margin from top to start appointment title
	PrintDate            bool
	OutputFileName       string  //name of the target file
	TitleFont            string  //controls appointment items font
	TitleFontSize        float64 //controls appointment items font size
	TitleText            string  //Calendar Title text
	ItemFont             string  //controls appointment items font
	ItemFontSize         float64 //controls appointment items font size
	ItemPadding          float64 //controls spacing between appointment items
	GoogleCalendarID     string
}

func main() {
	cfg := loadConfig()

	// Test Data
	appointments := getAppointmentsStub()
	// appointments := getGoogleAppointments(cfg.GoogleCalendarID)
	writeAppointmentsToImage(appointments, cfg)
}

func loadConfig() *config {
	configFileName := "config.json"

	configFile, err := os.Open(configFileName)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	cfg := &config{}
	err = json.NewDecoder(configFile).Decode(cfg)
	if err != nil {
		panic("parsing config: " + err.Error())
	}
	return cfg
}

// Test Data
func getAppointmentsStub() []string {
	return []string{
		"Call Sunny (2019-06-10)",
		"Parents Teacher Meet (2019-06-10)",
		"Pay Electricity Bill (2019-06-10)",
		"John's Birthday Reminder (2019-06-12)",
		"Pay Broadband Bill (2019-06-13)",
		"Pay SGSAO Maint. Charges (2019-06-14)",
	}
}
