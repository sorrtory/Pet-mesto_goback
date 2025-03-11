package common

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

func GenerateID() string {
	// Create a byte slice of 8 random bytes
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes) // Generate random bytes

	// Get the current Unix timestamp in seconds
	timestamp := time.Now().Unix()

	// Convert the timestamp to a 4-byte slice
	timestampBytes := make([]byte, 4)
	timestampBytes[0] = byte(timestamp >> 24)
	timestampBytes[1] = byte(timestamp >> 16)
	timestampBytes[2] = byte(timestamp >> 8)
	timestampBytes[3] = byte(timestamp)

	// Combine timestamp and random bytes
	combined := append(timestampBytes, randomBytes...)

	// Return the hex-encoded string (24 characters)
	return hex.EncodeToString(combined)
}
