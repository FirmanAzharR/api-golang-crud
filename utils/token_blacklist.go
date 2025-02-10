package utils

import "sync"

var tokenBlacklist = make(map[string]bool)
var lock = sync.Mutex{}

// Tambahkan token ke blacklist
func BlacklistToken(token string) {
	lock.Lock()
	defer lock.Unlock()
	tokenBlacklist[token] = true
}

// Periksa apakah token masuk blacklist
func IsTokenBlacklisted(token string) bool {
	lock.Lock()
	defer lock.Unlock()
	return tokenBlacklist[token]
}
