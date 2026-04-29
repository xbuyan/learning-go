package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type AuditEntry struct {
	ID           int    `json:"id"`
	Timestamp    string `json:"timestamp"`
	Actor        string `json:"actor"`
	Action       string `json:"action"`
	Document     string `json:"document"`
	DocumentHash string `json:"document_hash"`
	PreviousHash string `json:"previous_hash"`
	EntryHash    string `json:"entry_hash"`
}

func computeEntryHash(entry AuditEntry) string {
	entry.EntryHash = ""

	data, err := json.Marshal(entry)
	if err != nil {
		return ""
	}
	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}
