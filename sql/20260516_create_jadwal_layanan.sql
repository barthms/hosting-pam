-- Migration: create jadwal_layanan table

CREATE TABLE IF NOT EXISTS jadwal_layanan (
    id SERIAL PRIMARY KEY,
    posyandu_id INTEGER REFERENCES posyandu(id) ON DELETE SET NULL,
    layanan VARCHAR(255) NOT NULL,
    tanggal DATE,
    waktu_mulai TIME,
    waktu_selesai TIME,
    keterangan TEXT,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_jadwal_layanan_posyandu ON jadwal_layanan(posyandu_id);
CREATE INDEX IF NOT EXISTS idx_jadwal_layanan_tanggal ON jadwal_layanan(tanggal);
