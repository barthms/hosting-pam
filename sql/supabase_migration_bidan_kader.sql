-- Supabase SQL Editor migration
-- Modul: Dashboard Desa - Manajemen Bidan & Kader
-- Safe to run multiple times (idempotent)

BEGIN;

ALTER TABLE IF EXISTS penduduk
  ADD COLUMN IF NOT EXISTS kecamatan TEXT,
  ADD COLUMN IF NOT EXISTS desa TEXT;

-- 1) Ensure required roles exist
INSERT INTO roles (name, created_at, updated_at)
VALUES
  ('Admin', NOW(), NOW()),
  ('Bidan', NOW(), NOW()),
  ('Kader', NOW(), NOW())
ON CONFLICT (name) DO NOTHING;

-- 2) Ensure pengguna has penduduk_id
ALTER TABLE pengguna
  ADD COLUMN IF NOT EXISTS penduduk_id BIGINT;

-- FK pengguna -> penduduk (id)
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM pg_constraint
    WHERE conname = 'fk_pengguna_penduduk_id'
  ) THEN
    ALTER TABLE pengguna
      ADD CONSTRAINT fk_pengguna_penduduk_id
      FOREIGN KEY (penduduk_id)
      REFERENCES penduduk(id)
      ON UPDATE CASCADE
      ON DELETE SET NULL;
  END IF;
END $$;

-- Enforce 1 penduduk : 1 akun (for non-null penduduk_id)
CREATE UNIQUE INDEX IF NOT EXISTS ux_pengguna_penduduk_id
  ON pengguna(penduduk_id)
  WHERE penduduk_id IS NOT NULL;

-- 3) Table bidan
CREATE TABLE IF NOT EXISTS bidan (
  id BIGSERIAL PRIMARY KEY,
  penduduk_id BIGINT NOT NULL,
  no_str VARCHAR(100),
  no_sipb VARCHAR(100),
  desa TEXT,
  status VARCHAR(20) NOT NULL DEFAULT 'aktif',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);

ALTER TABLE IF EXISTS bidan
  ADD COLUMN IF NOT EXISTS desa TEXT;

-- FK bidan -> penduduk
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM pg_constraint
    WHERE conname = 'fk_bidan_penduduk_id'
  ) THEN
    ALTER TABLE bidan
      ADD CONSTRAINT fk_bidan_penduduk_id
      FOREIGN KEY (penduduk_id)
      REFERENCES penduduk(id)
      ON UPDATE CASCADE
      ON DELETE RESTRICT;
  END IF;
END $$;

-- Status constraint
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM pg_constraint
    WHERE conname = 'ck_bidan_status'
  ) THEN
    ALTER TABLE bidan
      ADD CONSTRAINT ck_bidan_status
      CHECK (status IN ('aktif', 'nonaktif'));
  END IF;
END $$;

-- Uniqueness: one penduduk can be one bidan only
CREATE UNIQUE INDEX IF NOT EXISTS ux_bidan_penduduk_id ON bidan(penduduk_id);

-- Helpful indexes
CREATE INDEX IF NOT EXISTS ix_bidan_desa ON bidan(desa);
CREATE INDEX IF NOT EXISTS ix_bidan_status ON bidan(status);
CREATE INDEX IF NOT EXISTS ix_bidan_deleted_at ON bidan(deleted_at);

-- 4) Table kader
CREATE TABLE IF NOT EXISTS kader (
  id BIGSERIAL PRIMARY KEY,
  penduduk_id BIGINT NOT NULL,
  desa TEXT,
  posyandu_id BIGINT,
  status VARCHAR(20) NOT NULL DEFAULT 'aktif',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);

ALTER TABLE IF EXISTS kader
  ADD COLUMN IF NOT EXISTS desa TEXT;

-- FK kader -> penduduk
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM pg_constraint
    WHERE conname = 'fk_kader_penduduk_id'
  ) THEN
    ALTER TABLE kader
      ADD CONSTRAINT fk_kader_penduduk_id
      FOREIGN KEY (penduduk_id)
      REFERENCES penduduk(id)
      ON UPDATE CASCADE
      ON DELETE RESTRICT;
  END IF;
END $$;

-- Status constraint
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM pg_constraint
    WHERE conname = 'ck_kader_status'
  ) THEN
    ALTER TABLE kader
      ADD CONSTRAINT ck_kader_status
      CHECK (status IN ('aktif', 'nonaktif'));
  END IF;
END $$;

-- Uniqueness: one penduduk can be one kader only
CREATE UNIQUE INDEX IF NOT EXISTS ux_kader_penduduk_id ON kader(penduduk_id);

-- Helpful indexes
CREATE INDEX IF NOT EXISTS ix_kader_desa ON kader(desa);
CREATE INDEX IF NOT EXISTS ix_kader_posyandu_id ON kader(posyandu_id);
CREATE INDEX IF NOT EXISTS ix_kader_status ON kader(status);
CREATE INDEX IF NOT EXISTS ix_kader_deleted_at ON kader(deleted_at);

-- 5) Optional timestamp update trigger function
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_trigger WHERE tgname = 'trg_bidan_set_updated_at'
  ) THEN
    CREATE TRIGGER trg_bidan_set_updated_at
    BEFORE UPDATE ON bidan
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1 FROM pg_trigger WHERE tgname = 'trg_kader_set_updated_at'
  ) THEN
    CREATE TRIGGER trg_kader_set_updated_at
    BEFORE UPDATE ON kader
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();
  END IF;
END $$;

UPDATE bidan b
SET desa = COALESCE(NULLIF(b.desa, ''), p.desa)
FROM penduduk p
WHERE p.id = b.penduduk_id;

UPDATE kader k
SET desa = COALESCE(NULLIF(k.desa, ''), p.desa)
FROM penduduk p
WHERE p.id = k.penduduk_id;

COMMIT;
