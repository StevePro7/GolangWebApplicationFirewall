package main

// https://golangbyexample.com/generate-uuid-guid-golang/
import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func main() {
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}
