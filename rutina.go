package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var tipoCom string
	var tClientes int
	var tEspera int

	fmt.Print("Comportamiento a seguir [pymes/retail]: ")
	fmt.Scanf("%s", &tipoCom)

	fmt.Print("Tiempo de espera entre envio de ordenes de clientes: ")
	fmt.Scanf("%d", &tClientes)

	fmt.Print("Tiempo de espera del camion: ")
	fmt.Scanf("%d", &tEspera)

	csvArch, err := os.Open("pymes.csv")
	if err != nil {
		log.Fatalln("No se pudo abrir el archivo csv", err)
	}

	r := csv.NewReader(csvArch)

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(record[0])
	}

}
