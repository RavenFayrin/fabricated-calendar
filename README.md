# Fabricated Calendar

A desktop application for creating and managing **fully customizable fantasy calendars**.

Fabricated Calendar is designed for worldbuilders, tabletop RPG creators, game developers, writers, and anyone who needs a calendar that doesn't follow the rules of the Gregorian calendar.

Instead of assuming 12 months, 7 weekdays, or 365 days per year, Fabricated Calendar lets users define the structure of their own fictional calendar, including the number and names of weekdays, the number and length of months, and how those months are ordered.

> **Status:** 🚧 Early Development

<img width="1119" height="696" alt="image" src="https://github.com/user-attachments/assets/8238bb7e-3a33-436f-89e1-309f664091ad" />

---

## ✨ Features

### 👤 User Management

Fabricated Calendar includes basic account management:

* Create a user account
* Log in with a username and password
* Log out of the current account
* Delete a user account
* Store passwords using Argon2id hashing
* Validate email addresses during account creation

User and calendar data are persisted in PostgreSQL.

---

### 📅 Calendar Management

Users can create and manage multiple custom calendars.

* Create calendars
* View available calendars
* Edit calendar names and descriptions
* Delete calendars
* Select a calendar to work with

Each calendar belongs to a user, allowing multiple independent calendars to be created and managed within an account.

---

### 🗓️ Custom Weekdays

Weekdays are completely configurable for each calendar.

* Create weekdays
* Edit weekdays
* Delete weekdays
* Set weekday order
* Use any number of weekdays
* Give weekdays custom names

A calendar does not have to use seven weekdays. A fictional calendar could have five, eight, ten, or any other number of weekdays.

---

### 🌙 Custom Months

Months are also completely configurable.

* Create months
* Edit months
* Delete months
* Set month order
* Set the number of days in each month
* Give months custom names
* Use any number of months

Month lengths are independent of the Gregorian calendar.

---

### 🗓️ Custom Calendar Display

The application includes a custom calendar renderer rather than relying on the Gregorian-based calendar.

The calendar display supports:

* Custom month names
* Custom weekday names
* Custom month lengths
* Custom weekday counts
* Year display
* Month/year selection
* Previous/next month pagination
* Day number placement based on the calendar's own rules

---

## 🚀 Getting Started

### Prerequisites

You will need:

* Go 1.26.3 or compatible
* PostgreSQL
* Goose
* A system capable of running Fyne applications

Fyne may require additional system libraries depending on the operating system.

---

## 🗄️ Database Setup

Fabricated Calendar uses PostgreSQL for persistent storage.

The database currently consists of four primary tables:

```text
users
  │
  └── calendar
       ├── weekday
       └── month
```

### Users

Stores account information including:

* Username
* Hashed password
* Email address
* Creation timestamp
* Update timestamp

### Calendar

Stores the user's calendars and their descriptions.

### Weekday

Stores the custom weekdays belonging to a calendar.

Each weekday has:

* Name
* Display order
* Calendar ID
* User ID

### Month

Stores the custom months belonging to a calendar.

Each month has:

* Name
* Month order
* Days in month
* Calendar ID
* User ID

Foreign keys use `ON DELETE CASCADE`, so deleting a user removes their calendars and deleting a calendar removes its associated months and weekdays.

---

## ⚙️ Configuration

The application reads its database connection string from the `DATABASE_URL` environment variable.

Create a `.env` file in the project root:

```env
DATABASE_URL=postgres://username:password@localhost:5432/fabricated_calendar?sslmode=disable
```

Do not commit your `.env` file or database credentials to source control.

---

## 🐘 Database Migrations

Database migrations are located in:

```text
sql/schema/
```

They are applied in dependency order:

```text
001_users.sql
002_calendar.sql
003_weekday.sql
004_months.sql
```

With Goose installed, migrations can be applied using:

```bash
goose -dir sql/schema postgres "$DATABASE_URL" up
```

---

## ▶️ Running the Application

After PostgreSQL is running and `DATABASE_URL` has been configured:

```bash
go run .
```

The project also includes:

```bash
./run.sh
```

The development script sets the locale to `C.utf8` before starting the application. This helps avoid locale parsing issues with Fyne in environments such as WSL2.

---

## 🏗️ Architecture

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

### GUI

The `gui/` package contains the Fyne desktop interface.

It is responsible for:

* Login and account screens
* Calendar selection
* Calendar management
* Month management
* Weekday management
* Calendar rendering
* Month/year navigation
* Refreshing the interface after database changes
* Custom application theming

### Authentication

The `internal/auth/` package contains user-related business logic.

It handles:

* Password hashing
* Password verification
* Email validation
* User creation
* User deletion
* Login

### Calendar Logic

The `internal/calendar/` package contains calendar-specific application logic.

It handles:

* Calendar CRUD operations
* Month CRUD operations
* Weekday CRUD operations
* Loading complete calendar data
* Calendar calculations
* Month starting weekday calculations

### Database

PostgreSQL provides persistent storage.

SQL queries are written manually under `sql/queries/` and generated into type-safe Go code using SQLC.

Database schema changes are managed using Goose migrations under `sql/schema/`.

### 🧪 Testing

Run the project's Go tests with:

```bash
go test ./...
```

For verbose output:

```bash
go test -v ./...
```

Current tests include authentication validation and calendar calculation tests.

Calendar calculation tests cover:

* Days per year
* Multiple month lengths
* Different weekday counts
* Month starting weekdays
* Month-to-month weekday continuation
* Year-to-year weekday continuation

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

## 📋 Current Release

The current release provides the core functionality needed to create and use a custom fantasy calendar.

## 🔮 Future Development

Fabricated Calendar is still under active development. Planned improvements include:

* Event creation and display
* Recurring events
* Holiday support
* Expanded calendar rules
* Additional customization options

---

## 🎯 Project Goals

Fabricated Calendar is intended to make fictional timekeeping easy to create and manage.

The goal is not to reproduce the Gregorian calendar, but to provide a foundation where the **rules of time itself can be customized**.

A finished calendar might have:

* Five weekdays
* Thirteen months
* Months with completely different lengths
* A year that doesn't contain 365 days
* Custom month and weekday names
* A unique fictional world's own concept of time

The application provides the tools needed to define those rules and visualize the resulting calendar.

---

## 📜 License

Fabricated Calendar is licensed under the **Apache License 2.0**.

See [`LICENSE`](LICENSE) for the complete license text.
