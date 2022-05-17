package gui

import (
	"fmt"
	"strconv"

	"passwordgenerator/entrydigital"
	"passwordgenerator/passwordgenerator"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	MAINWINDOWTITLE = "Password Generator"
)

func LoadAndRunWindow() {
	// démarre l'application
	app := app.New()
	mainWindow := app.NewWindow(MAINWINDOWTITLE)
	mainWindow.CenterOnScreen()

	// variables booléennes pour les checkbox
	var blowLetter bool = false
	var bbigLetter bool = false
	var bnumber bool = false
	var bspecialChar bool = false
	var bcheckPasswordLeak bool = false

	// crée et charge les widget
	lowLetterCheckbox := widget.NewCheck("Lettres minuscules", func(b bool) {
		if b {
			blowLetter = true
		} else {
			blowLetter = false
		}

	})
	bigLetterCheckbox := widget.NewCheck("Lettres majusculles", func(b bool) {
		if b {
			bbigLetter = true
		} else {
			bbigLetter = false
		}
	})
	numberCheckbox := widget.NewCheck("Nombres", func(b bool) {
		if b {
			bnumber = true
		} else {
			bnumber = false
		}
	})

	specialCharCheckbox := widget.NewCheck("Caractères spéciaux", func(b bool) {
		if b {
			bspecialChar = true
		} else {
			bspecialChar = false
		}
	})

	minimalLengthPasswordLabel := widget.NewLabel("Longueur minimale du mot de passe:")
	minimalLengthPassword := entrydigital.NewNumericalEntry()
	minimalLengthPassword.SetText("8")

	maximalLengthPasswordLabel := widget.NewLabel("Longueur maximale du mot de passe:")
	maximalLengthPassword := entrydigital.NewNumericalEntry()
	maximalLengthPassword.SetText("50")

	checkPasswordLeak := widget.NewCheck("Générer un mot de passe qui n'a pas fuité sur have i been pwned", func(b bool) {
		if b {
			bcheckPasswordLeak = true
		} else {
			bcheckPasswordLeak = false
		}
	})

	passwordGenerateLabel := widget.NewLabel("Mot de passe généré:")
	passwordGenerate := widget.NewLabel("")

	validationButton := widget.NewButton("Générer un mot de passe", func() {
		//fmt.Println("génère un mot de passe")
		//fmt.Println("lettre maj: ", bbigLetter, "\nlettre min: ", blowLetter, "\nnombre:", bnumber, "\nchar speciaux: ", bspecialChar, "\npassword leak: ", bcheckPasswordLeak)

		// génère le mot de passe
		lenMin, _ := strconv.Atoi(minimalLengthPassword.Text)
		lenMax, _ := strconv.Atoi(maximalLengthPassword.Text)

		if blowLetter || bbigLetter || bnumber || bspecialChar {
			passwordGenerate.SetText(passwordgenerator.GeneratePassword(lenMin, lenMax, blowLetter, bbigLetter, bnumber, bspecialChar, bcheckPasswordLeak))
		}
	})

	clipbordButton := widget.NewButton("Copier le mot de passe", func() {
		mainWindow.Clipboard().SetContent(passwordGenerate.Text)
		fmt.Println("copier le mot de passe ...")
	})

	// crée le layout

	containerForm := fyne.NewContainerWithLayout(layout.NewGridLayout(2), lowLetterCheckbox, bigLetterCheckbox, numberCheckbox, specialCharCheckbox, minimalLengthPasswordLabel, minimalLengthPassword, maximalLengthPasswordLabel, maximalLengthPassword, checkPasswordLeak)

	containerResult := fyne.NewContainerWithLayout(layout.NewGridLayout(2), passwordGenerateLabel, passwordGenerate)
	container2 := fyne.NewContainerWithLayout(layout.NewGridLayout(1), validationButton, containerResult, clipbordButton)

	containerGlobal := fyne.NewContainerWithLayout(layout.NewGridLayout(1), containerForm, container2)

	mainWindow.SetContent(containerGlobal)

	mainWindow.Show()

	app.Run()

}
