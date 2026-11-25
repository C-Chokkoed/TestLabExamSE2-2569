# ภาคผนวก ข — คำสั่งที่จำเป็น

> แทน `<container>` ด้วยชื่อกล่องจริงที่ได้จาก `docker ps` (คอลัมน์ NAMES)

| ต้องการ | คำสั่ง |
|---|---|
| ดู container ที่เปิดอยู่ | `docker ps` |
| ดู container ทั้งหมด (รวมที่ปิด) | `docker ps -a` |
| เปิด container | `docker start <container>` |
| รันไฟล์ .sql เข้าฐาน (PowerShell) | `Get-Content schema.sql \| docker exec -i <container> psql -U lab -d rental_sql` |
| เข้าโหมดพิมพ์ SQL โต้ตอบ | `docker exec -it <container> psql -U lab -d rental_sql` |
| สั่ง SQL เดียวแล้วออก | `docker exec -it <container> psql -U lab -d rental_sql -c "SELECT * FROM bikes;"` |
| ดูโครงสร้างตาราง | `\d rentals` |
| ดูรายชื่อตารางทั้งหมด | `\dt` |
| ออกจาก psql | `\q` |
| เริ่มโปรเจกต์ Go | `go mod init se69-exam` |
| ติดตั้ง GORM | `go get gorm.io/gorm gorm.io/driver/postgres` |
| รันโปรแกรม Go | `go run .` |

## ถ้าแก้ model แล้วตารางไม่เปลี่ยน (AutoMigrate ไม่ลบของเดิม)

ลบตารางในฐาน rental_gorm ทั้งหมดก่อน แล้วรัน `go run .` ใหม่:
```
docker exec -it <container> psql -U lab -d rental_gorm -c "DROP TABLE rentals, bikes, members;"
```
