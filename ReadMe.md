# Programmierworkshop am 19.6.2026

## Namen

Bastian Knebel  
Leon Jungkind

## Link zum Git-Repository

https://github.com/leonjungkind-cmyk/SWE_Workshop_4_Abgabe

## KI-Werkzeuge

Claude Sonnet 4.6
Claude integriert in VS code
ChatGPT-5.5

### Agenten

-

### Chat-URLs, z.B. https://chatgpt.com

-

## Frameworks und Bibliotheken

| Bereich | Entscheidung | Warum diese Wahl? |
|---|---|---|
| HTTP-Framework | Gin | Sehr verbreitet, übersichtliches Routing, gute REST-Unterstützung |
| Datenbankzugriff | GORM | Weniger Boilerplate, schnelle CRUD-Umsetzung, hohe Verbreitung |
| Datenbanktreiber | GORM PostgreSQL Driver | Passt direkt zu GORM und zur vorgegebenen PostgreSQL-Datenbank |
| Validierung | go-playground/validator/v10 | Gute Gin-Integration, Validierung direkt über Struct-Tags |
| Dependency Management | Go Modules | Offizieller Standard in Go, keine Zusatzinstallation nötig |
| Formatierung | gofmt | Einheitlicher Go-Code ohne eigene Style-Diskussionen |
| Linting | golangci-lint | Viele Go-Linter in einem Tool, gut für Teamarbeit und CI/CD |
| Lokales Setup | Docker Compose | Einheitliche lokale Umgebung für benötigte Dienste |

### REST-Schnittstelle (Lesen und Neuanlegen)

[github.com/gin-gonic/gin](https://github.com/gin-gonic/gin). Aktuell nur ein
Grundgerüst mit Routen-Gruppen `/api/public` (z.B. `GET /health`) und
`/api/secured` (geschützt durch OIDC-Middleware), siehe
[internal/server/server.go](internal/server/server.go). Es sind noch keine
fachlichen Endpunkte (Lesen/Neuanlegen) implementiert.

### Validierung (nur Neuanlegen)

[github.com/go-playground/validator/v10](https://github.com/go-playground/validator)
über Gins eingebautes Binding (`c.ShouldBindJSON` mit `validate:"..."`-Tags).
Platzhalter-Package: [internal/validation/](internal/validation/). Es gibt
noch keine fachlichen Validierungsregeln.

### OR-Mapping (für PostgreSQL)

[gorm.io/gorm](https://gorm.io/) mit `gorm.io/driver/postgres`. Verbindung
zur bestehenden PostgreSQL-Instanz aus
[deployments/postgres/compose.yml](deployments/postgres/compose.yml), siehe
[internal/database/database.go](internal/database/database.go). Es gibt noch
keine fachlichen Modelle, Repositories oder Migrationen.

### Optional: OIDC mit Keycloak

[github.com/coreos/go-oidc/v3](https://github.com/coreos/go-oidc) +
`golang.org/x/oauth2`. Middleware in
[internal/middleware/auth.go](internal/middleware/auth.go) prüft den
`Authorization: Bearer <token>`-Header gegen den konfigurierten Keycloak-Realm
und liefert `401` bei fehlendem/ungültigem Token. Nutzt die bestehende
Keycloak-Infrastruktur aus [deployments/keycloak/compose.yml](deployments/keycloak/compose.yml); ein
eigener Realm/Client für diese API muss in der Keycloak-Admin-Konsole noch
angelegt werden (siehe [deployments/keycloak/ReadMe.md](deployments/keycloak/ReadMe.md)).

### Einfacher Integrationstest

`net/http/httptest` + [github.com/stretchr/testify](https://github.com/stretchr/testify),
siehe [tests/health_test.go](tests/health_test.go). Ruft die
Public-Health-Route direkt über den Gin-Router auf, ohne echten HTTP-Server,
Docker oder Datenbank.

## Projekt-Setup (Go-Backend-Grundgerüst)

Dieser Abschnitt beschreibt das technische Setup des Go-Backends, das auf der
bestehenden PostgreSQL- und Keycloak-Infrastruktur aus früheren Abgaben
aufbaut (Ordner [deployments/postgres/](deployments/postgres/) und
[deployments/keycloak/](deployments/keycloak/), unverändert übernommen).

### Tech-Stack

| Bereich          | Bibliothek |
|-------------------|------------|
| Sprache           | Go 1.23+ |
| REST-Framework    | gin-gonic/gin |
| ORM               | gorm.io/gorm + gorm.io/driver/postgres |
| Validierung       | go-playground/validator/v10 |
| Security (OIDC)   | coreos/go-oidc/v3 + golang.org/x/oauth2 |
| Tests             | net/http/httptest + testify |
| Linting           | golangci-lint, gofmt, goimports |

Begründung der Auswahl: siehe [DECISIONS.md](DECISIONS.md).

### Bestehendes PostgreSQL-Setup

[deployments/postgres/compose.yml](deployments/postgres/compose.yml) startet PostgreSQL 18 mit TLS
(`ssl=on`) auf Host-Port `5432`. Das Superuser-Passwort steht in
`deployments/postgres/password.txt` (Docker Secret), unverändert aus der Vorabgabe.
`deployments/postgres/init/kunde/sql/` zeigt beispielhaft, wie eine eigene Datenbank,
ein Schema und ein DB-User angelegt werden (`create-db.sql`,
`create-schema.sql`); für diese API muss analog eine eigene DB angelegt
werden (siehe [deployments/postgres/ReadMe.md](deployments/postgres/ReadMe.md)).

### Bestehendes Keycloak-Setup

[deployments/keycloak/compose.yml](deployments/keycloak/compose.yml) startet Keycloak (inkl. der
PostgreSQL-Compose-Datei via `include:`) auf Host-Port `8880` (HTTP) und
`8843` (HTTPS). Der bisherige Realm `javascript` mit Client
`javascript-client` (siehe [deployments/keycloak/ReadMe.md](deployments/keycloak/ReadMe.md)) stammt
aus einer anderen Abgabe und wird nicht direkt wiederverwendet — für diese
API muss ein neuer Realm (z.B. `workshop`) und Client (z.B. `go-rest-api`)
nach demselben Schema angelegt werden. Laut [Aufgabe.md](Aufgabe.md) ist
Keycloak optional: ist der OIDC-Provider beim Start nicht erreichbar, startet
der Server trotzdem; `/api/secured` liefert dann für jede Anfrage `401`.

### Wie starten

```shell
cp .env.example .env
# .env mit echten lokalen Werten befüllen

go mod tidy
make run
# alternativ: go run ./cmd/api
```

PostgreSQL lokal starten:

```shell
cd deployments/postgres
docker compose up
```

Keycloak optional zusätzlich starten:

```shell
cd deployments/keycloak
docker compose up
```

### Wie testen

```shell
make test
# alternativ: go test ./... -v
```

Der enthaltene Test benötigt kein Docker, keine echte Datenbank und kein
Keycloak.

### Umgebungsvariablen

Siehe [.env.example](.env.example) für die vollständige, kommentierte Liste
(Server-, DB- und Keycloak-Variablen). Es sind keine Secrets im Repository
hinterlegt; `.env` ist über `.gitignore` ausgeschlossen.

## Prompts/Requests an KI-Agent/en

Beispiel Prompt:
dieser wurde nach einer langen Diskussion von Claude erstellt un anschließend in VScode mit Claude verwendet.

## Prompt 1 — Setup & Projektstruktur

Dieser Prompt kommt zuerst, noch bevor irgendeine Zeile Business-Logic geschrieben wird.

```
CONTEXT:
Go REST API project for a university workshop. Must be customer-deliverable quality.
No existing code yet — this is greenfield setup.

TECH STACK (fixed, do not deviate):
- Language:    Go 1.23+
- HTTP Router: github.com/gin-gonic/gin
- ORM:         gorm.io/gorm + gorm.io/driver/postgres
- Validation:  github.com/go-playground/validator/v10 (via Gin binding)
- Linter:      github.com/golangci/golangci-lint
- Formatter:   gofmt (built-in)

TASK:
Set up the complete project foundation. Create a TODO.md with every 
setup step, then execute them one by one and check each off.

Run all shell commands yourself:
- go mod init
- go get for every dependency above  
- go mod tidy
- Install the VS Code Go extension: code --install-extension golang.go
- Install golangci-lint via the official install script for Linux/macOS

Create:
- Project folder structure (follow standard Go project layout)
- Makefile with targets: build, lint, test, run
- .golangci.yml configured for: errcheck, gofmt, goimports, govet, staticcheck
- .env.example with all required environment variables
- DECISIONS.md explaining why each tool in the tech stack was chosen
- README.md skeleton with all required sections pre-filled where possible

DONE looks like:
- `make build` exits 0
- `make lint` exits 0 with zero warnings
- All files are committed with a meaningful initial commit message
- DECISIONS.md and README.md exist and are not empty

Do not write any business logic or handlers yet. Setup only.
```

---

## Prompt 2 — Datenbankanbindung

Kommt direkt nach Prompt 1, nachdem die Struktur steht.

```
CONTEXT:
Project setup is complete. Now connect to the existing PostgreSQL database.
The DB schema already exists — do not run migrations that create new tables.
Connection string comes from the DATABASE_URL environment variable.

TASK:
- Configure GORM to connect to PostgreSQL using DATABASE_URL from .env
- Implement connection pooling with sensible defaults
- Introspect the existing DB schema and create matching GORM model structs
- Add a startup health check that verifies the DB connection on launch
- If DATABASE_URL is missing or the connection fails, exit with a clear error message

DONE looks like:
- `make run` starts without errors when DATABASE_URL is set
- `make run` exits with a descriptive error when DATABASE_URL is missing
- All models reflect the actual existing DB tables
```

---

## Prompt 3 — REST Endpoints

```
CONTEXT:
DB connection and models are working.
Now implement the REST API layer.

TASK:
For each model from the existing DB, implement:
- GET  /[resource]      → list all records
- GET  /[resource]/:id  → get single record  
- POST /[resource]      → create new record with full input validation

All POST endpoints must validate input using Gin binding struct tags.
Return consistent JSON error responses for validation failures and not-found cases.

DONE looks like:
- `make run` serves all endpoints
- POST with missing required fields returns 400 with a readable error body
- GET /[resource]/999999 returns 404, not a 500
- `make lint` still exits 0
```

---

## Prompt 4 — Tests & Abgabe

```
CONTEXT:
REST API is working. Final step before submission.

TASK:
- Write one integration test per endpoint using net/http/httptest
- Tests must not require a real DB — use a mock or in-memory approach
- Run all tests and fix anything that fails
- Fill out README.md completely with: names, Git repo URL, 
  all AI tools used, all frameworks/libraries with versions,
  and the prompts that were used in this session
- Final check: make build, make lint, make test all exit 0

DONE looks like:
- `make test` is green
- README.md has no empty sections
- The repo is in a state that can be zipped and handed to a customer
```

---

Die Aufteilung hat einen klaren Grund: jeder Prompt hat ein abgeschlossenes, verifizierbares Ergebnis bevor der nächste startet. Dadurch bleibt der Context von Claude Code sauber und ihr merkt sofort wenn ein Schritt nicht funktioniert hat.

