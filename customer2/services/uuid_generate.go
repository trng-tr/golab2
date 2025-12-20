package services

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func generateUuid(str string) string {
	var cleaned string = strings.ToUpper(strings.ReplaceAll(str, " ", "-"))
	cleaned = cleaned[:4]
	var shortUuid string = uuid.New().String()[:8]

	return fmt.Sprintf("%s-%s", cleaned, shortUuid)
}
