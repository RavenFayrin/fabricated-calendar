# Fabricated Calendar

A desktop application for creating and managing fully customizable fantasy calendars.

Fabricated Calendar is designed for worldbuilders, tabletop RPG creators, game developers, writers, and anyone who needs a calendar that doesn't follow the rules of the Gregorian calendar.

Instead of assuming 12 months, 7 weekdays, or 365 days per year, Fabricated Calendar lets users define the structure of their own fictional calendar, including the number and names of weekdays, the number and length of months, and how those months are ordered.

> **Status:** 🚧 Early Development

<img width="1119" height="696" alt="Fabricated Calendar screenshot" src="https://github.com/user-attachments/assets/8238bb7e-3a33-436f-89e1-309f664091ad" />

---

## Motivation

I enjoy worldbuilding and wanted a simple, flexible way to keep track of dates, recurring events, and everything in between. When I looked for an existing solution, I found only one that seemed to fit my needs—but it was so complicated that it made my head spin.

So, I decided to build my own from scratch.

Fabricated Calendar is intended to make fictional timekeeping easy to create and manage without forcing every world to follow the Gregorian calendar.

---

## Quick Start

### Prerequisites

You will need:

* Go 1.26.3 or compatible
* PostgreSQL
* Goose

### 1. Configure the Database

Create a PostgreSQL database and set the `DATABASE_URL` environment variable.

For local development, create a `.env` file in the project root:

```env
DATABASE_URL=postgres://username:password@localhost:5432/fabricated_calendar?sslmode=disable
```

### 2. Run Migrations

With Goose installed:

```bash
goose -dir sql/schema postgres "$DATABASE_URL" up
```

### 3. Launch the Application

```bash
go run .
```

Alternatively:

```bash
./run.sh
```

The included script also handles the locale configuration needed by Fyne in environments such as WSL2.

---

## Usage

### 👤 User Management

* Create a user account
* Log in with a username and password
* Log out
* Delete a user account
* Securely store passwords using Argon2id
* Validate email addresses during account creation

User and calendar data are persisted in PostgreSQL.

### 📅 Calendar Management

Users can create and manage multiple custom calendars.

* Create calendars
* Select between calendars
* Edit calendar names and descriptions
* Delete calendars

Each calendar belongs to a user, allowing multiple independent calendars to be managed within the same account.

### 🗓️ Custom Weekdays

Weekdays are completely configurable for each calendar.

* Create, edit, and delete weekdays
* Set weekday order
* Use any number of weekdays
* Give weekdays custom names

A calendar can have five, eight, ten, or any other number of weekdays.

### 🌙 Custom Months

Months are also completely configurable.

* Create, edit, and delete months
* Set month order
* Set the number of days in each month
* Give months custom names
* Use any number of months

Month lengths are independent of the Gregorian calendar.

### 🗓️ Calendar Display

Fabricated Calendar uses its own calendar renderer rather than relying on a Gregorian-based calendar.

The current display supports:

* Custom month and weekday names
* Custom month lengths and weekday counts
* Year display
* Month/year selection
* Previous/next month navigation
* Day placement based on the calendar's own rules

---

## Contributing

Contributions, ideas, bug reports, and feedback are welcome!

If you'd like to contribute:

1. Fork the repository.
2. Create a branch for your changes.
3. Make your changes and add tests where appropriate.
4. Run the test suite:

```bash
go test ./...
```

5. Open a pull request with a description of your changes.

For larger changes, opening an issue first is recommended so the proposed direction can be discussed before significant work is done.

---

### Architecture

The project is separated into several major components.

```text
┌───────────────────────┐
│       Fyne GUI        │
│         gui/          │
└───────────┬───────────┘
            │
            ▼
┌───────────────────────┐
│   Application Logic   │
│   internal/auth/      │
│   internal/calendar/  │
└───────────┬───────────┘
            │
            ▼
┌───────────────────────┐
│       SQLC            │
│  internal/database/   │
└───────────┬───────────┘
            │
            ▼
┌───────────────────────┐
│     PostgreSQL        │
└───────────────────────┘
```

#### GUI

The `gui/` package contains the Fyne desktop interface and is responsible for:

* Login and account screens
* Calendar selection and management
* Month and weekday management
* Calendar rendering and navigation
* Refreshing the interface after database changes
* Custom application theming

#### Authentication

The `internal/auth/` package contains user-related business logic, including:

* Password hashing and verification
* Email validation
* User creation and deletion
* Login

#### Calendar Logic

The `internal/calendar/` package contains calendar-specific application logic, including:

* Calendar, month, and weekday CRUD operations
* Loading complete calendar data
* Calendar calculations
* Month starting weekday calculations

#### Database

PostgreSQL provides persistent storage.

SQL queries are written under `sql/queries/` and generated into type-safe Go code using SQLC. Database schema changes are managed through Goose migrations under `sql/schema/`.

---

## 🛠️ Tech Stack

| Technology     | Purpose                            |
| -------------- | ---------------------------------- |
| **Go**         | Application language               |
| **Fyne**       | Desktop GUI framework              |
| **PostgreSQL** | Persistent database                |
| **SQLC**       | Type-safe database code generation |
| **Goose**      | Database schema migrations         |
| **Argon2id**   | Password hashing                   |
| **UUID**       | Database identifiers               |
| **godotenv**   | Local environment configuration    |

---

## 🔮 Future Development

Fabricated Calendar is still under active development. Planned improvements include:

* Event creation and display
* Recurring events
* Holiday support
* Expanded calendar rules
* Additional customization options

---

## 📜 License

Fabricated Calendar is licensed under the **Apache License 2.0**.

See [`LICENSE`](LICENSE) for the complete license text.
