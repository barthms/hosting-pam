# Migration Runner - PostgreSQL (Windows PowerShell)
# Description: Run the trimester improvements migration
# Usage: .\run_migration.ps1

$ErrorActionPreference = "Stop"

Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "Running Migration: Trimester Improvements" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

# Get database configuration from environment variables
$DB_USER = $env:DB_POSTGRES_USER ? $env:DB_POSTGRES_USER : "postgres"
$DB_PASSWORD = $env:DB_POSTGRES_PASSWORD
$DB_HOST = $env:DB_POSTGRES_HOST ? $env:DB_POSTGRES_HOST : "localhost"
$DB_PORT = $env:DB_POSTGRES_PORT ? $env:DB_POSTGRES_PORT : "5432"
$DB_NAME = $env:DB_POSTGRES_NAME
$DB_SCHEMA = $env:DB_POSTGRES_SCHEMA ? $env:DB_POSTGRES_SCHEMA : "public"

Write-Host "🔗 Database Configuration:" -ForegroundColor Yellow
Write-Host "   Host: $DB_HOST"
Write-Host "   Port: $DB_PORT"
Write-Host "   User: $DB_USER"
Write-Host "   Database: $DB_NAME"
Write-Host "   Schema: $DB_SCHEMA"
Write-Host ""

# Validate required variables
if (-not $DB_NAME) {
    Write-Host "❌ ERROR: DB_POSTGRES_NAME not set in environment" -ForegroundColor Red
    exit 1
}

if (-not $DB_PASSWORD) {
    Write-Host "❌ ERROR: DB_POSTGRES_PASSWORD not set in environment" -ForegroundColor Red
    exit 1
}

# Build connection string
$connectionString = "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}"

# Run migration
Write-Host "🚀 Running migration..." -ForegroundColor Green
Write-Host ""

# Check if psql is available
$psqlPath = Get-Command psql -ErrorAction SilentlyContinue
if (-not $psqlPath) {
    Write-Host "❌ ERROR: psql not found in PATH" -ForegroundColor Red
    Write-Host "   Please install PostgreSQL client tools or add to PATH" -ForegroundColor Yellow
    exit 1
}

# Run the migration file
$env:PGPASSWORD = $DB_PASSWORD
& psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "sql/migration_trimester_improvements_postgresql.sql"

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "✅ Migration completed successfully!" -ForegroundColor Green
    Write-Host ""
    Write-Host "📊 Verification results shown above" -ForegroundColor Green
    Write-Host ""
} else {
    Write-Host ""
    Write-Host "❌ Migration failed with exit code: $LASTEXITCODE" -ForegroundColor Red
    exit 1
}

Remove-Item Env:\PGPASSWORD -ErrorAction SilentlyContinue
