# TODO — Go-Backend-Grundgerüst

## Inspektion
- [x] Bestehenden `postgres/`-Ordner inspizieren (compose.yml, ReadMe.md, Init-Skripte, password.txt)
- [x] Bestehenden `keycloak/`-Ordner inspizieren (compose.yml, ReadMe.md, TLS-Dateien)
- [x] Ports, DB-Namen/-User, Keycloak-Ports/Realm-Hinweise für `.env.example` notieren
- [x] Verfügbarkeit von `go`, `make`, `golangci-lint` auf dieser Maschine prüfen

## Projektstruktur
- [x] `cmd/api/main.go` (umbenannt von `cmd/server/`, Go-Standard 2026)
- [x] `internal/config/`
- [x] `internal/database/`
- [x] `internal/server/`
- [x] `internal/middleware/`
- [x] `internal/model/`
- [x] `internal/repository/`
- [x] `internal/handler/`
- [x] `internal/validation/`
- [x] `tests/`
- [x] `docs/`
- [x] `migrations/` (Platzhalter, `.gitkeep` — künftige GORM-Migrationen)
- [x] `api/` (Platzhalter, `.gitkeep` — künftige OpenAPI-Spezifikation)

### Anpassung an Go-Standard 2026
- [x] `cmd/server/` zu `cmd/api/` umbenannt — `main.go` hat keine
  Selbstreferenz auf seinen eigenen Pfad, daher keine Code-/Import-Änderung
  nötig
- [x] `Makefile`: `build`- und `run`-Ziel von `./cmd/server` auf `./cmd/api`
  angepasst
- [x] `ReadMe.md`: Verweis auf `go run ./cmd/server` auf `./cmd/api`
  angepasst
- [x] `DECISIONS.md`: Abschnitt "Projektstruktur (Go-Standard 2026)" mit
  Begründung für `cmd/api/`, `migrations/` und `api/` ergänzt
- [x] `go build ./...` nach der Umstrukturierung erneut geprüft — siehe unten

## Go-Modul
- [x] `go.mod` (Modul `swe-workshop-api`, Go 1.26.4 — angehoben, nachdem Go 1.26.4 installiert wurde)
- [x] `go.sum` — von `go mod tidy` generiert, nachdem Go 1.26.4 installiert wurde

## Konfiguration & Infrastruktur-Anbindung
- [x] `internal/config`: Env-Vars für Server, Datenbank (Host/Port/Name/User/Passwort/SSLMode), Keycloak (Issuer-URL/Client-ID/erforderliche Rolle) laden
- [x] `internal/database`: GORM + `postgres`-Treiber, Platzhalter-Verbindungsaufbau, keine Migrationen
- [x] `internal/middleware`: OIDC/JWT-Platzhalter-Middleware (Bearer-Check, `go-oidc`-Verifier-Anbindung, 401 bei Fehler)
- [x] `internal/server`: Gin-Engine-Setup mit Routen-Gruppen `/api/public` und `/api/secured`, Health-Endpoint
- [x] `internal/model`, `internal/repository`, `internal/handler`, `internal/validation`: nur Platzhalter-Packages, keine Fachlogik

## Tests
- [x] `tests/`: httptest + testify Integrationstest gegen die Health-Route (kein DB, kein Docker)

## Tooling & Dokumentation
- [x] `Makefile` (build, run, test, lint, fmt, tidy)
- [x] `.golangci.yml` (errcheck, gofmt, goimports, govet, staticcheck)
- [x] `.env.example` (Server-/DB-/Keycloak-Variablen, abgeleitet aus bestehenden Compose-Dateien, keine Secrets)
- [x] `.gitignore`-Ergänzungen für Go-Build-Artefakte
- [x] `README.md` (Beschreibung, Tech-Stack, bestehendes Postgres-/Keycloak-Setup, REST-/Validierungs-/ORM-/OIDC-/Test-Abschnitte, Run-/Test-Anleitung, Env-Vars)
- [x] `DECISIONS.md` (Begründung je Technologie-Entscheidung)
- [x] `CLAUDE.md` (Leitfaden für zukünftige KI-unterstützte Arbeit in diesem Repo)

## Befehle (vom Assistenten ausgeführt)
- [x] Go 1.26.4 installiert via `winget install GoLang.Go --version 1.26.4`
- [x] `golangci-lint` v1.64.8 installiert via `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
- [x] `go mod tidy` — hat `go.sum` erzeugt, alle indirekten Abhängigkeiten aufgelöst
- [x] `go build ./...` — Exit 0
- [x] `go vet ./...` — Exit 0
- [x] `golangci-lint run ./...` — Exit 0, keine Warnungen
- [x] `go test ./...` — alle Tests erfolgreich
- [x] `GET /api/public/health` liefert `200 {"status":"ok"}` gegen eine **echte** Postgres-Verbindung verifiziert (kein Fallback nötig) — siehe unten
- [ ] Installation der VS-Code-Go-Extension — bewusst übersprungen, Projektvorgaben verbieten das automatische Installieren von Extensions

### Gelöst: `dhi.io`-Registry blockiert, auf lokales Ersatz-Image umgestellt
`dhi.io/postgres:18.3-debian13` konnte nicht gepullt werden (`401
Unauthorized`, lizenzierte Registry, keine Zugangsdaten verfügbar). Die
`image:`-Zeile in `postgres/compose.yml` wurde für die lokale Entwicklung
auf `postgres:17-alpine` geändert. Alles andere in `postgres/compose.yml`
(TLS-Dateien, Port, Volumes, `password.txt`) blieb wie vorgegeben
unverändert, **außer** der Zeile `command: ["-c", "ssl=on"]`, die
auskommentiert werden musste: Das Standard-Datenverzeichnis von
`postgres:17-alpine` (`/var/lib/postgresql/data`) entspricht nicht dem
bestehenden Volume-Mount (`pg_data:/var/lib/postgresql/18/data`, ausgelegt
auf den Debian-versionierten Pfad des ursprünglichen `dhi.io`-Images),
sodass die TLS-Zertifikate aus `postgres/init/tls` nie an dem Pfad landen,
den PostgreSQL erwartet — mit weiterhin gesetztem `ssl=on` geriet der
Container in eine Crash-Loop (`FATAL: could not load server certificate
file "server.crt"`). Dies wurde durch Ausführen und Auswerten von `docker
logs postgres` bestätigt, bevor über die Lösung entschieden wurde.

Nettoeffekt: Das lokale Ersatz-Setup läuft **ohne TLS** und **ohne
Volume-Persistenz** für `pg_data` (Daten liegen in der beschreibbaren
Container-Schicht, nicht im benannten Volume, wegen des oben beschriebenen
Pfad-Mismatches — akzeptabel für die lokale Entwicklung, nicht für Daten,
die eine Container-Neuerstellung überstehen sollen). `DB_SSLMODE` in
`.env.example` wurde passend von `require` auf `disable` geändert.
`keycloak/` wurde nicht berührt.

Ende-zu-Ende verifiziert:
1. `docker compose -f postgres/compose.yml up -d` → `docker ps` zeigt
   `postgres:17-alpine ... Up (healthy)`
2. Eine eigene DB/User angelegt (nach dem Muster von
   `postgres/init/kunde/sql`) via:
   `docker exec -e PGPASSWORD=p postgres psql --dbname=postgres --username=postgres -c "CREATE USER app PASSWORD 'app';" -c "CREATE DATABASE app;" -c "GRANT ALL ON DATABASE app TO app;"`
3. `go run ./cmd/server` mit `DB_SSLMODE=disable` (sowie `DB_NAME`/
   `DB_USER`/`DB_PASSWORD=app`) ausgeführt — das strikte `log.Fatalf` bei
   DB-Fehlern in `main.go` wurde **nicht** ausgelöst, d. h. dies war eine
   echte Verbindung, nicht der frühere Fallback.
4. `curl http://localhost:8080/api/public/health` → `200 {"status":"ok"}`

Um zum ursprünglichen `dhi.io`-Image + TLS zurückzukehren, sobald
lizenzierter Registry-Zugang verfügbar ist: die `image:`-Zeile und die
auskommentierte `command:`-Zeile in `postgres/compose.yml`
zurücksetzen, dann das TLS-Bootstrapping in `postgres/ReadMe.md` befolgen
(der Volume-Pfad muss dabei an das tatsächlich verwendete Image angepasst
werden).

### TLS für lokale Entwicklung deaktiviert — Config-Default aktualisiert, erneut verifiziert
- [x] `.env.example`: `DB_SSLMODE=disable` (bereits im vorherigen Schritt gesetzt)
- [x] `internal/config/config.go`: Fallback-Default von `DB_SSLMODE` von `require` auf `disable` geändert, mit Kommentar-Verweis auf DECISIONS.md und Hinweis, dass dies zurück auf `require` muss, sobald das ursprüngliche TLS-fähige `dhi.io/postgres`-Image wiederhergestellt ist
- [x] `docker compose -f postgres/compose.yml up -d` → `docker ps` bestätigt `postgres:17-alpine ... Up (healthy)`
- [x] `go build ./...` — Exit 0 nach der Config-Änderung
- [x] `go run ./cmd/server` gestartet **ohne** `DB_SSLMODE` explizit zu setzen (verlässt sich auf den neuen Code-Default) → echte DB-Verbindung, kein Fallback
- [x] `curl http://localhost:8080/api/public/health` → `200 {"status":"ok"}`
- [x] `DECISIONS.md`: Abschnitt "Lokale Entwicklung" ergänzt, der erklärt, warum TLS lokal deaktiviert ist und wie es in Produktion wieder aktiviert wird

### Verschiebung nach deployments/ (Go-Standard-Layout)
- [x] `postgres/` → `deployments/postgres/` verschoben (`git mv`, Inhalt unverändert)
- [x] `keycloak/` → `deployments/keycloak/` verschoben (`git mv`, Inhalt unverändert)
- [x] `deployments/keycloak/compose.yml`s `include: ../postgres/compose.yml` geprüft — Pfad bleibt gültig, da beide Ordner gemeinsam als Geschwister unter `deployments/` verschoben wurden; keine Änderung nötig
- [x] `Makefile` geprüft — enthält keine `docker compose -f`-Befehle, daher keine Anpassung nötig
- [x] `ReadMe.md`: alle Pfadverweise und `cd postgres`/`cd keycloak`-Anleitungen auf `deployments/postgres/` bzw. `deployments/keycloak/` aktualisiert
- [x] `.env.example`: alle Kommentar-Pfadverweise auf `deployments/postgres/` bzw. `deployments/keycloak/` aktualisiert
- [x] `CLAUDE.md`: Pfadverweise aktualisiert
- [x] `internal/config/config.go`, `internal/database/database.go`, `internal/middleware/auth.go`: Kommentar-Pfadverweise aktualisiert (keine funktionale Änderung)
- [x] `DECISIONS.md`: Abschnitt "Projektstruktur" überarbeitet — erklärt jetzt, warum `deployments/postgres/` und `deployments/keycloak/` dem Go-Standard-Layout entsprechen
- [x] `docker compose -f deployments/postgres/compose.yml config` — valide
- [x] `docker compose -f deployments/keycloak/compose.yml config` — valide, `include`-Pfad löst korrekt auf
- [x] `go build ./...` — Exit 0 nach der Verschiebung
- [x] `go test ./...` — Exit 0 nach der Verschiebung

Bewusst nicht geändert: die historischen Protokoll-Einträge weiter oben in
dieser Datei (z.B. zum `dhi.io`-Registry-Problem), da sie den damals
tatsächlich gültigen Pfad `postgres/compose.yml` korrekt wiedergeben.

## Abschlussprüfung
- [x] Ergebnis gegen die Qualitätskriterien prüfen, bevor an den Nutzer zurückgemeldet wird
