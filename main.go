package main

import "os"
import g "github.com/AllenDang/giu"

var texto string

func formulario() {
	g.SingleWindow().Layout(g.Label("Argumentos: "),g.InputTextMultiline(&texto))
	texto = ""
	if len(os.Args) != 1 {
		for i := 1; i < len(os.Args); i++ {
			texto = texto + os.Args[i] + "\n"
		}
	}
}

func main() {
	ventana := g.NewMasterWindow("AppURL", 500, 500, 0)
	ventana.Run(formulario)
}
