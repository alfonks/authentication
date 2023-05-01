package debug

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(raw interface{}, forceStop bool) {
	data, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		log.Printf("error pretty print: %v", raw)
		return
	}

	printData := fmt.Sprintf(">>>>> %v <<<<<", string(data))

	if forceStop {
		log.Fatalln(printData)
	}
	log.Println(printData)
}
