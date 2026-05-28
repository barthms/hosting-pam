-- Migration: tabel warna_tinja_anak
-- Jalankan di database PostgreSQL/Supabase

CREATE TABLE IF NOT EXISTS warna_tinja_anak (
    id BIGSERIAL PRIMARY KEY,
    anak_id BIGINT NOT NULL,
    periode_key VARCHAR(20) NOT NULL,
    periode_label VARCHAR(30) NOT NULL,
    tanggal_catat DATE NOT NULL,
    nomor_warna INTEGER NOT NULL CHECK (nomor_warna BETWEEN 1 AND 7),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,

    CONSTRAINT fk_warna_tinja_anak_anak
        FOREIGN KEY (anak_id)
        REFERENCES anak(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT uq_warna_tinja_anak_periode
        UNIQUE (anak_id, periode_key)
);

CREATE INDEX IF NOT EXISTS idx_warna_tinja_anak_deleted_at
    ON warna_tinja_anak(deleted_at);

CREATE INDEX IF NOT EXISTS idx_warna_tinja_anak_anak_id
    ON warna_tinja_anak(anak_id);
