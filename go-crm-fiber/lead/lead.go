package lead

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/Jonathansoufer/go-crm-fiber/database"
	"github.com/gofiber/fiber/v2"
)

type Lead struct {
	gorm.Model
	ID        	uint   `json:"id"`
	Name 		string `json:"name"`
	Company  	string `json:"company"`
	Email     	string `json:"email"`
	Status    	string `json:"status"`
	Phone    	uint   `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.GetDB()
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	db := database.GetDB()
	id := c.Params("id")
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(500).SendString("No lead found with ID")
	}
	return c.JSON(lead)
}

func CreateLead(c *fiber.Ctx) error {
	db := database.GetDB()
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(500).SendString("Couldn't parse JSON")
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id:= c.Params("id")
	db := database.GetDB()

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		return c.Status(500).SendString("No lead found with ID")
	}
	db.Delete(&lead)
	return c.SendString("Lead successfully deleted")
}
