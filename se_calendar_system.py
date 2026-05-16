from calendar_system import CalendarSystem
from calendar_classes import WeekDay, Month, Date

se_calendar_system = CalendarSystem(
    "S.E. Calendar",
    [
        WeekDay("Setacyde", 1),
        WeekDay("M’férst", 2),
        WeekDay("Teztor", 3),
        WeekDay("Wysellor", 4),
        WeekDay("TanéRhéim", 5),
        WeekDay("Fedrist", 6),
        WeekDay("Scynthay", 7),
    ],
    [
        Month("First", 1, 30),
        Month("Second", 2, 30),
        Month("Third", 3, 30),
        Month("Fourth", 4, 30),
        Month("Fith", 5, 30),
        Month("Sixth", 6, 30),
        Month("Seventh", 7, 30),
        Month("Eighth", 8, 30),
        Month("Nineth", 9, 30),
        Month("Tenth", 10, 30),
        Month("Eleventh", 11, 30),
        Month("Twelfth", 12, 30),
    ])

def main():
    print("test, test, one, two, three")


if __name__ == "__main__":
    main()