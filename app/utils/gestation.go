package utils

import "time"

// HitungUsiaKehamilan menghitung usia kehamilan dalam minggu berdasarkan string HPHT (format YYYY-MM-DD)
func HitungUsiaKehamilan(hpht string) int {
    if hpht == "" {
        return 0
    }
    layout := "2006-01-02"
    hphtDate, err := time.Parse(layout, hpht)
    if err != nil {
        return 0
    }
    today := time.Now()
    today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
    diffDays := int(today.Sub(hphtDate).Hours() / 24)
    if diffDays < 0 {
        return 0
    }
    return diffDays / 7
}

// HitungUsiaKehamilanFromTime menghitung usia kehamilan dari time.Time
func HitungUsiaKehamilanFromTime(hpht time.Time) int {
    if hpht.IsZero() {
        return 0
    }
    today := time.Now()
    today = time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
    diffDays := int(today.Sub(hpht).Hours() / 24)
    if diffDays < 0 {
        return 0
    }
    return diffDays / 7
}

func DetermineTrimester(usia int) string {
    switch {
    case usia <= 12:
        return "TRIMESTER 1"
    case usia <= 27:
        return "TRIMESTER 2"
    default:
        return "TRIMESTER 3"  // tetap Trimester 3 meskipun >40
    }
}