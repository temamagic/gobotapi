package tgbotapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// ValidateWebAppData validate data received via the Web App
// https://core.telegram.org/bots/webapps#validating-data-received-via-the-web-app
func ValidateWebAppData(token, telegramInitData string) (bool, error) {
	initData, err := url.ParseQuery(telegramInitData)
	if err != nil {
		return false, fmt.Errorf("error parsing data %w", err)
	}

	dataCheckString := make([]string, 0, len(initData))
	for k, v := range initData {
		if k == "hash" {
			continue
		}
		if len(v) > 0 {
			dataCheckString = append(dataCheckString, fmt.Sprintf("%s=%s", k, v[0]))
		}
	}

	sort.Strings(dataCheckString)

	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(token))

	hHash := hmac.New(sha256.New, secret.Sum(nil))
	hHash.Write([]byte(strings.Join(dataCheckString, "\n")))

	hash := hex.EncodeToString(hHash.Sum(nil))

	if initData.Get("hash") != hash {
		return false, errors.New("hash not equal")
	}

	return true, nil
}
