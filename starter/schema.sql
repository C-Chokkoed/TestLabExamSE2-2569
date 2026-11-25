-- ============================================================
-- ข้อ 1: สร้างตารางด้วย SQL  (ฐานข้อมูล rental_sql)
-- เขียนคำสั่ง CREATE TABLE ทั้ง 3 ตารางตามข้อกำหนดของ Schema ในข้อสอบ
-- คำใบ้: สร้าง members และ bikes ให้เสร็จก่อน rentals
--        เพราะ rentals อ้างถึงทั้งสองตาราง
-- ============================================================

-- TODO: ตาราง members
--   id BIGSERIAL PRIMARY KEY
--   member_code VARCHAR(10) NOT NULL UNIQUE
--   full_name   VARCHAR(100) NOT NULL
--   phone       VARCHAR(20) NOT NULL UNIQUE


-- TODO: ตาราง bikes
--   id BIGSERIAL PRIMARY KEY
--   code    VARCHAR(10) NOT NULL UNIQUE
--   model   VARCHAR(100) NOT NULL
--   deposit INTEGER NOT NULL DEFAULT 500 CHECK (deposit >= 0)


-- TODO: ตาราง rentals
--   id BIGSERIAL PRIMARY KEY
--   member_id   -> members(id)  ON DELETE CASCADE
--   bike_id     -> bikes(id)    ON DELETE RESTRICT
--   status      DEFAULT 'active', เฉพาะ 'active' / 'returned'
--   start_at    NOT NULL
--   due_at      NOT NULL, ต้อง > start_at
--   returned_at เว้นว่างได้
