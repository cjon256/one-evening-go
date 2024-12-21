package main

import "strings"

func CountCreatedEvents(events []string) int {
	count := 0
	for _, e := range events {
		if strings.HasSuffix(e, "_created") {
			count += 1
		} else if strings.HasSuffix(e, "_deleted") {
			break
		} else {
			// this seems pointless... how exactly did they think this should be structured?
			continue
		}
	}
	return count
}

func main() {
	events := []string{
		"product_created",
		"product_updated",
		"product_assigned",
		"order_created",
		"order_updated",
		"client_created",
		"client_updated",
		"client_refreshed",
		"client_deleted",
		"order_updated",
	}

	CountCreatedEvents(events)
}
