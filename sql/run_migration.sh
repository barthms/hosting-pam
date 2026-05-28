#!/bin/bash

# Migration Runner - PostgreSQL
# Description: Run the trimester improvements migration
# Usage: ./run_migration.sh

set -e

echo "=========================================="
echo "Running Migration: Trimester Improvements"
echo "=========================================="
echo ""

# Load environment variables if .env file exists
if [ -f ".env" ]; then
    echo "📋 Loading environment from .env file..."
    export $(cat .env | grep -v '^#' | xargs)
else
    echo "⚠️  No .env file found, using system environment variables"
fi

# Get database configuration from environment
DB_USER=${DB_POSTGRES_USER:-postgres}
DB_PASSWORD=${DB_POSTGRES_PASSWORD}
DB_HOST=${DB_POSTGRES_HOST:-localhost}
DB_PORT=${DB_POSTGRES_PORT:-5432}
DB_NAME=${DB_POSTGRES_NAME}
DB_SCHEMA=${DB_POSTGRES_SCHEMA:-public}

echo ""
echo "🔗 Database Configuration:"
echo "   Host: $DB_HOST"
echo "   Port: $DB_PORT"
echo "   User: $DB_USER"
echo "   Database: $DB_NAME"
echo "   Schema: $DB_SCHEMA"
echo ""

# Validate required variables
if [ -z "$DB_NAME" ]; then
    echo "❌ ERROR: DB_POSTGRES_NAME not set in environment"
    exit 1
fi

if [ -z "$DB_PASSWORD" ]; then
    echo "❌ ERROR: DB_POSTGRES_PASSWORD not set in environment"
    exit 1
fi

# Set password for psql
export PGPASSWORD="$DB_PASSWORD"

# Run migration
echo "🚀 Running migration..."
psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -f "sql/migration_trimester_improvements_postgresql.sql"

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ Migration completed successfully!"
    echo ""
    echo "Verification queries:"
    echo ""
    
    # Run verification
    psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" << EOF
-- Verify gambar_usg columns
SELECT 
    table_name, 
    column_name, 
    data_type
FROM information_schema.columns
WHERE table_schema = 'public'
  AND table_name IN ('pemeriksaan_dokter_trimester_1', 'pemeriksaan_dokter_trimester_3')
  AND column_name = 'gambar_usg'
ORDER BY table_name;

-- Verify foreign key constraints
SELECT
    tc.constraint_name,
    tc.table_name,
    kcu.column_name,
    'CASCADE' as delete_action
FROM information_schema.table_constraints AS tc
  JOIN information_schema.key_column_usage AS kcu
    ON tc.constraint_name = kcu.constraint_name
    AND tc.table_schema = kcu.table_schema
WHERE tc.constraint_type = 'FOREIGN KEY'
  AND tc.table_schema = 'public'
  AND tc.table_name IN ('catatan_pelayanan_trimester_1', 'catatan_pelayanan_trimester_3')
  AND kcu.column_name = 'kehamilan_id'
ORDER BY tc.table_name;
EOF
else
    echo ""
    echo "❌ Migration failed!"
    exit 1
fi

unset PGPASSWORD
