package main

import (
	"log"
	"os"

	"pankreatitmed/internal/api"
	"pankreatitmed/internal/app/handler"
	"pankreatitmed/internal/app/repository"
)

func main() {
	minioBase := os.Getenv("MINIO_PUBLIC_URL")
	if minioBase == "" {
		// В ЛР1 можно так: положить картинки в публичный бакет и указать его URL.
		// Во время защиты показать, что это именно адрес MinIO.
		minioBase = "http://127.0.0.1:9000/services-images"
	}
	repo := repository.NewRepository(minioBase)
	h := handler.NewHandler(repo, "templates")
	s := api.NewServer(h)

	log.Println("Listening on http://localhost:8080/services")
	if err := s.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
