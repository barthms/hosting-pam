-- Migration: Rename edukasi_tanda_bahaya_trimester → edukasi_trimester
-- Jalankan di Supabase SQL Editor

BEGIN;

-- Buat tabel baru edukasi_trimester jika belum ada
CREATE TABLE IF NOT EXISTS edukasi_trimester (
    id         SERIAL PRIMARY KEY,
    judul      VARCHAR(255) NOT NULL,
    gambar_url TEXT,
    isi        TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Jika tabel lama edukasi_tanda_bahaya_trimester ada, pindahkan datanya
INSERT INTO edukasi_trimester (id, judul, gambar_url, isi, created_at, updated_at)
SELECT id, judul, gambar_url, isi, created_at, updated_at
FROM edukasi_tanda_bahaya_trimester
ON CONFLICT (id) DO NOTHING;

COMMIT;

-- Verifikasi:
-- SELECT COUNT(*) FROM edukasi_trimester;
