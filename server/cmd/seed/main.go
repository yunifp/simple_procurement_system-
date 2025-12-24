package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"procurement-system/internal/models"
)

func main() {
	
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found (checking environment vars)")
	}


	dsn := "root:@tcp(127.0.0.1:3306)/procurement_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal konek database untuk seeding. Pastikan MySQL nyala & DB 'procurement_db' sudah dibuat!")
	}

	fmt.Println("Mulai Seeding Data...")

	// --- SEED USER ---
	// Password: password123
	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), 14)
	
	users := []models.User{
		{Username: "admin", Password: string(hash), Role: "admin"},
		{Username: "staff", Password: string(hash), Role: "staff"},
	}

	for _, user := range users {
		if err := db.FirstOrCreate(&models.User{}, models.User{Username: user.Username}).Updates(user).Error; err != nil {
			log.Printf("Gagal seed user %s: %v", user.Username, err)
		}
	}
	fmt.Println("✅ Users created (Pass: password123)")

	// --- SEED SUPPLIERS ---
	suppliers := []models.Supplier{
		{Name: "PT. Teknologi Maju", Email: "sales@tekno.com", Address: "Jakarta Selatan"},
		{Name: "CV. Sumber Makmur", Email: "admin@sumber.com", Address: "Surabaya"},
		{Name: "Toko Besi Jaya", Email: "owner@besijaya.com", Address: "Bandung"},
	}

	for _, s := range suppliers {
		db.FirstOrCreate(&s, models.Supplier{Email: s.Email})
	}
	fmt.Println("✅ Suppliers created")

	// --- SEED ITEMS ---
	items := []models.Item{
		{Name: "Laptop Thinkpad X1", Stock: 10, Price: 15000000},
		{Name: "Mouse Logitech", Stock: 50, Price: 150000},
		{Name: "Monitor Dell 24 Inch", Stock: 20, Price: 2500000},
		{Name: "Kabel HDMI", Stock: 100, Price: 45000},
	}

	for _, item := range items {
		db.FirstOrCreate(&item, models.Item{Name: item.Name})
	}
	fmt.Println("✅ Items created")

	fmt.Println("🎉 Seeding Selesai! Sekarang jalankan 'go run cmd/main.go'")
}