package transformer

import (
	"log"
	"time"
)


type Workout struct{
	ID int `json:"id"`
	Fecha time.Time `json:"fecha"`
	TipoDeporte string `json:"tipo_deporte"`
	DistanciaKm float64 `json:"distancia_km"`
	Duration float64 `json:"duracion_minutos"`
	FrecuenciaCard int `json:"frecuencia_cardiaca_media,omitempty"`
}

type WorkoutProcessed struct{
	Workout
	RitmoMedio float64 `json:"ritmo_medio_min_km"`
}

// ProcessData takes raw workout data, filters invalid entries, and calculates derived metrics.
func ProcessData(data []Workout) []WorkoutProcessed{
	
	results := make([]WorkoutProcessed, 0, len(data))

	for _, w := range data{
		if w.DistanciaKm <= 0 || w.Duration <= 0 {
			log.Printf("ID %d descartado: distancia (%v) o duración (%v) inválida\n", w.ID, w.DistanciaKm, w.Duration)
			continue
		}

		if w.FrecuenciaCard != 0 && (w.FrecuenciaCard < 40 || w.FrecuenciaCard > 220) {
			log.Printf("ID %d descartado: frecuencia cardíaca anómala (%d)\n", w.ID, w.FrecuenciaCard)
			continue
		}

		if w.Fecha.IsZero(){
			log.Printf("ID %d descartado: fecha nula o no parseable\n", w.ID)
			continue
		}

		ritmo := w.Duration / w.DistanciaKm

		newWorkout := WorkoutProcessed{
			Workout: w,
			RitmoMedio: ritmo,
		}

		results = append(results, newWorkout)

	}
	return results
}