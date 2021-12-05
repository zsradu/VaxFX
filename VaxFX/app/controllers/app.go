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
	"VaxFX/app/models"
	"VaxFX/app"
)

type App struct {
	*revel.Controller
}

func (c App) checkForQR(path string) revel.Result {
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

	return c.RenderTemplate("App/submit.html")

	// TODO: pass through the Green Pass API to check if it is valid

	// TODO: if it is valid, query the DB to see if it has been uploaded in the past.
}

func (c App) Submit() revel.Result {
	var form models.Form
	var effectNames1 [4]string
	var effectNames2 [4]string
	var effectNames3 [4]string
	dataNames := []string{"Durere de cap", "Febra", "Ameteala", "Greata"}
	c.Params.Bind(&form, "Form")
	c.Params.Bind(&effectNames1[0], "E1Name1")
	c.Params.Bind(&effectNames1[1], "E1Name2")
	c.Params.Bind(&effectNames1[2], "E1Name3")
	c.Params.Bind(&effectNames1[3], "E1Name4")

	c.Params.Bind(&effectNames2[0], "E2Name1")
	c.Params.Bind(&effectNames2[1], "E2Name2")
	c.Params.Bind(&effectNames2[2], "E2Name3")
	c.Params.Bind(&effectNames2[3], "E2Name4")

	c.Params.Bind(&effectNames3[0], "E3Name1")
	c.Params.Bind(&effectNames3[1], "E3Name2")
	c.Params.Bind(&effectNames3[2], "E3Name3")
	c.Params.Bind(&effectNames3[3], "E3Name4")
	c.Validation.Required(form.Vac1).Message("First dose of vaccine required for submit")
	if !c.Validation.HasErrors() {
		/*var effx models.Effects
		effx.Name = "Febra"
		effx.FormID = 1
		slice := []models.Effects{effx}
		form.Effects1 = slice*/
		var effects1 []models.Effects
		for i, name := range effectNames1 {
			if name != "" {
				var effect models.Effects
				effect.Name = dataNames[i]
				effects1 = append(effects1, effect)
			}
		}
		var effects2 []models.Effects
		for i, name := range effectNames2 {
			if name != "" {
				var effect models.Effects
				effect.Name = dataNames[i]
				effects2 = append(effects2, effect)
			}
		}
		var effects3 []models.Effects
		for i, name := range effectNames3 {
			if name != "" {
				var effect models.Effects
				effect.Name = dataNames[i]
				effects3 = append(effects3, effect)
			}
		}
		form.Effects1 = effects1
		form.Effects2 = effects2
		form.Effects3 = effects3
		app.DB.Create(&form)
	}
		
	return c.Render()
}

func (c App) DataSecFX() revel.Result {
	return c.Render()
}

func (c App) DataCovid() revel.Result {
	return c.Render()
}