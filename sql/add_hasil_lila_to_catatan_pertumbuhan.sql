-- Migration: Add hasil_lila column to catatan_pertumbuhan table
-- Purpose: Support LILA (Lingkar Lengan Atas) measurement for stunting prediction
-- Date: 2026-05-04

BEGIN;

-- Add hasil_lila column if not exists
ALTER TABLE catatan_pertumbuhan 
ADD COLUMN IF NOT EXISTS hasil_lila DECIMAL(5,2) DEFAULT 0;

-- Add comment for documentation
COMMENT ON COLUMN catatan_pertumbuhan.hasil_lila IS 'Lingkar Lengan Atas (LILA) dalam cm - digunakan untuk prediksi stunting';

COMMIT;

-- Verification query (run separately after migration):
-- SELECT column_name, data_type, column_default 
-- FROM information_schema.columns 
-- WHERE table_name = 'catatan_pertumbuhan' AND column_name = 'hasil_lila';
