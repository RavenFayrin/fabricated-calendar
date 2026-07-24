# Fabricated Calendar

A Go library for creating and working with fully customizable calendar systems.

Fabricated Calendar is designed for worldbuilders, game developers, tabletop RPG creators, and anyone who needs a calendar that doesn't follow the rules of the Gregorian calendar.

Instead of assuming 12 months, 7-day weeks, or 365-day years, every part of the calendar can be customized to fit a fictional world.

> **Status:** 🚧 Early Development

---

## Features

Current functionality includes:

- Create custom calendar systems
- Define any number of months
- Define any number of weekdays
- Configure the number of days in each month
- Create dates within a calendar
- Convert dates to an absolute day number
- Convert an absolute day number back into a date
- Calculate weekday names
- Compare dates
- Sort collections of dates
- Calculate the number of days between two dates
- Format the time difference between two dates

---

## Example

```go
calendar := calendar.CalendarSystem{
    Name: "Example Calendar",

    Weekdays: []*calendar.WeekDay{
        {Name: "Sun", Order: 1},
        {Name: "Moon", Order: 2},
        {Name: "Star", Order: 3},
    },

    Months: []*calendar.Month{
        {Name: "Spring", Order: 1, NumDays: 30},
        {Name: "Summer", Order: 2, NumDays: 30},
        {Name: "Autumn", Order: 3, NumDays: 30},
    },
}

date, _ := calendar.NewDate(&calendar, 5, 2, 14)

fmt.Println(calendar.WeekdayName(date))
```

---

## Project Structure

```
fabricated-calendar/
│
├── calendar/
│   ├── calendar.go      # Core calendar calculations
│   ├── date.go          # Date type and helpers
│   ├── month.go         # Month definition
│   └── weekday.go       # Weekday definition
│
├── main.go              # Example entry point
└── go.mod
```

---

## Core Concepts

### CalendarSystem

A `CalendarSystem` defines the rules of a calendar.

It contains:

- calendar name
- list of weekdays
- list of months

Many calculations are performed directly by the calendar system, including:

- days in a year
- weeks in a year
- date conversion
- weekday lookup
- date sorting
- elapsed time calculations

---

### Date

A `Date` represents a single point in time within a specific calendar.

Each date stores:

- calendar reference
- year
- month
- day

Dates can be compared using:

- `Equals()`
- `Before()`

---

### Month

Months define:

- name
- order
- number of days

---

### WeekDay

Weekdays define:

- name
- order

---

## Installation

Clone the repository:

```bash
git clone https://github.com/RavenFayrin/fabricated-calendar.git
```

Change into the project directory:

```bash
cd fabricated-calendar
```

Run the application:

```bash
go run .
```

---

## Current API

### CalendarSystem

- `DaysInYear()`
- `WeeksInYear()`
- `DateToAbsoluteDay()`
- `AbsoluteDayToDate()`
- `WeekdayName()`
- `SortDates()`
- `DaysBetween()`
- `FormattedTimeBetween()`

### Date

- `AbsoluteDay()`
- `Equals()`
- `Before()`

### Constructors

- `NewDate()`

---

## Planned Features

The long-term vision for Fabricated Calendar includes:

- Multiple calendar systems
- Leap year support
- Custom eras
- Event tracking
- Recurring events
- Holiday support
- Calendar serialization
- Import/export
- Database persistence
- REST API
- Desktop GUI
- Web interface
- Calendar visualization
- Timeline generation
- Random calendar generation

---

## Goals

Fabricated Calendar aims to become a reusable engine for fictional timekeeping rather than a single calendar implementation.

The project is intended to support worlds with completely custom calendars while providing an easy-to-use API for applications, games, and worldbuilding tools.

---

## License

No license has currently been specified.

```
