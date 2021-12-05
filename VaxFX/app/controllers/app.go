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
	dataNames := []string{"Headache", "Fever", "Dizziness", "Exhaustion", "Arm pain", "Nausea"}
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
		var effects1 []models.Effects
		for i, name := range effectNames1 {
			if name != "" {
				var effect models.Effects
				effect.Name = dataNames[i]
				effect.WhichEffect = 1
				effects1 = append(effects1, effect)
			}
		}
		var effects2 []models.Effects
		for i, name := range effectNames2 {
			if name != "" {
				var effect models.Effects
				effect.Name = dataNames[i]
				effect.WhichEffect = 2
				effects2 = append(effects2, effect)
			}
		}
		var effects3 []models.Effects
		for i, name := range effectNames3 {
			if name != "" {
				var effect models.Effects
				effect.Name = dataNames[i]
				effect.WhichEffect = 3
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
	TotalUsers1 := 0
	app.DB.Model(&models.Form{}).Where("Vac1 <> ?", "").Count(&TotalUsers1)
	
	var forms []models.Form
	app.DB.Find(&forms)
	var formsPfi []models.Form
	app.DB.Model(&models.Form{}).Where("Vac1 = ?", "Pfizer").Find(&formsPfi)
	var formsMrn []models.Form
	app.DB.Model(&models.Form{}).Where("Vac1 = ?", "Moderna").Find(&formsMrn)
	var formsAzn []models.Form
	app.DB.Model(&models.Form{}).Where("Vac1 = ?", "AstraZeneca").Find(&formsAzn)
	var formsJJ []models.Form
	app.DB.Model(&models.Form{}).Where("Vac1 = ?", "Johnson & Johnson").Find(&formsJJ)
	Headaches1 := 0
	Fevers1 := 0
	Dizzinesses1 := 0
	Exhaustions1 := 0
	ArmPains1 := 0
	Nauseas1 := 0
	maxVax := []int{0, 0, 0, 0}
	for _, form := range forms {
		app.DB.Model(&form.Effects1).Where("Name = ? AND Which_Effect = ?", "Headache", 1).Count(&Headaches1)
		app.DB.Model(&form.Effects1).Where("Name = ? AND Which_Effect = ?", "Fever", 1).Count(&Fevers1)
		app.DB.Model(&form.Effects1).Where("Name = ? AND Which_Effect = ?", "Dizziness", 1).Count(&Dizzinesses1)
		app.DB.Model(&form.Effects1).Where("Name = ? AND Which_Effect = ?", "Exhaustion", 1).Count(&Exhaustions1)
		app.DB.Model(&form.Effects1).Where("Name = ? AND Which_Effect = ?", "Arm pain", 1).Count(&ArmPains1)
		app.DB.Model(&form.Effects1).Where("Name = ? AND Which_Effect = ?", "Nausea", 1).Count(&Nauseas1)
	}
	oldValues := []int{0, 0, 0, 0}
	oldValues[0] = maxVax[0]
	oldValues[1] = maxVax[1]
	oldValues[2] = maxVax[2]
	oldValues[3] = maxVax[3]
	for _, form := range formsPfi {
		app.DB.Model(&form.Effects1).Where("Which_Effect = ?", 1).Count(&maxVax[0])
	}
	oldValues[0] = maxVax[0]
	maxVax[1] = oldValues[1]
	maxVax[2] = oldValues[2]
	maxVax[3] = oldValues[3]
	for _, form := range formsMrn {
		app.DB.Model(&form.Effects1).Where("Which_Effect = ?", 1).Count(&maxVax[1])
	}
	oldValues[1] = maxVax[1]
	maxVax[0] = oldValues[0]
	maxVax[2] = oldValues[2]
	maxVax[3] = oldValues[3]
	for _, form := range formsAzn {
		app.DB.Model(&form.Effects1).Where("Which_Effect = ?", 1).Count(&maxVax[2])
	}
	oldValues[2] = maxVax[2]
	maxVax[1] = oldValues[1]
	maxVax[0] = oldValues[0]
	maxVax[3] = oldValues[3]
	for _, form := range formsJJ {
		app.DB.Model(&form.Effects1).Where("Which_Effect = ?", 1).Count(&maxVax[3])
	}
	oldValues[3] = maxVax[3]
	maxVax[1] = oldValues[1]
	maxVax[2] = oldValues[2]
	maxVax[0] = oldValues[0]
	mx := -1
	mxi := -1
	for i, _ := range maxVax {
		fmt.Printf("Uite: %v %v\n", i, maxVax[i])
		if maxVax[i] > mx {
			mx = maxVax [i]
			mxi = i
		}
	}
	WorstVax1 := ""
	if mxi == 0 {
		WorstVax1 = "Pfizer"
	} else if mxi == 1 {
		WorstVax1 = "Moderna"
	} else if mxi == 2 {
		WorstVax1 = "AstraZeneca"
	} else {
		WorstVax1 = "Johnson & Johnson"
	}

	TotalUsers2 := 0
	app.DB.Model(&models.Form{}).Where("Vac2 <> ?", "").Count(&TotalUsers2)
	
	//var forms []models.Form
	app.DB.Find(&forms)
	//var formsPfi []models.Form
	app.DB.Model(&models.Form{}).Where("Vac2 = ?", "Pfizer").Find(&formsPfi)
	//var formsMrn []models.Form
	app.DB.Model(&models.Form{}).Where("Vac2 = ?", "Moderna").Find(&formsMrn)
	//var formsAzn []models.Form
	app.DB.Model(&models.Form{}).Where("Vac2 = ?", "AstraZeneca").Find(&formsAzn)
	//var formsJJ []models.Form
	app.DB.Model(&models.Form{}).Where("Vac2 = ?", "Johnson & Johnson").Find(&formsJJ)
	Headaches2 := 0
	Fevers2 := 0
	Dizzinesses2 := 0
	Exhaustions2 := 0
	ArmPains2 := 0
	Nauseas2 := 0
	maxVax2 := []int{0, 0, 0, 0}
	for _, form := range forms {
		app.DB.Model(&form.Effects2).Where("Name = ? AND Which_Effect = ?", "Headache", 2).Count(&Headaches2)
		app.DB.Model(&form.Effects2).Where("Name = ? AND Which_Effect = ?", "Fever", 2).Count(&Fevers2)
		app.DB.Model(&form.Effects2).Where("Name = ? AND Which_Effect = ?", "Dizziness", 2).Count(&Dizzinesses2)
		app.DB.Model(&form.Effects2).Where("Name = ? AND Which_Effect = ?", "Exhaustion", 2).Count(&Exhaustions2)
		app.DB.Model(&form.Effects2).Where("Name = ? AND Which_Effect = ?", "Arm pain", 2).Count(&ArmPains2)
		app.DB.Model(&form.Effects2).Where("Name = ? AND Which_Effect = ?", "Nausea", 2).Count(&Nauseas2)
	}
	for _, form := range formsPfi {
		app.DB.Model(&form.Effects1).Where("Which_Effect = ?", 1).Count(&maxVax2[0])
	}
	for _, form := range formsMrn {
		app.DB.Model(&form.Effects1).Where("Which_Effect = ?", 1).Count(&maxVax2[1])
	}
	for _, form := range formsAzn {
		app.DB.Model(&form.Effects1).Where("Which_Effect = ?", 1).Count(&maxVax2[2])
	}
	for _, form := range formsJJ {
		app.DB.Model(&form.Effects1).Where("Which_Effect = ?", 1).Count(&maxVax2[3])
	}
	mx2 := -1
	mxi2 := -1
	for i, _ := range maxVax2 {
		fmt.Printf("Uite: %v %v\n", i, maxVax2[i])
		if maxVax2[i] > mx2 {
			mx2 = maxVax [i]
			mxi2 = i
		}
	}
	WorstVax2 := ""
	if mxi2 == 0 {
		WorstVax2 = "Pfizer"
	} else if mxi2 == 1 {
		WorstVax2 = "Moderna"
	} else if mxi2 == 2 {
		WorstVax2 = "AstraZeneca"
	} else {
		WorstVax2 = "Johnson & Johnson"
	}
	

	TotalUsers3 := 0
	app.DB.Model(&models.Form{}).Where("Booster <> ?", "").Count(&TotalUsers3)
	
	//var forms []models.Form
	app.DB.Find(&forms)
	//var formsPfi []models.Form
	app.DB.Model(&models.Form{}).Where("Booster = ?", "Pfizer").Find(&formsPfi)
	//var formsMrn []models.Form
	app.DB.Model(&models.Form{}).Where("Booster = ?", "Moderna").Find(&formsMrn)
	//var formsAzn []models.Form
	app.DB.Model(&models.Form{}).Where("Booster = ?", "AstraZeneca").Find(&formsAzn)
	//var formsJJ []models.Form
	app.DB.Model(&models.Form{}).Where("Booster = ?", "Johnson & Johnson").Find(&formsJJ)
	Headaches3 := 0
	Fevers3 := 0
	Dizzinesses3 := 0
	Exhaustions3 := 0
	ArmPains3 := 0
	Nauseas3 := 0
	maxVax3 := []int{0, 0, 0, 0}
	for _, form := range forms {
		app.DB.Model(&form.Effects3).Where("Name = ? AND Which_Effect = ?", "Headache", 3).Count(&Headaches3)
		app.DB.Model(&form.Effects3).Where("Name = ? AND Which_Effect = ?", "Fever", 3).Count(&Fevers3)
		app.DB.Model(&form.Effects3).Where("Name = ? AND Which_Effect = ?", "Dizziness", 3).Count(&Dizzinesses3)
		app.DB.Model(&form.Effects3).Where("Name = ? AND Which_Effect = ?", "Exhaustion", 3).Count(&Exhaustions3)
		app.DB.Model(&form.Effects3).Where("Name = ? AND Which_Effect = ?", "Arm pain", 3).Count(&ArmPains3)
		app.DB.Model(&form.Effects3).Where("Name = ? AND Which_Effect = ?", "Nausea", 3).Count(&Nauseas3)
	}
	for _, form := range formsPfi {
		app.DB.Model(&form.Effects3).Where("Which_Effect = ?", 3).Count(&maxVax3[0])
	}
	for _, form := range formsMrn {
		app.DB.Model(&form.Effects3).Where("Which_Effect = ?", 3).Count(&maxVax3[1])
	}
	for _, form := range formsAzn {
		app.DB.Model(&form.Effects3).Where("Which_Effect = ?", 3).Count(&maxVax3[2])
	}
	for _, form := range formsJJ {
		app.DB.Model(&form.Effects3).Where("Which_Effect = ?", 3).Count(&maxVax3[3])
	}
	mx3 := -1
	mxi3 := -1
	for i, _ := range maxVax3 {
		if maxVax3[i] > mx3 {
			mx3 = maxVax [i]
			mxi3 = i
		}
	}
	WorstVax3 := ""
	if mxi3 == 0 {
		WorstVax3 = "Pfizer"
	} else if mxi3 == 1 {
		WorstVax3 = "Moderna"
	} else if mxi3 == 2 {
		WorstVax3 = "AstraZeneca"
	} else {
		WorstVax3 = "Johnson & Johnson"
	}

	return c.Render(TotalUsers1, Headaches1, Fevers1, Dizzinesses1, Exhaustions1, ArmPains1, Nauseas1, WorstVax1,
		TotalUsers2, Headaches2, Fevers2, Dizzinesses2, Exhaustions2, ArmPains2, Nauseas2, WorstVax2,
		TotalUsers3, Headaches3, Fevers3, Dizzinesses3, Exhaustions3, ArmPains3, Nauseas3, WorstVax3)
}

func (c App) DataCovid() revel.Result {
	return c.Render()
}