package tgbotapi

import (
	"strconv"
	"strings"
	"unicode"
)

// ValidateToken checks if the given token is a valid BOT token and returns ID
//
// The BOT token consists of two parts like
// BOT_ID:BOT_ALPHANUMERIC_PART
// The BOT_ID is 8 to 10 digit long and the BOT_ALPHANUMERIC_PART has a length of 35 characters.
// So the total length is 43 to 45 characters.
func ValidateToken(token string) (bool, int64) {
	// Check if the token is of correct length
	if len(token) != 43 && len(token) != 44 && len(token) != 45 {
		return false, 0
	}

	// Split the token into two parts
	parts := strings.Split(token, ":")
	if len(parts) != 2 {
		return false, 0
	}

	// Check that the BOT_ID is 8-10 digits long and numeric
	id, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil || id < 10000000 || id > 9999999999 {
		return false, 0
	}

	// Check that the BOT_ALPHANUMERIC_PART has length 35 and contains only valid characters
	botAlphaPart := parts[1]
	if len(botAlphaPart) != 35 {
		return false, 0
	}

	for _, c := range botAlphaPart {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '-' && c != '_' {
			return false, 0
		}
	}

	// If all checks pass, the token is valid and we return the BOT_ID as int64
	return true, id
}
