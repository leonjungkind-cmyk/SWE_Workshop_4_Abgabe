# Technologiewahl im Go-Backend

## Übersicht

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

---

# 1. Gin

## Beschreibung

| Punkt | Erklärung |
|---|---|
| Aufgabe | Aufbau der REST-API |
| Einsatz | Definition von Endpunkten wie `GET`, `POST`, `PUT`, `DELETE` |
| Hilft bei | Routing, JSON-Antworten, Middleware, Request Binding, Pfadparametern |
| Vorteil | API-Struktur bleibt klar und gut lesbar |

## Vergleich mit Go-Alternativen

| Alternative | Kurzbeschreibung | Bewertung |
|---|---|---|
| `net/http` | Standardbibliothek von Go | Sehr stabil, aber mehr Handarbeit bei Routing und Middleware |
| Echo | Vollwertiges Go-Webframework | Gute Alternative, aber weniger verbreitet als Gin |
| Fiber | Express-ähnliches Framework auf Basis von `fasthttp` | Schnell, aber weniger nah am Go-Standard-HTTP-Modell |
| Chi | Sehr schlanker Router | Gut für minimalistische APIs, aber weniger Komfortfunktionen |
| Gin | REST-freundliches Go-Webframework | Gute Mischung aus Einfachheit, Funktionalität und Verbreitung |

## Nutzungs-/Beliebtheitsvergleich

| Paket | Imported by auf pkg.go.dev |
|---|---:|
| Gin | 183.243 |
| Echo | 31.336 |
| Fiber | 30.850 |
| Chi | 17.309 |

## Entscheidung

| Grund | Erklärung |
|---|---|
| Hohe Verbreitung | Gin wird deutlich häufiger importiert als Echo, Fiber und Chi |
| Verständlicher Aufbau | Routen und Handler sind einfach nachvollziehbar |
| Gute REST-Unterstützung | JSON, Binding und Middleware sind direkt nutzbar |
| Passend für das Projekt | Nicht zu minimalistisch, aber auch nicht unnötig schwergewichtig |

---

# 2. GORM

## Beschreibung

| Punkt | Erklärung |
|---|---|
| Aufgabe | Datenbankzugriff aus Go heraus |
| Einsatz | Datensätze speichern, lesen, aktualisieren und löschen |
| Arbeitsweise | Go-Structs werden auf Tabellen abgebildet |
| Vorteil | Weniger manuelles SQL für Standardoperationen |

## Vergleich mit Go-Alternativen

| Alternative | Kurzbeschreibung | Bewertung |
|---|---|---|
| `database/sql` | Standardbibliothek für SQL-Zugriff | Sehr kontrollierbar, aber viel manueller Code |
| sqlx | Erweiterung von `database/sql` | Mehr Komfort, aber weiterhin SQL-lastig |
| sqlc | Generiert typsicheren Go-Code aus SQL | Sehr sauber, aber zusätzlicher Generierungsprozess |
| Ent | Schema-basiertes ORM mit Codegenerierung | Mächtig, aber für ein überschaubares Projekt aufwendiger |
| Bun | SQL-first ORM/Query Builder | Gute Lösung, aber deutlich weniger verbreitet als GORM |
| GORM | Vollwertiges ORM | Gute Balance aus Produktivität, Lesbarkeit und Verbreitung |

## Nutzungs-/Beliebtheitsvergleich

| Paket | Imported by auf pkg.go.dev |
|---|---:|
| GORM | 86.926 |
| sqlx | 25.632 |
| Ent | 4.006 |
| Bun | 2.433 |

## Entscheidung

| Grund | Erklärung |
|---|---|
| Schnelle CRUD-Umsetzung | Standardoperationen lassen sich mit wenig Code umsetzen |
| Weniger Boilerplate | Weniger wiederholter SQL- und Mapping-Code |
| Gute Lesbarkeit | Datenbankzugriffe bleiben nah an den Go-Structs |
| Hohe Verbreitung | GORM ist im Vergleich zu anderen Go-Lösungen sehr stark genutzt |

---

# 3. GORM PostgreSQL Driver

## Beschreibung

| Punkt | Erklärung |
|---|---|
| Paket | `gorm.io/driver/postgres` |
| Aufgabe | Verbindung zwischen GORM und PostgreSQL |
| Einsatz | Öffnen der Datenbankverbindung über `gorm.Open(...)` |
| Vorteil | PostgreSQL kann direkt im GORM-Ansatz verwendet werden |

## Vergleich mit Go-Alternativen

| Alternative | Kurzbeschreibung | Bewertung |
|---|---|---|
| `lib/pq` | Klassischer PostgreSQL-Treiber | Stabil, aber eher klassischer direkter SQL-Zugriff |
| `pgx` | Moderner PostgreSQL-Treiber | Sehr leistungsfähig, aber ohne GORM-Abstraktion direkter zu verwenden |
| GORM PostgreSQL Driver | PostgreSQL-Dialector für GORM | Passt direkt zur Entscheidung für GORM |

## Nutzungs-/Beliebtheitswert

| Paket | Imported by auf pkg.go.dev |
|---|---:|
| GORM PostgreSQL Driver | 16.201 |

## Entscheidung

| Grund | Erklärung |
|---|---|
| Direkte GORM-Integration | Der Treiber passt direkt zum gewählten ORM |
| Passend zur Vorgabe | PostgreSQL ist vorgegeben, daher wird ein passender GORM-Treiber benötigt |
| Einheitlicher Zugriff | Kein gemischter Zugriffsstil zwischen ORM und direktem SQL nötig |

---

# 4. go-playground/validator/v10

## Beschreibung

| Punkt | Erklärung |
|---|---|
| Aufgabe | Prüfung eingehender Request-Daten |
| Einsatz | Validierung von DTOs und Request-Structs |
| Arbeitsweise | Regeln werden direkt als Struct-Tags definiert |
| Beispiel | `binding:"required,email"` oder `validate:"required"` |

## Vergleich mit Go-Alternativen

| Alternative | Kurzbeschreibung | Bewertung |
|---|---|---|
| Eigene Validierung | Manuelle `if`-Prüfungen im Handler oder Service | Einfach, aber schnell unübersichtlich |
| govalidator | Externe Validierungsbibliothek | Alternative, aber weniger passend zur Gin-Integration |
| ozzo-validation | Validierung über Code-Regeln | Flexibel, aber mehr Schreibaufwand |
| validator/v10 | Tag-basierte Struct-Validierung | Sehr passend für Gin und Request Binding |

## Nutzungs-/Beliebtheitswert

| Paket | Imported by auf pkg.go.dev |
|---|---:|
| go-playground/validator/v10 | 24.133 |

## Entscheidung

| Grund | Erklärung |
|---|---|
| Gute Gin-Integration | Passt direkt zum Binding-Ansatz von Gin |
| Übersichtliche Regeln | Validierungen stehen direkt am Request-Struct |
| Weniger Fehlerquellen | Keine verstreuten manuellen Prüfungen |
| Erweiterbar | Eigene Validierungsregeln sind möglich |

---

# 5. Go Modules

## Beschreibung

| Punkt | Erklärung |
|---|---|
| Aufgabe | Verwaltung externer Go-Abhängigkeiten |
| Dateien | `go.mod` und `go.sum` |
| Befehle | `go get`, `go mod tidy`, `go install` |
| Vorteil | Kein zusätzlicher Package Manager nötig |

## Vergleich mit Go-Alternativen

| Alternative | Kurzbeschreibung | Bewertung |
|---|---|---|
| GOPATH-Workflow | Alter Go-Ansatz vor Go Modules | Für moderne Projekte nicht mehr sinnvoll |
| Vendor-Ordner | Abhängigkeiten werden ins Projekt kopiert | Kann groß und unübersichtlich werden |
| Go Modules | Offizieller Dependency-Standard | Richtige Wahl für moderne Go-Projekte |

## Entscheidung

| Grund | Erklärung |
|---|---|
| Offizieller Standard | Direkt in Go integriert |
| Reproduzierbarkeit | Versionen und Prüfsummen werden festgehalten |
| Teamfähig | Alle Entwickler nutzen dieselben Abhängigkeiten |
| Einfacher Workflow | Wenige klare Befehle reichen aus |

---

# 6. gofmt

## Beschreibung

| Punkt | Erklärung |
|---|---|
| Aufgabe | Automatische Formatierung von Go-Code |
| Einsatz | Einheitliches Einrücken und Formatieren |
| Vorteil | Keine Diskussion über Code-Stil |
| Nutzung | Direkt über `gofmt` oder automatisch in der IDE |

## Vergleich mit Alternativen

| Alternative | Kurzbeschreibung | Bewertung |
|---|---|---|
| Manuelle Formatierung | Jeder Entwickler formatiert selbst | Uneinheitlich und fehleranfällig |
| Eigene Style-Regeln | Team definiert eigene Regeln | Unnötiger Aufwand im Go-Umfeld |
| gofmt | Standardformatter von Go | Einheitlich, etabliert und direkt verfügbar |

## Entscheidung

| Grund | Erklärung |
|---|---|
| Go-Standard | Gehört direkt zum Go-Ökosystem |
| Einheitlichkeit | Alle Dateien sehen gleich formatiert aus |
| Weniger Review-Aufwand | Code-Reviews können sich auf Logik statt Stil konzentrieren |

---

# 7. golangci-lint

## Beschreibung

| Punkt | Erklärung |
|---|---|
| Aufgabe | Statische Codeanalyse |
| Einsatz | Finden von potenziellen Fehlern und unsauberen Stellen |
| Vorteil | Viele Go-Linter werden gemeinsam ausgeführt |
| Typische Nutzung | Lokal und später auch in CI/CD möglich |

## Vergleich mit Go-Alternativen

| Alternative | Kurzbeschreibung | Bewertung |
|---|---|---|
| `go vet` | Offizielles Go-Prüfwerkzeug | Wichtig, aber weniger umfangreich |
| Einzelne Linter | Jeder Linter wird separat eingerichtet | Mehr Konfigurationsaufwand |
| golangci-lint | Linter-Runner für viele Go-Linter | Praktisch und professionell für Teamprojekte |

## Entscheidung

| Grund | Erklärung |
|---|---|
| Zentrale Konfiguration | Regeln können in einer Datei gesammelt werden |
| Viele Prüfungen | Mehr Abdeckung als nur `go vet` |
| CI/CD-geeignet | Gut für automatische Qualitätsprüfung |
| Teamqualität | Gleiche Regeln für alle Entwickler |

---

# 8. Docker Compose

## Beschreibung

| Punkt | Erklärung |
|---|---|
| Aufgabe | Start mehrerer lokaler Dienste |
| Einsatz | Einheitliches lokales Projektsetup |
| Datei | `docker-compose.yml` oder `compose.yml` |
| Vorteil | Gleiche Umgebung für alle Teammitglieder |

## Vergleich mit Alternativen

| Alternative | Kurzbeschreibung | Bewertung |
|---|---|---|
| Manuelle Installation | Dienste werden lokal direkt installiert | Fehleranfälliger und je Rechner unterschiedlich |
| Einzelne Docker-Befehle | Container werden einzeln gestartet | Unübersichtlich bei mehreren Diensten |
| Docker Compose | Dienste werden gemeinsam definiert und gestartet | Beste Wahl für ein lokales Team-Setup |

## Entscheidung

| Grund | Erklärung |
|---|---|
| Einheitliche Umgebung | Alle Teammitglieder nutzen dasselbe Setup |
| Schneller Start | Ein Befehl startet die benötigten Dienste |
| Gute Nachvollziehbarkeit | Konfiguration steht sichtbar in einer Compose-Datei |
| Weniger Setup-Probleme | Mehr Fokus auf Anwendungscode statt lokaler Installation |

---

# Zusammenfassung

| Technologie | Hauptgrund für die Auswahl |
|---|---|
| Gin | Beste Balance aus Einfachheit, REST-Funktionen und Verbreitung |
| GORM | Produktiver Datenbankzugriff mit wenig Boilerplate |
| GORM PostgreSQL Driver | Direkte Verbindung zwischen GORM und der vorgegebenen PostgreSQL-Datenbank |
| validator/v10 | Gute Gin-Integration und einfache Struct-Tag-Validierung |
| Go Modules | Offizieller Standard für Go-Abhängigkeiten |
| gofmt | Einheitliche Formatierung im gesamten Projekt |
| golangci-lint | Automatische Qualitätsprüfung für Go-Code |
| Docker Compose | Einheitliche lokale Umgebung für benötigte Dienste |

---

# Quellenstand

Die Nutzungswerte wurden über pkg.go.dev geprüft. Bei Go-Paketen wird dort nicht mit klassischen Downloadzahlen gearbeitet, sondern mit dem Wert **Imported by**. Dieser Wert zeigt, von wie vielen anderen Go-Modulen das jeweilige Paket importiert wird.

Verwendete Quellen:

- https://pkg.go.dev/github.com/gin-gonic/gin
- https://pkg.go.dev/github.com/labstack/echo/v4
- https://pkg.go.dev/github.com/gofiber/fiber/v2
- https://pkg.go.dev/github.com/go-chi/chi/v5
- https://pkg.go.dev/gorm.io/gorm
- https://pkg.go.dev/github.com/jmoiron/sqlx
- https://pkg.go.dev/entgo.io/ent
- https://pkg.go.dev/github.com/uptrace/bun
- https://pkg.go.dev/gorm.io/driver/postgres
- https://pkg.go.dev/github.com/go-playground/validator/v10
