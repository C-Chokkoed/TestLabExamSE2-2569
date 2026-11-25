# ภาคผนวก ก — ตาราง GORM Tags ที่จำเป็น

| ต้องการ (SQL) | GORM struct tag |
|---|---|
| PRIMARY KEY (นับเองอัตโนมัติ) | `gorm:"primaryKey"` |
| VARCHAR(n) | `gorm:"size:n"` |
| NOT NULL | `gorm:"not null"` |
| UNIQUE | `gorm:"unique"` |
| DEFAULT ค่า | `gorm:"default:ค่า"` |
| CHECK เงื่อนไข | `gorm:"check:ชื่อ,เงื่อนไข"` |
| Foreign key + ON DELETE (Belongs To) | `gorm:"constraint:OnDelete:CASCADE"` |
| คอลัมน์ที่เว้นว่าง (NULL) ได้ | ใช้ชนิด pointer เช่น `*time.Time` |
| ชนิดให้ได้ bigint | ใช้ `uint` (ไม่ใช่ `int`) |

รวมหลาย tag ในช่องเดียวได้ด้วยเครื่องหมาย `;`
เช่น `gorm:"size:10;not null;unique"`

## รูปแบบ Belongs To (foreign key)

```go
type Rental struct {
    ID       uint
    MemberID uint   `gorm:"not null"`
    Member   Member `gorm:"constraint:OnDelete:CASCADE"`
    BikeID   uint   `gorm:"not null"`
    Bike     Bike   `gorm:"constraint:OnDelete:RESTRICT"`
}
```
ฟิลด์ `MemberID` (ชื่อ = ชื่อ struct + ID) คือคอลัมน์ foreign key จริง
ส่วนฟิลด์ `Member` (ชนิดเป็น struct) เป็นตัวบอกความสัมพันธ์ ไม่ได้สร้างคอลัมน์
