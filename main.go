package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas()

	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizza)
	router.DELETE("/pizzas/:id", deletePizzaById)
	router.PUT("/pizzas/:id", updatePizzaById)
	router.Run(":8080")
}

func getPizzas(c *gin.Context) {
	c.JSON(200, pizzas)
}

func deletePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H {
			"error": err.Error() })
		return
	}

	for i, p := range pizzas {
		if p.ID == id {
			pizzas = append(pizzas[:i], pizzas[i+1:]...)
			savePizza()
			c.JSON(200, gin.H {
				"message": "Pizza deleted" })
			return
		}
	}
	c.JSON(404, gin.H { "message": "Pizza not found" })
}

func updatePizzaById(c *gin.Context) {

}

func getPizza(c *gin.Context) {
	idParam := c.Param("id")
	id, errr := strconv.Atoi(idParam)
	if errr != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	for _, pizza := range pizzas {
		if pizza.ID == id {
			c.JSON(200, pizza)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Pizza not found"})
}

func postPizzas(c *gin.Context) {
	var pizza models.Pizza
	if err := c.ShouldBindJSON(&pizza); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	
	pizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, pizza)
	
	savePizza()
	c.JSON(201, pizzas)
}

func loadPizzas() {
	file, err := os.Open("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
}

func savePizza() {
	file, err := os.Create("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}