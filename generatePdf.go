package main

import (
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"image/png"
	"os"
	"os/exec"
	"strconv"
)

type WatermarkData struct {
	Text   string
	IsDark bool
	X      int
	Y      int
}

const descWhite = "rot:0, scale:0.7 abs, pos:tl, fillc:#ffffff, strokec:#ffffff, "
const descDark = "rot:0, scale:0.6 abs, pos:tl, fillc:#000000, strokec:#000000, "

func textAt(x, y int, text string, dark bool) (*model.Watermark, error) {
	x = int(float64(x) * 2.85)
	y = int(float64(y) * 2.85)
	desc := descWhite
	if dark {
		desc = descDark
	}
	return pdfcpu.TextWatermark(text, desc+"off:"+strconv.Itoa(x)+" -"+strconv.Itoa(y), true, false, types.POINTS)
}

func generatePdf(payload *InscriptionPayload) (*string, error) {
	inscriptionPath := "inscriptions/" + payload.Runner.CheckoutID + ".pdf"
	err := exec.Command("cp", "template.pdf", inscriptionPath).Run()
	if err != nil {
		return nil, err
	}

	var watermarks map[int][]*model.Watermark
	watermarks = make(map[int][]*model.Watermark)
	for _, wmData := range GetWatermarksData(payload) {
		var wm *model.Watermark
		wm, err = textAt(wmData.X, wmData.Y, wmData.Text, wmData.IsDark)
		if err != nil {
			return nil, err
		}
		watermarks[1] = append(watermarks[1], wm)
	}

	err = pdfcpu.AddWatermarksSliceMapFile(inscriptionPath, inscriptionPath, watermarks, nil)
	if err != nil {
		return nil, err
	}

	enc := qrcode.NewQRCodeWriter()
	img, err := enc.Encode(payload.Runner.CheckoutID, gozxing.BarcodeFormat_QR_CODE, 256, 256, nil)
	if err != nil {
		return nil, err
	}

	qrcodePath := "qrcodes/" + payload.Runner.CheckoutID + ".png"
	file, err := os.Create(qrcodePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return nil, err
	}

	pages := []string{"1"}
	err = pdfcpu.AddImageWatermarksFile(inscriptionPath, inscriptionPath, pages, true, qrcodePath, "rot:0, scale:0.6 abs, pos:tl, off:70 -50", nil)
	if err != nil {
		return nil, err
	}

	return &inscriptionPath, nil
}
