package exporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Arnaugb/etl_go/internal/transformer"
)

// WriteJSON serializes the processed workout data and saves it to a JSON file.
func WriteJSON(filepath string, data []transformer.WorkoutProcessed) error{

	file, err := os.Create(filepath)
	if err != nil{
		return fmt.Errorf("Error al crear el archivo %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil{
		return fmt.Errorf("Error al codificar los datos a JSON %w", err)
	}

	return nil
}