package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sapphirenw/jakeruns.com/src/logger"
)

func WriteObj(w http.ResponseWriter, status int, body any) {
	w.Header().Add("Content-Type", "application/json")
	enc := encode(body)
	if status < 299 {
		logger.Info.Println(string(enc))
	} else {
		logger.Error.Println(string(enc))
	}
	w.WriteHeader(status)
	w.Write(enc)
}

func WriteStr(w http.ResponseWriter, status int, message string, args ...any) {
	w.Header().Add("Content-Type", "application/json")
	enc := encode(map[string]string{"message": fmt.Sprintf(message, args...)})
	if status < 299 {
		logger.Info.Println(string(enc))
	} else {
		logger.Error.Println(string(enc))
	}
	w.WriteHeader(status)
	w.Write(enc)
}

func encode(body any) []byte {
	encoded, err := json.Marshal(body)
	if err != nil {
		return errEnc(fmt.Sprintf("There was an issue encoding the data: %s", err))
	} else {
		return encoded
	}
}

// for when there is an issue with json encoding, so do not rely on json encoding for this error
func errEnc(message string) []byte {
	logger.Critical.Println("there was an issue serializing the json!")
	return []byte(fmt.Sprintf(`{"status": 400, "message": "%s", "body": null}`, message))
}
