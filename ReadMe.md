# Programmierworkshop am 19.6.2026

## Namen

Bastian Knebel  
Leon Jungkind

## Link zum Git-Repository

## KI-Werkzeuge

### Agenten

### Chat-URLs, z.B. https://chatgpt.com

## Frameworks und Bibliotheken

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
[postgres/compose.yml](postgres/compose.yml), siehe
[internal/database/database.go](internal/database/database.go). Es gibt noch
keine fachlichen Modelle, Repositories oder Migrationen.

### Optional: OIDC mit Keycloak

[github.com/coreos/go-oidc/v3](https://github.com/coreos/go-oidc) +
`golang.org/x/oauth2`. Middleware in
[internal/middleware/auth.go](internal/middleware/auth.go) prüft den
`Authorization: Bearer <token>`-Header gegen den konfigurierten Keycloak-Realm
und liefert `401` bei fehlendem/ungültigem Token. Nutzt die bestehende
Keycloak-Infrastruktur aus [keycloak/compose.yml](keycloak/compose.yml); ein
eigener Realm/Client für diese API muss in der Keycloak-Admin-Konsole noch
angelegt werden (siehe [keycloak/ReadMe.md](keycloak/ReadMe.md)).

### Einfacher Integrationstest

`net/http/httptest` + [github.com/stretchr/testify](https://github.com/stretchr/testify),
siehe [tests/health_test.go](tests/health_test.go). Ruft die
Public-Health-Route direkt über den Gin-Router auf, ohne echten HTTP-Server,
Docker oder Datenbank.

## Projekt-Setup (Go-Backend-Grundgerüst)

Dieser Abschnitt beschreibt das technische Setup des Go-Backends, das auf der
bestehenden PostgreSQL- und Keycloak-Infrastruktur aus früheren Abgaben
aufbaut (Ordner [postgres/](postgres/) und [keycloak/](keycloak/), unverändert
übernommen).

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

[postgres/compose.yml](postgres/compose.yml) startet PostgreSQL 18 mit TLS
(`ssl=on`) auf Host-Port `5432`. Das Superuser-Passwort steht in
`postgres/password.txt` (Docker Secret), unverändert aus der Vorabgabe.
`postgres/init/kunde/sql/` zeigt beispielhaft, wie eine eigene Datenbank,
ein Schema und ein DB-User angelegt werden (`create-db.sql`,
`create-schema.sql`); für diese API muss analog eine eigene DB angelegt
werden (siehe [postgres/ReadMe.md](postgres/ReadMe.md)).

### Bestehendes Keycloak-Setup

[keycloak/compose.yml](keycloak/compose.yml) startet Keycloak (inkl. der
PostgreSQL-Compose-Datei via `include:`) auf Host-Port `8880` (HTTP) und
`8843` (HTTPS). Der bisherige Realm `javascript` mit Client
`javascript-client` (siehe [keycloak/ReadMe.md](keycloak/ReadMe.md)) stammt
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
cd postgres
docker compose up
```

Keycloak optional zusätzlich starten:

```shell
cd keycloak
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
