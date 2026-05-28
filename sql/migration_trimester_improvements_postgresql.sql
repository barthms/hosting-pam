-- Migration: Add CASCADE DELETE dan GambarUSG field untuk Trimester 1 & 3
-- Database: PostgreSQL
-- Date: 2026-05-06
-- Description: 
--   1. Add GambarUSG field untuk menyimpan gambar USG (Base64 encoded)
--   2. Add CASCADE DELETE untuk catatan yang terkait dengan pemeriksaan trimester
--   3. Hilangkan data redundan (tanggal akan di-set otomatis ke hari ini di backend)

-- =============== TRIMESTER 1 ===============

-- 1. Add GambarUSG column to pemeriksaan_dokter_trimester_1
ALTER TABLE pemeriksaan_dokter_trimester_1
ADD COLUMN IF NOT EXISTS gambar_usg TEXT NULL;

-- 2. Update foreign key untuk catatan_pelayanan_trimester_1 dengan CASCADE DELETE
-- Drop existing constraint if exists
ALTER TABLE catatan_pelayanan_trimester_1
DROP CONSTRAINT IF EXISTS catatan_pelayanan_trimester_1_kehamilan_id_fkey;

-- Add new constraint with CASCADE DELETE
ALTER TABLE catatan_pelayanan_trimester_1
ADD CONSTRAINT catatan_pelayanan_trimester_1_kehamilan_id_fkey
FOREIGN KEY (kehamilan_id) REFERENCES kehamilan(id) ON DELETE CASCADE;

-- =============== TRIMESTER 3 ===============

-- 1. Add GambarUSG column to pemeriksaan_dokter_trimester_3
ALTER TABLE pemeriksaan_dokter_trimester_3
ADD COLUMN IF NOT EXISTS gambar_usg TEXT NULL;

-- 2. Update foreign key untuk catatan_pelayanan_trimester_3 dengan CASCADE DELETE
-- Drop existing constraint if exists
ALTER TABLE catatan_pelayanan_trimester_3
DROP CONSTRAINT IF EXISTS catatan_pelayanan_trimester_3_kehamilan_id_fkey;

-- Add new constraint with CASCADE DELETE
ALTER TABLE catatan_pelayanan_trimester_3
ADD CONSTRAINT catatan_pelayanan_trimester_3_kehamilan_id_fkey
FOREIGN KEY (kehamilan_id) REFERENCES kehamilan(id) ON DELETE CASCADE;

-- =============== Verification ===============
-- Verify the columns were added
SELECT table_name, column_name, data_type
FROM information_schema.columns
WHERE table_schema = 'public'
  AND table_name IN ('pemeriksaan_dokter_trimester_1', 'pemeriksaan_dokter_trimester_3')
  AND column_name = 'gambar_usg';

-- Check the foreign key constraints
SELECT
  tc.constraint_name,
  tc.table_name,
  kcu.column_name,
  ccu.table_name AS referenced_table_name
FROM information_schema.table_constraints AS tc
  JOIN information_schema.key_column_usage AS kcu
    ON tc.constraint_name = kcu.constraint_name
    AND tc.table_schema = kcu.table_schema
  JOIN information_schema.constraint_column_usage AS ccu
    ON ccu.constraint_name = tc.constraint_name
    AND ccu.table_schema = tc.table_schema
WHERE tc.constraint_type = 'FOREIGN KEY'
  AND tc.table_schema = 'public'
  AND tc.table_name IN ('catatan_pelayanan_trimester_1', 'catatan_pelayanan_trimester_3')
  AND kcu.column_name = 'kehamilan_id';
