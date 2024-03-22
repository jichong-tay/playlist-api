package util

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomName generates a random name
func RandomName() string {
	return RandomString(6)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomImageURL generates a random image URL
func RandomImageURL(width, height int) (string, error) {
	url := fmt.Sprintf("https://source.unsplash.com/random/%dx%d/?food", width, height)
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error get request: %v", err)
	}
	redirectURL := res.Request.URL.String()

	// Close the response body with error handling
	if err := res.Body.Close(); err != nil {
		return "", fmt.Errorf("error closing response body: %v", err)
	}

	return redirectURL, nil
}
