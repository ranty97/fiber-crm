package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/ranty97/fiber-crm-basic/database"
)

type Lead struct {
	gorm.Model
	Name  string `json:"name"`
	Team  string `json:"team"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBconnection
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBconnection
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func Delete(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBconnection
	var lead Lead
	db.First(&lead, id)
	db.Delete(&lead)
	c.JSON(lead)
}

func CreateLead(c *fiber.Ctx) {
	db := database.DBconnection
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}
