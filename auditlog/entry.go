package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
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

func hashFile(filepath string) (string, error) {
	content, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer content.Close()

	hasher := sha256.New()
	_, err = io.Copy(hasher, content)
	if err != nil {
		return "", err
	}

	bytes := hasher.Sum(nil)
	return hex.EncodeToString(bytes), nil
}
