-- Migration: Create Table Puskesmas dan Posyandu
-- Database: PostgreSQL
-- Date: 2026-05-11
-- Description: Tabel untuk manajemen Puskesmas dan Posyandu sesuai design database
-- Safe to run multiple times (idempotent)

BEGIN;

-- 1) Create puskesmas table
CREATE TABLE IF NOT EXISTS puskesmas (
  id BIGSERIAL PRIMARY KEY,
  nama VARCHAR(255) NOT NULL,
  alamat TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);

-- 2) Create posyandu table
CREATE TABLE IF NOT EXISTS posyandu (
  id BIGSERIAL PRIMARY KEY,
  id_puskesmas BIGINT NOT NULL,
  nama VARCHAR(255) NOT NULL,
  alamat TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);

-- 3) Add FK constraint: posyandu -> puskesmas
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM pg_constraint
    WHERE conname = 'fk_posyandu_id_puskesmas'
  ) THEN
    ALTER TABLE posyandu
      ADD CONSTRAINT fk_posyandu_id_puskesmas
      FOREIGN KEY (id_puskesmas)
      REFERENCES puskesmas(id)
      ON UPDATE CASCADE
      ON DELETE RESTRICT;
  END IF;
END $$;

-- 4) Create helpful indexes
CREATE INDEX IF NOT EXISTS ix_puskesmas_deleted_at ON puskesmas(deleted_at);
CREATE INDEX IF NOT EXISTS ix_posyandu_id_puskesmas ON posyandu(id_puskesmas);
CREATE INDEX IF NOT EXISTS ix_posyandu_deleted_at ON posyandu(deleted_at);
CREATE INDEX IF NOT EXISTS ix_posyandu_nama ON posyandu(nama);

-- 5) Update kader table to reference posyandu if needed (jika belum ada)
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.columns
    WHERE table_schema = 'public'
      AND table_name = 'kader'
      AND column_name = 'posyandu_id'
  ) THEN
    ALTER TABLE kader ADD COLUMN IF NOT EXISTS posyandu_id BIGINT;
    
    ALTER TABLE kader
      ADD CONSTRAINT fk_kader_posyandu_id
      FOREIGN KEY (posyandu_id)
      REFERENCES posyandu(id)
      ON UPDATE CASCADE
      ON DELETE SET NULL;
  END IF;
END $$;

-- 6) Optional: Insert sample puskesmas data
-- Uncomment untuk menambahkan data contoh
-- INSERT INTO puskesmas (nama, alamat, created_at, updated_at)
-- VALUES 
--   ('Puskesmas Desa Sentosa', 'Jl. Kesehatan No. 1', NOW(), NOW()),
--   ('Puskesmas Desa Maju', 'Jl. Raya No. 5', NOW(), NOW())
-- ON CONFLICT DO NOTHING;

COMMIT;
