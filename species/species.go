package species

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/wakatara/spt/database"
)

type Species struct {
	gorm.Model
	Name          string `json:"name"`
	CommonName    string `json:"common_name"`
	Analog        string `json:"analog"`
	Class         string `json:"class"`
	Order         string `json:"order"`
	Family        string `json:"family"`
	Genus         string `json:"genus"`
	Species       string `json:"species"`
	HasPhyloStudy bool   `json:"has_phylo_study"`
}

func GetAllSpecies(c *fiber.Ctx) {
	db := database.DBConn
	var species []Species
	db.Find(&species)
	c.JSON(species)
}

func GetSpecies(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var species Species
	db.Find(&species, id)
	c.JSON(species)
}

func NewSpecies(c *fiber.Ctx) {
	db := database.DBConn

	species := new(Species)
	if err := c.BodyParser(species); err != nil {
		c.Status(503).Send(err)
		return
	}

	// Test fixture
	// species.Name = "Green and black poison dart frog"
	// species.CommonName = "Curare frog"
	// species.Analog = ""
	// species.Class = "Amphibia"
	// species.Order = "Anura"
	// species.Family = "Dendrobatidae"
	// species.Genus = "Dendrobates"
	// species.Species = "auratus"
	// species.HasPhyloStudy = true

	db.Create(&species)
	c.JSON(species)
}

func DeleteSpecies(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var species Species

	db.First(&species, id)
	if species.Name == "" {
		c.Status(500).Send("No such species exists with given ID.")
		return
	}
	db.Delete(&species)
	c.Send("Species successfully deleted.")
}
