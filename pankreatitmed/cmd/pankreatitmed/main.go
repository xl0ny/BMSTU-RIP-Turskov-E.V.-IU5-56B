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
		minioBase = "http://127.0.0.1:9000/services-images"
	}
	repo := repository.NewRepository(minioBase)
	h := handler.NewHandler(repo, "templates")
	s := api.NewServer(h)

	log.Println("Listening on http://localhost:8080/criteria")
	if err := s.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
