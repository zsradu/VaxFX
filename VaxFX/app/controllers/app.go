package controllers

import (
	"github.com/revel/revel"
	"github.com/liyue201/goqr"
	"fmt"
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
)

type App struct {
	*revel.Controller
}

func (c App) Submit() revel.Result {
 
	
	imgdata, err := ioutil.ReadFile("public/img/CertificatVaccin-1.png") // TODO: pass as argument the pic uploaded as Green Form
	if err != nil {
		fmt.Printf("%v\n", err)
		return c.RenderTemplate("errors/customError.html")
	}
	img, _, err := image.Decode(bytes.NewReader(imgdata)) // Decodes from byte to image
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
		return c.RenderTemplate("errors/customError.html")
	}
	qrCodes, err := goqr.Recognize(img) // Recognizes if it is a QR image and passes to qrCodes a slice of QRData
	if err != nil {
		fmt.Printf("Recognize failed: %v\n", err)
		return c.RenderTemplate("errors/customError.html")
	}
	fmt.Println(qrCodes)
	// TODO: pass through the Green Pass API to check if it is valid

	// TODO: if it is valid, query the DB to see if it has been uploaded in the past.
	

	
	return c.Render()
}

func (c App) DataSecFX() revel.Result {
	return c.Render()
}

func (c App) DataCovid() revel.Result {
	return c.Render()
}