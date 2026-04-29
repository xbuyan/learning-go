## Real-World Applications

- Hospital medical records — detect post-incident record modification
- Land registry — prove title deed tampering
- Electoral evidence — chain of custody for tallying forms
- Court evidence — unbroken evidence handling records
- NGO document systems — verify institutional record integrity

## Structure

- `main.go` — CLI entry point and command routing
- `entry.go` — AuditEntry struct and hash computation
- `log.go` — append, load, and verify chain operations

## Part Of

auditlog is a standalone tool and a component of the Aegis Anchor
evidence preservation system.

github.com/xbuyan