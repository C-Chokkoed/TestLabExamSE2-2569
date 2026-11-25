# ชุดข้อสอบจำลอง — SE69 Lab03 (SQL + ORM + Schema)

ข้อสอบปฏิบัติการจำลอง ประยุกต์จาก Lab 03 รายวิชา ENG23 3032 Software Engineering
สถานการณ์: **ระบบเช่าจักรยานสาธารณะ (Public Bike Rental)** — 3 ตาราง, ความสัมพันธ์ 1:N สองคู่, constraint ครบชุด

> โครงสร้างคู่ขนานกับระบบจองห้องติวในชีทจริง แต่เปลี่ยนสถานการณ์และรายละเอียด เพื่อฝึกคิดเอง ไม่ใช่ท่องจำ

## 📂 โครงสร้างไฟล์

```
se69-exam-mock/
├── ข้อสอบ-SE69-Midterm-จำลอง.pdf   # โจทย์ + เงื่อนไข + ภาคผนวก (อ่านอันนี้ก่อน)
├── starter/                        # โค้ดเริ่มต้น (ใช้ตอนทำข้อสอบ)
│   ├── schema.sql                  # โครง ข้อ 1 (มี TODO)
│   ├── main.go                     # โครง ข้อ 2 (เชื่อม DB ให้แล้ว, model เป็น TODO)
│   ├── go.mod
│   └── reference/
│       ├── gorm-tags.md            # ตาราง GORM tags (ภาคผนวก ก)
│       └── commands.md             # คำสั่งที่จำเป็น (ภาคผนวก ข)
└── solution/                       # เฉลย (ดูหลังลองทำเองแล้ว)
    ├── schema.sql
    ├── main.go
    └── constraint-test.sql
```

## 🎯 โจทย์ 3 ข้อ (10 คะแนน)

1. **(4) สร้างตารางด้วย SQL** — เขียน `schema.sql` สร้าง 3 ตารางในฐาน `rental_sql`
2. **(4) สร้างตารางด้วย GORM** — เติม model ใน `main.go` ให้ AutoMigrate สร้างตารางเหมือนข้อ 1 ในฐาน `rental_gorm`
3. **(2) ทดสอบ constraint** — บันทึกคำสั่งที่โดนปฏิเสธ + ตรวจ DEFAULT และ CASCADE ลง `constraint-test.sql`

## ▶️ วิธีฝึก

1. อ่าน `ข้อสอบ-SE69-Midterm-จำลอง.pdf`
2. เตรียมฐานข้อมูล 2 ฐาน: `rental_sql` และ `rental_gorm`
3. ลองทำในโฟลเดอร์ `starter/` ด้วยตัวเองก่อน โดย**ไม่เปิดเฉลย**
4. ติดตรงไหนค่อยเปิด `solution/` เทียบ

## 🗄️ เตรียมฐานข้อมูล (เลือกอย่างใดอย่างหนึ่ง)

**ก) Docker** (เหมือนห้องสอบ)
```bash
docker run --name se69-exam-db -e POSTGRES_USER=lab -e POSTGRES_PASSWORD=lab1234 -e POSTGRES_DB=rental_sql -p 5432:5432 -d postgres:17
docker exec -it se69-exam-db psql -U lab -d rental_sql -c "CREATE DATABASE rental_gorm;"
```

**ข) PostgreSQL ติดตั้งตรงบนเครื่อง**
```
psql -U postgres -c "CREATE USER lab WITH PASSWORD 'lab1234';"
psql -U postgres -c "CREATE DATABASE rental_sql OWNER lab;"
psql -U postgres -c "CREATE DATABASE rental_gorm OWNER lab;"
```

---
*ชุดฝึกเพื่อการเรียนรู้ส่วนบุคคล*
