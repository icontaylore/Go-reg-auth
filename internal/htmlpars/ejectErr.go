package htmlpars

import (
	"fmt"
	"log"
)

func EjectErr(data map[string]interface{}) string {
	if err, ok := data["Error"].(string); ok {
		log.Println(err)
		return err
	}
	fmt.Println("not err")
	return ""
}
