# Welcome to Chit

> Pronounced "Cheat".

A Go scaffolding tool and CRUD code generator for [chi](https://github.com/go-chi/chi) + [GORM](https://gorm.io) API projects.

## Features

- **Scaffold** a complete chi API project with dependency injection, router, and DB connection
- **Interactive DB prompt** — choose PostgreSQL, MySQL, or SQLite during scaffold
- **Dependency bundling** — auto-installs chi, GORM, Cobra, CORS, JWT, UUID, bcrypt
- **CRUD generation** — parses your model structs via Go AST and generates repository + controller files
- **FIFO tracking** — skips already-generated models so you can incrementally add models

## Prerequisites

- Go 1.26+
- git

## Installation

### From source

```bash
git clone https://github.com/Mizael-go/chit
cd chit
go build
```

This produces a `chit` binary in the current directory.

### Via `go install`

```bash
go install github.com/Mizael-go/chit@latest
```

## Usage

| Command              | Description                                  |
| -------------------- | -------------------------------------------- |
| `chit create <name>` | Scaffold a new project (prompts for DB type) |
| `chit generate`      | Generate CRUD from structs in `model/*.go`   |
| `chit help`          | Show available commands                      |

---

### `chit create <name>`

Creates a full project skeleton under `./<name>/`.

```bash
chit create myapp
```

During creation you choose a database driver:

- **PostgreSQL** — reads `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_PORT` from `.env`
- **MySQL** — same env vars as PostgreSQL
- **SQLite** — uses `database.db` local file, no `.env` needed

What gets created:

```
myapp/
├── main.go              # entrypoint, calls cmd.Execute()
├── .env                 # DB credentials (Postgres/MySQL only)
├── chit.md              # placeholder (unused)
├── go.mod
├── cmd/
│   └── root.go          # Cobra root command, inits DB/repo/controller/router, starts :4400
├── app/
│   └── db.go            # GORM connection setup
├── controller/
│   └── controller.go    # base Controller struct with *gorm.DB + *repository.Repository
├── repository/
│   └── repository.go    # base Repository struct with *gorm.DB
├── router/
│   └── router.go        # chi router with CORS (localhost:5173), stub routes
└── model/               # place your Go struct files here
```

The scaffold also runs `go mod init`, `git init`, and installs all dependencies.

---

### `chit generate`

Generates CRUD repository and controller files for every struct found in `model/*.go`.

```bash
cd myapp
# Create your model structs in model/ first
chit generate
```

**How it works:**

1. Scans all `.go` files in `model/`
2. Parses them with `go/parser` + `go/ast` to extract top-level struct names
3. Skips files already tracked in `generateModels.json` (FIFO-based dedup)
4. For each new struct generates:
   - `repository/<struct>Repository.go` — `GetAll<Model>`, `Get<Model>ById`, `Create<Model>`, `Update<Model>`, `Delete<Model>ById`
   - `controller/<struct>Controller.go` — corresponding HTTP handlers

**Example:**

Given `model/user.go` with `type User struct`, running `chit generate` creates `repository/userRepository.go` and `controller/userController.go`.

To add CRUD for another model later, just add the struct file to `model/` and run `chit generate` again.

---

## Dependencies (scaffolded project)

These are installed into the generated project (not into chit itself):

| Package                      | Purpose                     |
| ---------------------------- | --------------------------- |
| `gorm.io/gorm`               | ORM                         |
| `github.com/go-chi/chi/v5`   | HTTP router                 |
| `github.com/joho/godotenv`   | Load `.env`                 |
| `gorm.io/driver/<db>`        | DB driver (per your choice) |
| `github.com/spf13/cobra`     | CLI framework               |
| `github.com/go-chi/cors`     | CORS middleware             |
| `github.com/golang-jwt/jwt`  | JWT auth                    |
| `github.com/google/uuid`     | UUID generation             |
| `golang.org/x/crypto/bcrypt` | Password hashing            |

## Known issues / Caveats

- `--help` is listed in `IsAllowed` but not handled in `cmd/app.go` — use `help` or `-h` instead
- `chit.md` placeholder file is written to scaffolded projects but has no purpose
- The `generate` command requires an existing `model/` directory with valid Go struct files

## Development

```bash
git clone https://github.com/Mizael-go/chit
cd chit
go build
./chit help
```
