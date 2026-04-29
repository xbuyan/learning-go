package main

import (
	"encoding/json"
	"os"
)

func loadLog() ([]AuditEntry, error) {
	data, err := os.ReadFile("audit_log.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []AuditEntry{}, nil
		}
		return nil, err
	}

	var entries []AuditEntry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
