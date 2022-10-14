package main

import (
	"os"
	"fmt"
	"strconv"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"desafio-goweb-juanurbano/cmd/server/handler"
	"desafio-goweb-juanurbano/internal/domain"
	"desafio-goweb-juanurbano/internal/tickets"
)

func main() {

	// Cargo csv.
	db, err := LoadTicketsFromFile("./tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

        repo := tickets.NewRepository(db)
	service := tickets.NewService(repo)
	t := handler.NewService(service)

	r := gin.Default()
	r.GET("/ping", t.Status())

        // Rutas a desarollar:
	pr := r.Group("ticket") 
        {
	        // GET - “/ticket/getByCountry/:dest”
                pr.GET("/getByCountry/:dest", t.GetTicketsByCountry())
	        // GET - “/ticket/getAverage/:dest”
                pr.GET("/getAverage/:dest", t.AverageDestination())
        }

	if err := r.Run(); err != nil {
		panic(err)
	}
}



func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
