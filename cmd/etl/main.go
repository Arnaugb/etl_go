package main

import (
	"log"
	"path/filepath"

	"github.com/Arnaugb/etl_go/internal/exporter"
	"github.com/Arnaugb/etl_go/internal/parser"
	"github.com/Arnaugb/etl_go/internal/transformer"
)


func main(){
	inputPath := filepath.Join("..", "..", "data", "raw_workouts.csv")
	outputPath := filepath.Join("..", "..", "data", "clean_workouts.json")

	log.Println("Iniciando proceso ETL...")

	rawData, err := parser.ReadCSV(inputPath)
	if err != nil{
		log.Fatalf("Fallo crítico en la extracción: %v", err)
	}
	log.Printf("Extracción completada. %d registros leídos.\n", len(rawData))

	cleanData := transformer.ProcessData(rawData)
	log.Printf("Transformación completada. %d registros válidos tras la limpieza.\n", len(cleanData))

	err = exporter.WriteJSON(outputPath, cleanData)
	if err != nil{
		log.Fatalf("Fallo crítico en la exportación: %v", err)
	}

	log.Println("ETL finalizado con éxito. Revisa el archivo clean_workouts.json.")
}