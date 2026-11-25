package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ============================================================
// ข้อ 2: สร้างตารางเดียวกันด้วย GORM  (ฐานข้อมูล rental_gorm)
// เติม struct และ struct tag ให้ AutoMigrate สร้างตารางเหมือนข้อ 1
// ดูตาราง GORM Tags ได้ที่ reference/gorm-tags.md
// ============================================================

// TODO: struct Member
//   ID, MemberCode, FullName, Phone
type Member struct {
	ID uint `gorm:"primaryKey"`
	// TODO: MemberCode string `gorm:"size:10;not null;unique"`
	// TODO: FullName
	// TODO: Phone
}

// TODO: struct Bike
//   ID, Code, Model, Deposit (DEFAULT 500, CHECK >= 0)
type Bike struct {
	ID uint `gorm:"primaryKey"`
	// TODO: Code
	// TODO: Model
	// TODO: Deposit int `gorm:"not null;default:500;check:deposit >= 0"`
}

// TODO: struct Rental
//   ID, MemberID, Member, BikeID, Bike, Status, StartAt, DueAt, CancelledAt/ReturnedAt
//   - Member คู่กับ MemberID แบบ Belongs To + constraint:OnDelete:CASCADE
//   - Bike   คู่กับ BikeID   แบบ Belongs To + constraint:OnDelete:RESTRICT
//   - Status DEFAULT 'active', CHECK เฉพาะ 'active'/'returned'
//   - DueAt CHECK due_at > start_at
//   - ReturnedAt เว้นว่างได้ (ใช้ *time.Time)
type Rental struct {
	ID uint `gorm:"primaryKey"`
	// TODO: fields
	_ time.Time // ลบบรรทัดนี้ออกเมื่อเพิ่มฟิลด์เวลาจริง
}

func main() {
	// การเชื่อมต่อฐานข้อมูลเตรียมไว้ให้แล้ว
	dsn := "host=localhost user=lab password=lab1234 dbname=rental_gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connect failed: ", err)
	}

	if err := db.AutoMigrate(&Member{}, &Bike{}, &Rental{}); err != nil {
		log.Fatal("migrate failed: ", err)
	}
	log.Println("AutoMigrate done")
}
