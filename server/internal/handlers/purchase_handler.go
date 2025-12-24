package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"procurement-system/config"
	"procurement-system/internal/models"
)

type PurchaseItemRequest struct {
	ItemID uint `json:"item_id"`
	Qty    int  `json:"qty"`
}

type PurchaseRequest struct {
	SupplierID uint                  `json:"supplier_id"`
	Items      []PurchaseItemRequest `json:"items"`
}

func CreatePurchase(c *fiber.Ctx) error {
	var req PurchaseRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	userID := uint(c.Locals("user_id").(float64))

	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "DB Transaction failed"})
	}

	purchasing := models.Purchasing{
		Date:       time.Now(),
		SupplierID: req.SupplierID,
		UserID:     userID,
		GrandTotal: 0,
	}

	if err := tx.Create(&purchasing).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create header"})
	}

	var grandTotal int64 = 0

	for _, i := range req.Items {
		var item models.Item
		if err := tx.First(&item, i.ItemID).Error; err != nil {
			tx.Rollback()
			return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
		}

		subTotal := item.Price * int64(i.Qty)
		grandTotal += subTotal

		detail := models.PurchasingDetail{
			PurchasingID: purchasing.ID,
			ItemID:       item.ID,
			Qty:          i.Qty,
			SubTotal:     subTotal,
		}

		if err := tx.Create(&detail).Error; err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create detail"})
		}

		newStock := item.Stock + i.Qty
		if err := tx.Model(&item).Update("stock", newStock).Error; err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update stock"})
		}
	}

	if err := tx.Model(&purchasing).Update("grand_total", grandTotal).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update grand total"})
	}

	tx.Commit()

	go func(p models.Purchasing) {
		webhookURL := os.Getenv("WEBHOOK_URL")
		if webhookURL != "" {
			jsonData, _ := json.Marshal(p)
			http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
		}
	}(purchasing)

	return c.Status(201).JSON(fiber.Map{"message": "Success", "data": purchasing})
}