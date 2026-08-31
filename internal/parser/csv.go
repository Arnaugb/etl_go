package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Arnaugb/etl_go/internal/transformer"
)

// ReadCSV opens a CSV file and decodes the workout data into a slice.
func ReadCSV(filepath string) ([]transformer.Workout, error){
	
	file, err := os.Open(filepath)
	if err != nil{
		return nil, fmt.Errorf("Error al abrir el archivo: %w", err)
	}
	defer file.Close()
	

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil && err != io.EOF{
		return nil, fmt.Errorf("error leyendo la cabecera: %w", err)
	}

	var workouts []transformer.Workout

	for {
		record, err := reader.Read()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Printf("Error leyendo linea: %v", err)
			continue
		}

		if len(record) != 6{
			log.Printf("Linea corrupta (longitud incorrecta): %v\n", record)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil{
			log.Printf("ID inválido en registro %v: %v\n", record[0], err)
			continue
		}

		fecha, _ := time.Parse("2006-01-02", record[1])

		tipoDeporte := record[2]

		distance, err := strconv.ParseFloat(record[3], 64)
		if err != nil{
			log.Printf("Distancia inválida en registro %v: %v\n", record[3], err)
			continue
		}

		duration, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Printf("Duración inválida en registro %v: %v\n", record[4], err)
			continue
		}

		frecuenciaCard, err := strconv.Atoi(record[5])
		if err != nil{
			log.Printf("Frecuencia Cardíaca inválida en registro %v: %v\n", record[5], err)
			continue
		}


		workout := transformer.Workout{
			ID: id,
			Fecha: fecha,
			TipoDeporte: tipoDeporte,
			DistanciaKm: distance,
			Duration: duration,
			FrecuenciaCard: frecuenciaCard,
		}		
		workouts = append(workouts, workout)
	}
	return workouts, nil 
}