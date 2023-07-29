package utils

import (
	"fmt"
	"time"
)

func GenerateInvoiceCode() string {
	uniqueCode := time.Now().Unix()
	return fmt.Sprintf("INV-%d", uniqueCode)
}
