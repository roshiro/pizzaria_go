package handler

import (
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(200, data.Pizzas)
}

func DeletePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H {
			"error": err.Error() })
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizza()
			c.JSON(200, gin.H {
				"message": "Pizza deleted" })
			return
		}
	}
	c.JSON(404, gin.H { "message": "Pizza not found" })
}

func UpdatePizzaById(c *gin.Context) {
	idParam := c.Param("id")
	id, errr := strconv.Atoi(idParam)
	if errr != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedPizza models.Pizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			updatedPizza.ID = id
			data.Pizzas[i] = updatedPizza
			data.SavePizza()
			c.JSON(200, updatedPizza)
			return
		}
	}
}

func GetPizza(c *gin.Context) {
	idParam := c.Param("id")
	id, errr := strconv.Atoi(idParam)
	if errr != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	for _, pizza := range data.Pizzas {
		if pizza.ID == id {
			c.JSON(200, pizza)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Pizza not found"})
}

func PostPizzas(c *gin.Context) {
	var pizza models.Pizza
	if err := c.ShouldBindJSON(&pizza); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := service.ValidatePizzaPrice(&pizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	pizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, pizza)
	
	data.SavePizza()
	c.JSON(201, data.Pizzas)
}