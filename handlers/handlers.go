package handlers
import (
	"fmt"
	"encoding/json"
    "log"
    "net/http"
	"structs"
)

func PostTemperature(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Handling Temperature Post")
    decoder :=json.NewDecoder(r.Body)
    var temp structs.TempData
    err := decoder.Decode(&temp)
      if err != nil {
        panic(fmt.Sprintf("%v", &temp))
    }
    log.Println(temp.Temperature)
}