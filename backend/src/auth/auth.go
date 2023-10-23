package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var AuthHost = "http://172.17.0.1:8043"

func extractStringFromHeader(r *http.Request, key string) (string, error) {
	strSlice, ok := r.Header[key]
	if !ok || strSlice == nil || len(strSlice) == 0 {
		return "", errors.New("No " + key + " header provided")
	}
	return strSlice[0], nil
}

func extractIntFromHeader(r *http.Request, key string) (int, error) {
	s, err := extractStringFromHeader(r, key)
	if err != nil {
		return 0, err
	}
	retInt, err := strconv.Atoi(s)
	return retInt, err
}

func extractUintFromHeader(r *http.Request, key string) (uint, error) {
	retInt, err := extractIntFromHeader(r, key)
	return uint(retInt), err
}

func ExtractAuthParamsNoUser(r *http.Request) (int, string, string, string, error) {
	timeInt, err := extractIntFromHeader(r, "Time-Sent")
	if err != nil {
		return 0, "", "", "", err
	}

	// just the url after marktai.com/t9
	// one example "/users/1231/games"
	path := r.URL.String()

	messageHMACString, err := extractStringFromHeader(r, "Hmac")
	if err != nil {
		return 0, "", "", "", err
	}

	encoding, err := extractStringFromHeader(r, "Encoding")
	if err != nil {
		encoding = "hex"
	}

	return timeInt, path, messageHMACString, encoding, nil
}

func ExtractAuthParams(r *http.Request) (uint, int, string, string, string, error) {
	userID, err := extractUintFromHeader(r, "Userid")
	if err != nil {
		return 0, 0, "", "", "", err
	}

	timeInt, path, messageHMACString, encoding, err := ExtractAuthParamsNoUser(r)
	if err != nil {
		return 0, 0, "", "", "", err
	}

	return userID, timeInt, path, messageHMACString, encoding, nil
}

func CheckAuthParams(userID uint, timeInt int, path string, HMAC string, encoding string) (bool, error) {
	req, err := http.NewRequest("POST", AuthHost+"/authHeaders", nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("UserID", fmt.Sprintf("%d", userID))
	req.Header.Set("Time-Sent", fmt.Sprintf("%d", timeInt))
	req.Header.Set("Path", path)
	req.Header.Set("HMAC", HMAC)
	req.Header.Set("Encoding", encoding)
	resp, err := http.DefaultClient.Do(req)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		decoder := json.NewDecoder(resp.Body)
		var parsedJson map[string]string
		err := decoder.Decode(&parsedJson)
		if err != nil {
			return false, errors.New("Auth: " + err.Error() + " in parsing response body (JSON)")
		}

		if errText, ok := parsedJson["Error"]; !ok {
			return false, errors.New("Auth: Error key not present in response body")
		} else {
			return false, errors.New("Auth: " + errText)
		}
	} else {
		return true, nil
	}

}
