# CLAUDE.md

Leitfaden für KI-unterstützte Arbeit in diesem Repository.

## Projektkontext

Eine kleine Go-REST-API für einen Universitäts-Programmierworkshop.
Greenfield-Go-Backend; `deployments/postgres/` und `deployments/keycloak/`
sind wiederverwendete, unveränderte Infrastruktur aus früheren Abgaben —
diese nicht löschen, ersetzen oder neu erstellen.

## Festgelegter Tech-Stack — nicht davon abweichen

- Go 1.23+
- `github.com/gin-gonic/gin` (REST)
- `gorm.io/gorm` + `gorm.io/driver/postgres` (ORM)
- `github.com/go-playground/validator/v10` via Gin-Binding (Validierung)
- `github.com/coreos/go-oidc/v3` + `golang.org/x/oauth2` (OIDC/JWT)
- `net/http/httptest` + `github.com/stretchr/testify` (Tests)
- `golangci-lint`, `gofmt`, `goimports` (Lint/Formatierung)

Begründung siehe [DECISIONS.md](DECISIONS.md).

## Aktueller Stand

Nur das Grundgerüst: Projektstruktur, Konfigurations-Laden,
GORM-Verbindungsaufbau, Gin-Router mit `/api/public`-/`/api/secured`-Gruppen,
OIDC-Middleware-Platzhalter, ein Health-Check-Integrationstest. Keine
Fachlogik, keine Domänenmodelle, kein vollständiges CRUD.

## Regeln

- Niemals DB-Passwörter, Keycloak-Secrets, Ports oder Tokens hart kodieren —
  alles aus Umgebungsvariablen lesen (siehe `internal/config/`,
  `.env.example`).
- `deployments/postgres/` oder `deployments/keycloak/` nicht verändern, außer eine Änderung ist für
  den Betrieb des Go-Backends zwingend notwendig — die Änderung vorher
  erklären.
- Keine Testcontainers verwenden; Integrationstests müssen ohne Docker oder
  eine echte Datenbank laufen (siehe `tests/health_test.go`).
- Den Code für Go-Einsteiger verständlich halten — keine vorzeitige
  Abstraktion, keine zusätzlichen Abhängigkeiten außerhalb des festgelegten
  Stacks.
- `go`, `make` und `golangci-lint` waren auf der Maschine, auf der dieses
  Grundgerüst erstellt wurde, nicht installiert. Falls sie nicht verfügbar
  sind, dies ausdrücklich so benennen, statt anzunehmen, dass `go
  build`/`go test`/Lint erfolgreich durchgelaufen wären.
