-- Migration: Add CASCADE DELETE dan GambarUSG field untuk Trimester 1 & 3
-- Date: 2026-05-06
-- Description: 
--   1. Add GambarUSG field untuk menyimpan gambar USG (Base64 encoded)
--   2. Add CASCADE DELETE untuk catatan yang terkait dengan pemeriksaan trimester
--   3. Hilangkan data redundan (tanggal akan di-set otomatis ke hari ini di backend)

-- =============== TRIMESTER 1 ===============

-- 1. Add GambarUSG column to pemeriksaan_dokter_trimester_1
ALTER TABLE pemeriksaan_dokter_trimester_1
ADD COLUMN gambar_usg LONGTEXT NULL AFTER konsep_anamnesa_pemeriksaan;

-- 2. Update foreign key untuk catatan_pelayanan_trimester_1 dengan CASCADE DELETE
-- Drop existing constraint
ALTER TABLE catatan_pelayanan_trimester_1
DROP CONSTRAINT IF EXISTS catatan_pelayanan_trimester_1_ibfk_1;

-- Add new constraint with CASCADE DELETE
ALTER TABLE catatan_pelayanan_trimester_1
ADD CONSTRAINT catatan_pelayanan_trimester_1_ibfk_1
FOREIGN KEY (kehamilan_id) REFERENCES kehamilan(id) ON DELETE CASCADE;

-- =============== TRIMESTER 3 ===============

-- 1. Add GambarUSG column to pemeriksaan_dokter_trimester_3
ALTER TABLE pemeriksaan_dokter_trimester_3
ADD COLUMN gambar_usg LONGTEXT NULL AFTER konsep_anamnesa_pemeriksaan;

-- 2. Update foreign key untuk catatan_pelayanan_trimester_3 dengan CASCADE DELETE
-- Drop existing constraint
ALTER TABLE catatan_pelayanan_trimester_3
DROP CONSTRAINT IF EXISTS catatan_pelayanan_trimester_3_ibfk_1;

-- Add new constraint with CASCADE DELETE
ALTER TABLE catatan_pelayanan_trimester_3
ADD CONSTRAINT catatan_pelayanan_trimester_3_ibfk_1
FOREIGN KEY (kehamilan_id) REFERENCES kehamilan(id) ON DELETE CASCADE;

-- =============== Verification ===============
-- Verify the changes were applied
SELECT TABLE_NAME, COLUMN_NAME 
FROM INFORMATION_SCHEMA.COLUMNS 
WHERE TABLE_SCHEMA = DATABASE() 
  AND TABLE_NAME IN ('pemeriksaan_dokter_trimester_1', 'pemeriksaan_dokter_trimester_3')
  AND COLUMN_NAME = 'gambar_usg';

-- Check the foreign key constraints
SELECT CONSTRAINT_NAME, TABLE_NAME, COLUMN_NAME, REFERENCED_TABLE_NAME
FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE
WHERE TABLE_SCHEMA = DATABASE()
  AND TABLE_NAME IN ('catatan_pelayanan_trimester_1', 'catatan_pelayanan_trimester_3')
  AND COLUMN_NAME = 'kehamilan_id';
