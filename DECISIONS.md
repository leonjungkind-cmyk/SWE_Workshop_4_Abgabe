# Technologie-Entscheidungen

Dieses Dokument hält fest, warum jede Bibliothek im festgelegten Tech-Stack
gewählt wurde, für einen kleinen Go-REST-API-Prototyp im Rahmen eines
Universitäts-Workshops.

## Gin (`github.com/gin-gonic/gin`)

Minimales, schnelles, weit verbreitetes Go-HTTP/REST-Framework. Im Vergleich
zur reinen Standardbibliothek liefert es Routen-Gruppen, Middleware-Ketten
und JSON-Binding/Validierungs-Hooks direkt mit — genau das, was hier
benötigt wird, ohne ein großes Web-Framework mit mehr Konzepten einzuführen,
als ein Workshop braucht.

## GORM (`gorm.io/gorm` + `gorm.io/driver/postgres`)

Ausgereifter, gut dokumentierter Go-ORM mit erstklassiger
PostgreSQL-Unterstützung. Er lässt sich nahtlos an das bestehende
PostgreSQL-Setup anbinden (siehe `postgres/`) und hält den
Datenzugriffscode deklarativ und kurz, was für ein Projekt wichtig ist, das
für Go-Einsteiger verständlich bleiben soll.

## go-playground/validator/v10

Die de-facto Standard-Validierungsbibliothek für Go, und diejenige, mit der
Gins `ShouldBind*`-Methoden nativ über Struct-Tags zusammenarbeiten. Eine
andere Bibliothek zu verwenden würde bedeuten, Gins eingebaute
Binding-/Validierungs-Pipeline zu umgehen und mehr Code für das gleiche
Ergebnis zu schreiben.

## Keycloak mit OIDC

Aus einer früheren Abgabe wiederverwendet, statt selbst gebaut: es läuft
bereits (siehe `keycloak/`), und OIDC ist der Standard, sodass die
Go-Seite nur einen generischen OIDC-Client braucht, kein
Keycloak-spezifisches SDK.

## coreos/go-oidc/v3 (+ golang.org/x/oauth2)

`go-oidc` ist die Standard-Go-Bibliothek für OIDC-Discovery und
ID-Token-Verifikation (Issuer-/Audience-/Signatur-Prüfungen) und lässt sich
mit `golang.org/x/oauth2` für künftige Token-Exchange-Flows kombinieren.
Vermeidet selbst geschriebene JWT-/JWKS-Verifikation, bei der man leicht
subtile Fehler macht.

## net/http/httptest

Werkzeug aus der Standardbibliothek, um Gin-Routen zu testen, ohne einen
echten Socket zu binden oder einen echten Server zu starten — schnell,
ohne zusätzliche Abhängigkeit, und von der Aufgabenstellung ausdrücklich
verlangt, um Docker/Testcontainers in Tests zu vermeiden.

## testify

Fügt lesbare Assertions (`assert.Equal`, `assert.JSONEq`, ...) auf
`net/http/httptest` und Gos eingebautem `testing`-Paket hinzu, ohne den
Standard-Testworkflow zu ersetzen oder zu verschleiern — hält Tests auch
für Studierende zugänglich, die mit Gos Test-Idiomen noch nicht vertraut
sind.

## golangci-lint (+ gofmt, goimports)

Bündelt mehrere statische Go-Analysewerkzeuge (hier: `errcheck`, `govet`,
`staticcheck`, plus Formatierungsprüfungen) hinter einem Befehl und einer
Konfigurationsdatei (`.golangci.yml`), statt Studierende mehrere separate
Linter einzeln lernen und ausführen zu lassen.

## Lokale Entwicklung

Das ursprüngliche Image in `postgres/compose.yml`,
`dhi.io/postgres:18.3-debian13`, ist ein lizenziertes Docker Hardened
Image und konnte in dieser Umgebung nicht gepullt werden (`401
Unauthorized`, keine `dhi.io`-Zugangsdaten verfügbar). Für die lokale
Entwicklung wurde das Image auf das öffentliche `postgres:17-alpine`
umgestellt.

Das Standard-Datenverzeichnis dieses Ersatz-Images
(`/var/lib/postgresql/data`) entspricht nicht dem bestehenden
Volume-Mount (`pg_data:/var/lib/postgresql/18/data`, ausgelegt auf den
Debian-versionierten Pfad des ursprünglichen `dhi.io`-Images), sodass die
TLS-Zertifikate aus `postgres/init/tls` nie an dem Ort landen, an dem
PostgreSQL sie erwartet. Mit weiterhin gesetztem `command: ["-c",
"ssl=on"]` geriet der Container in eine Crash-Loop (`FATAL: could not load
server certificate file "server.crt"`). Die Lösung für die lokale
Entwicklung war, TLS zu deaktivieren: `ssl=on` in
`postgres/compose.yml` auskommentieren und `DB_SSLMODE` standardmäßig auf
`disable` setzen (sowohl in `.env.example` als auch in
`internal/config/config.go`).

Dies ist ein Workaround **nur für die lokale Entwicklung**. TLS wird nicht
als Konzept entfernt, sondern nur deaktiviert, weil das lokale
Ersatz-Image die Zertifikate, auf die das ursprüngliche Setup angewiesen
war, nicht bereitstellen kann. In Produktion (oder sobald
`dhi.io`-Registry-Zugang verfügbar ist) sollten das ursprüngliche
`dhi.io/postgres`-Image und sein `ssl=on`-Befehl wiederhergestellt und
`DB_SSLMODE` zurück auf `require` gesetzt werden — der Anwendungscode liest
`DB_SSLMODE` bereits aus der Konfiguration statt es hart zu kodieren, daher
ist für die Wiederaktivierung von TLS keine Code-Änderung nötig, nur
Konfiguration und das Postgres-Image/-Kommando in `postgres/compose.yml`.

## Projektstruktur (Go-Standard 2026)

`cmd/server/` wurde zu `cmd/api/` umbenannt, um dem inzwischen üblichen
Namensschema für das Hauptprogramm einer API zu folgen.

Der Ordner `migrations/` wurde angelegt (aktuell nur mit `.gitkeep`) als
zukünftiger Ablageort für GORM-Migrationen, sobald fachliche Modelle und
Schema-Änderungen hinzukommen.

Der Ordner `api/` wurde angelegt (aktuell nur mit `.gitkeep`) als
zukünftiger Ablageort für die OpenAPI-Spezifikation der REST-Schnittstelle.

## Projektstruktur

Die Ordner `postgres/` und `keycloak/` liegen bewusst im Projektstamm,
da sie aus vorherigen Abgaben übernommen wurden und unverändert
bleiben müssen. Nach Go-Best-Practices würden sie unter `deployments/`
liegen. Eine spätere Migration dorthin wäre der nächste saubere Schritt.
