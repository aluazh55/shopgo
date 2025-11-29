package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Hotel struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Rating    float64   `json:"rating"` // или float64, если надо
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func connectPdx() (*pgxpool.Pool, error) {
	connectionString := "postgres://postgres:Alua2004@localhost:5432/mydb"

	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil

}

func createHotel(c *gin.Context) {
	dbAny, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurs"})
		return
	}

	var hotel Hotel
	db := dbAny.(*pgxpool.Pool)
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.QueryRow(
		context.Background(),
		`INSERT INTO hotels (name, address,rating) 
    VALUES ($1,$2,$3)
    RETURNING id, created_at,updated_at`,
		hotel.Name, hotel.Address, hotel.Rating,
	).Scan(&hotel.Id, &hotel.CreatedAt, &hotel.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occure"})
	}
	c.JSON(http.StatusCreated, hotel)

}

func getHotels(c *gin.Context) {
	dbAny, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurs"})
		return
	}

	db := dbAny.(*pgxpool.Pool)
	rows, err := db.Query(
		context.Background(),
		"Select id, name, address,rating, created_at,updated_at FROM hotels")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurs"})
		return
	}
	defer rows.Close()

	hotels := []Hotel{}
	for rows.Next() {
		var h Hotel
		err := rows.Scan(&h.Id, &h.Name, &h.Address, &h.Rating, &h.CreatedAt, &h.UpdatedAt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurs"})
			return
		}
		hotels = append(hotels, h)
	}
	c.JSON(http.StatusOK, hotels)
}

func main() {

	pool, err := connectPdx()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", pool)
		c.Next()

	})

	// r.Get("/home", func())
	r.POST("/hotels", createHotel)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running")
	})

	r.GET("/hotels/:id", func(c *gin.Context) {
		dbAny, exists := c.Get("db")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurs"})
			return
		}
		db := dbAny.(*pgxpool.Pool)

		id := c.Param("id")
		var hotel Hotel
		err := db.QueryRow(
			context.Background(),
			"SELECT id, name, address, rating, created_at, updated_at FROM hotels WHERE id=$1",
			id,
		).Scan(&hotel.Id, &hotel.Name, &hotel.Address, &hotel.Rating, &hotel.CreatedAt, &hotel.UpdatedAt)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "hotel not found"})
			return
		}

		c.JSON(http.StatusOK, hotel)
	})

	r.GET("/gethotels", getHotels)
	r.Run()

}
