# Workout Data ETL Pipeline

Un pipeline ETL (Extract, Transform, Load) secuencial desarrollado en Go para procesar, limpiar y estandarizar datos de entrenamientos deportivos.

## Arquitectura

El proyecto sigue el estándar de diseño de directorios de la industria (Standard Go Project Layout), aislando la lógica de dominio del orquestador principal:

- **`cmd/etl/`**: Punto de entrada de la aplicación. Orquesta el flujo de lectura, transformación y escritura sin contener lógica de negocio.
- **`internal/parser/`**: Módulo de extracción. Decodifica archivos CSV y realiza el tipado estricto inicial, descartando registros con estructuras o tipos de datos inválidos.
- **`internal/transformer/`**: Módulo de procesamiento. Aplica las reglas de calidad de datos (filtrado de métricas biológicamente imposibles, duraciones nulas, errores de fecha) y calcula métricas derivadas (`ritmo_medio_min_km`).
- **`internal/exporter/`**: Módulo de carga. Serializa las estructuras limpias y las exporta a un archivo JSON utilizando _streaming_ (`json.NewEncoder`) para mantener un consumo de memoria estable (O(1)).

## Prerrequisitos

- Go (versión instalada en el sistema).

## Ejecución

Para ejecutar el pipeline, abre una terminal en la raíz del proyecto y ejecuta el orquestador:

```bash
go run cmd/etl/main.go
```

## Flujo de Datos

1. **Input**: Ingesta el dataset sucio desde data/raw_workouts.csv.
2. **Procesamiento**: Aplica el filtrado en memoria basado en las reglas de negocio estipuladas.
3. **Output**: Genera el archivo estandarizado final en data/clean_workouts.json.
