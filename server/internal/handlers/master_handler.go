package handlers

import (
	"github.com/gofiber/fiber/v2"
	"procurement-system/config"
	"procurement-system/internal/models"
)


func GetSuppliers(c *fiber.Ctx) error {
	var suppliers []models.Supplier
	config.DB.Order("id desc").Find(&suppliers)
	return c.JSON(suppliers)
}

func GetSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier
	if result := config.DB.First(&supplier, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Supplier not found"})
	}
	return c.JSON(supplier)
}

func CreateSupplier(c *fiber.Ctx) error {
	var supplier models.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	config.DB.Create(&supplier)
	return c.Status(201).JSON(supplier)
}

func UpdateSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier
	if result := config.DB.First(&supplier, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Supplier not found"})
	}
	
	var updateData models.Supplier
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	supplier.Name = updateData.Name
	supplier.Email = updateData.Email
	supplier.Address = updateData.Address
	config.DB.Save(&supplier)
	return c.JSON(supplier)
}

func DeleteSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier
	if result := config.DB.First(&supplier, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Supplier not found"})
	}
	config.DB.Delete(&supplier)
	return c.JSON(fiber.Map{"message": "Supplier deleted"})
}



func GetItems(c *fiber.Ctx) error {
	var items []models.Item
	config.DB.Order("id desc").Find(&items)
	return c.JSON(items)
}

func GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item
	if result := config.DB.First(&item, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}
	return c.JSON(item)
}

func CreateItem(c *fiber.Ctx) error {
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	config.DB.Create(&item)
	return c.Status(201).JSON(item)
}

func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item
	if result := config.DB.First(&item, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}

	var updateData models.Item
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	item.Name = updateData.Name
	item.Price = updateData.Price
	item.Stock = updateData.Stock
	config.DB.Save(&item)
	return c.JSON(item)
}

func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item
	if result := config.DB.First(&item, id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}
	config.DB.Delete(&item)
	return c.JSON(fiber.Map{"message": "Item deleted"})
}