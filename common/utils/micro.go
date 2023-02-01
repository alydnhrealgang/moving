package utils

import (
	"fmt"
	"os"
)

const ServiceNameTemplate = "com.ptit365.warehouse.%s.%s"

func GetMicroServiceName(name string) string {
	return fmt.Sprintf(ServiceNameTemplate, GetNamespace(), name)
}

func GetNamespace() string {
	warehouseNS, ok := os.LookupEnv("WAREHOUSE_NS")
	if !ok || EmptyOrWhiteSpace(warehouseNS) {
		panic("WAREHOUSE_NS_NOT_FOUND")
	}
	return warehouseNS
}
