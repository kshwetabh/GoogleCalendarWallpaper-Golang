package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"
)

func writeAppointmentsToImage(appointments []string, cfg *config) {
	log.Println("Overlaying the appointment list on the wallpaper...")

	var files []string
	if len(cfg.SourceImageDirectory) > 0 {
		files = getImageList(cfg.SourceImageDirectory)
	}
	// log.Println("got image list", files)

	for _, filename := range files {
		writeToImage(filename, appointments, cfg)
	}
	log.Println("Wallpaper updated with the appointments...")
}

func getImageList(root string) []string {
	var files []string

	// log.Printf("Loading wallpapers from %s\n", root)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func writeToImage(filename string, appointments []string, cfg *config) {
	// log.Printf("Loading image %s\n", filename)
	im, err := gg.LoadImage(filepath.Join(cfg.SourceImageDirectory, filename))
	if err != nil {
		log.Fatal("error loading image", err)
	}

	height := im.Bounds().Dy()
	width := im.Bounds().Dx()

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(1, 1, 1)
	if err := dc.LoadFontFace(cfg.TitleFont, cfg.TitleFontSize); err != nil {
		panic(err)
	}
	dc.DrawImage(im, 0, 0)

	//Draw semi-transparent background to display appointments
	grad := gg.NewRadialGradient(100, 100, 10, 100, 120, 80)
	grad.AddColorStop(0, color.RGBA{50, 50, 50, 150})
	grad.AddColorStop(1, color.RGBA{50, 50, 50, 150})
	dc.SetFillStyle(grad)
	dc.DrawRoundedRectangle(float64(width)-cfg.MarginRight-20, cfg.MarginTop-30, cfg.MarginRight, 512, 5)
	dc.Fill()

	//Draw Title Text
	titleText := cfg.TitleText
	if cfg.PrintDate == true {
		titleText += fmt.Sprintf(" ( %s )", time.Now().Format("Mon, 02-Jan"))
	}
	dc.DrawStringAnchored(titleText, float64(width)-cfg.MarginRight, cfg.MarginTop, 0, 0)

	//Underline Heading
	dc.SetLineWidth(2)
	dc.DrawLine(float64(width)-cfg.MarginRight, cfg.MarginTop+7, float64(width)-cfg.MarginTop, cfg.MarginTop+7)
	dc.DrawLine(float64(width)-cfg.MarginRight, cfg.MarginTop+10, float64(width)-cfg.MarginTop, cfg.MarginTop+10)
	dc.Stroke()

	// Load smaller and lighter font
	if err := dc.LoadFontFace(cfg.ItemFont, cfg.ItemFontSize); err != nil {
		panic(err)
	}

	//Draw Appointment Items
	i := cfg.MarginTop + 10
	for index, appt := range appointments {
		i += cfg.ItemPadding
		dc.DrawStringAnchored(fmt.Sprintf("%d. %s", index+1, appt), float64(width)-cfg.MarginRight, i, 0, 0)
	}
	dc.Clip()
	log.Printf("Writing images to: %s\n", filepath.Join(cfg.OutputImageDirectory, filename))
	dc.SavePNG(filepath.Join(cfg.OutputImageDirectory, filename))

}
