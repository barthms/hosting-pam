-- Migration: Fix Ibu table unique constraint
-- Purpose: Remove unique constraint on penduduk_id to allow multiple pregnancies per person
-- Safe to run multiple times (idempotent)

BEGIN;

-- Drop unique index if exists (PostgreSQL)
DROP INDEX IF EXISTS ux_ibu_penduduk_id CASCADE;

-- Drop unique constraint if exists (alternative method)
ALTER TABLE IF EXISTS ibu DROP CONSTRAINT IF EXISTS uq_ibu_penduduk_id CASCADE;

-- Create regular index if not exists (for better query performance)
CREATE INDEX IF NOT EXISTS idx_ibu_penduduk_id ON ibu(penduduk_id);

-- Create index on status_kehamilan for filtering
CREATE INDEX IF NOT EXISTS idx_ibu_status_kehamilan ON ibu(status_kehamilan);

-- Create composite index on penduduk_id and created_at for efficient queries
CREATE INDEX IF NOT EXISTS idx_ibu_penduduk_created ON ibu(penduduk_id, created_at DESC);

-- Create index on is_deleted for soft delete queries
CREATE INDEX IF NOT EXISTS idx_ibu_is_deleted ON ibu(is_deleted);

-- Add timestamp update trigger if not exists
CREATE OR REPLACE FUNCTION set_ibu_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_ibu_set_updated_at ON ibu;

CREATE TRIGGER trg_ibu_set_updated_at
BEFORE UPDATE ON ibu
FOR EACH ROW
EXECUTE FUNCTION set_ibu_updated_at();

COMMIT;
