package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const GenesisHash = "0000000000000000000000000000000000000000000000000000000000000000"

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

func appendEntry(actor, action, document, documentHash string) error {
	entries, err := loadLog()
	if err != nil {
		return err
	}
	entry := AuditEntry{
		ID:           len(entries) + 1,
		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Actor:        actor,
		Action:       action,
		Document:     document,
		DocumentHash: documentHash,
	}

	if len(entries) == 0 {
		entry.PreviousHash = GenesisHash
	} else {
		entry.PreviousHash = entries[len(entries)-1].EntryHash
	}
	entry.EntryHash = computeEntryHash(entry)
	entries = append(entries, entry)
	data, err := json.MarshalIndent(entries, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile("audit_log.json", data, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func verifyChain() error {
	entries, err := loadLog()
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		fmt.Println("Log is empty. Nothing to verify.")
		return nil
	}

	for i, entry := range entries {

		computed := computeEntryHash(entry)
		if computed != entry.EntryHash {
			return fmt.Errorf("CHAIN BROKEN at entry %d: entry hash mismatch", entry.ID)
		}

		if i == 0 {
			if entry.PreviousHash != GenesisHash {
				return fmt.Errorf("CHAIN BROKEN at entry %d: invalid genesis hash", entry.ID)
			}
		} else {
			if entry.PreviousHash != entries[i-1].EntryHash {
				return fmt.Errorf("CHAIN BROKEN at entry %d: previous hash mismatch", entry.ID)
			}
		}
	}

	fmt.Println("Chain verified. All", len(entries), "entries are intact.")
	return nil
}
