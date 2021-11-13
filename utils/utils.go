package utils

/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */
import (
	"encoding/json"
	. "fmt"
	"log"
)

// type Void struct{}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func ObjectToJSON(argStr string, objectToConvert interface{}) []byte {
	// ™ ____________
	objectToJSONFmt, err := json.MarshalIndent(
		objectToConvert, "", "  ",
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
	
	Printf("%s: %s\n", argStr, objectToJSONFmt)
	return objectToJSONFmt
}

/* ☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰ */
