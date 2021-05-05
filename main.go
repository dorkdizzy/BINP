package main

import (
	"math/rand"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("NoBugIsPossible")

	var theArray [5]string
	theArray[0] = "SOLI A RAID TON LIVE !"
	theArray[1] = "TU N'AS PAS MANGÉ DEPUIS LONGTEMPS ! "
	theArray[2] = "This software has finished with failccess"
	theArray[3] = "En vrai t'es moche"
	theArray[4] = "Are you Shure this website is safe ? "

	hello := widget.NewLabel("ERREUR B*TARD !")
	hello2 := widget.NewLabel("Nous n'avons pas pu integrer du son dans l'application. \n Normalement Il y a un son d'alarme, \n mais c'etait dur a implementer. Donc s'il vous plait, \n veuillez chanter \n la note LA pendant toute l'utilisation du logiciel, \n pour etre dans l'ambiance. Merci. \n Signé les dorkies ! ")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Cliquez ici pour fixer l'erreur", func() {
			randomnumber := rand.Intn(5 - 0)
			hello.SetText(theArray[randomnumber])
		}),
		hello2,
	))

	w.ShowAndRun()
}
