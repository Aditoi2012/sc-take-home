package folders

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

var TokenizedPagination = make(map[string]PaginatedResult)

/* 
	Explanation: I deicded to use the existing functions that were created in folders.go.
	My main thought process behind this was to separate the pagination code and the 
	method to get documents based off the orgID. Currently, when main.go is run, based
	off the page size, every page would have the name and at the end there would be the token.

	The overall logic behind this code was to divide the lfolders into pages based off the size.
	A global map is then used as tokens  for keys and then these are used to print it out.
*/

// GenerateRandomToken creates a random string of the specified length
func GenerateRandomToken(length int) (string, error) {
	numBytes := (length * 6) / 8
	if (length*6)%8 != 0 {
		numBytes++
	}

	randomBytes := make([]byte, numBytes)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(randomBytes)
	token = strings.ReplaceAll(token, "=", "")

	return token[:length], nil
}

func Request(token string) (PaginatedResult, bool) {
	result, exists := TokenizedPagination[token]
	return result, exists
}

// PaginateData splits the data into pages and generates tokens for each page
func PaginateData(data []*Folder, pageSize int) (error) {
	numPages := (len(data) + pageSize - 1) / pageSize
	var prevToken string
	prevToken = ""

	for i := 0; i < numPages; i++ {
		start := i * pageSize
		end := start + pageSize
		if end > len(data) {
			end = len(data)
		}

		pageData := data[start:end]

		token, err := GenerateRandomToken(5)
		if err != nil {
			return err
		}

		if i == numPages-1 {
			TokenizedPagination[prevToken] = PaginatedResult{
				Data:  pageData,
				Token: "",
			}
		} else {
			TokenizedPagination[prevToken] = PaginatedResult{
				Data:  pageData,
				Token: token,
			}
			prevToken = token
		}
	}

	return nil
}