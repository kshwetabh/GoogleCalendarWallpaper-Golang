package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	SourceImageName  string  //name of source file
	MarginRight      float64 //margin from right to start appointments text
	MarginTop        float64 //margin from top to start appointment title
	Width            int     //width and height of the wallpaper image
	Height           int
	PrintDate        bool
	OutputFileName   string  //name of the target file
	TitleFont        string  //controls appointment items font
	TitleFontSize    float64 //controls appointment items font size
	TitleText        string  //Calendar Title text
	ItemFont         string  //controls appointment items font
	ItemFontSize     float64 //controls appointment items font size
	ItemPadding      float64 //controls spacing between appointment items
	GoogleCalendarID string
}

func main() {
	// cfg := config{
	// 	sourceImageName:  "canyon.jpg",
	// 	outputFileName:   "wallpaper.png",
	// 	marginRight:      300,
	// 	marginTop:        50,
	// 	width:            2560,
	// 	height:           1440,
	// 	printDate:        true,
	// 	titleFont:        "./Fonts/Roboto/Roboto-Bold.ttf",
	// 	titleFontSize:    18,
	// 	itemFontSize:     14,
	// 	itemFont:         "./Fonts/Roboto/Roboto-Medium.ttf",
	// 	titleText:        "Today's Calendar",
	// 	itemPadding:      25,
	// 	googleCalendarID: "lts59adca39hk04r3vnlvir93s@group.calendar.google.com",
	// }

	cfg := loadConfig()
	// appointments := []string{
	// 	"Call Sunny",
	// 	"Parents Teacher Meet",
	// 	"Pay LIC Premium",
	// 	"Gautam's Happy Birthday",
	// 	"Pay BESCOM Bill",
	// 	"Pay SGSAO Maint. Charges",
	// }
	fmt.Println(cfg)
	appointments := getGoogleAppointments(cfg)
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
