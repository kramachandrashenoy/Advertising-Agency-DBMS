package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sikehish/Advertising-Agency-DBMS/configs"
	"github.com/sikehish/Advertising-Agency-DBMS/internal/models"
	"gorm.io/gorm"
)

func GetAllClients(c *fiber.Ctx) error {
	var clients []models.Client
	configs.DB.Find(&clients)
	return c.Status(200).JSON(clients)
}

func GetClientByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var client models.Client
	err := configs.DB.First(&client, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Client not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	//OR
	// client = config.Database.Find(&dog, id)

    // if client.RowsAffected == 0 {
    //     return c.SendStatus(404)
    // }

	return c.Status(200).JSON(client)
}


func AddClient(c *fiber.Ctx) error {
    client := new(models.Client) //returns a pointer

    if err := c.BodyParser(client); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    configs.DB.Create(&client)
    return c.Status(fiber.StatusCreated).JSON(client)
}

func 	DeleteClient(c *fiber.Ctx) error {

	clientID := c.Params("id")

	var existingClient models.Client
	result := configs.DB.First(&existingClient, clientID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Client not found"})
	}

	// Delete the client from the database
	configs.DB.Delete(&existingClient)

	// Return a success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Client deleted successfully"})
}
