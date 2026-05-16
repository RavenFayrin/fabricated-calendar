import unittest

from calendar_system import CalendarSystem
from calendar_classes import WeekDay, Month, Date

class TestCalendarSystem(unittest.TestCase):
    def setUp(self):
        self.calendar_system = CalendarSystem(
            "Test Calendar", 
            [
                WeekDay("WD1", 1),
                WeekDay("WD2", 2),
                WeekDay("WD3", 3),
                WeekDay("WD4", 4),
                WeekDay("WD5", 5),
                WeekDay("WD6", 6),
                WeekDay("WD7", 7)
                ],
            [
                Month("M1", 1, 31),
                Month("M2", 2, 28),
                Month("M3", 3, 31),
                Month("M4", 4, 30),
                Month("M5", 5, 31),
                Month("M6", 6, 30),
                Month("M7", 7, 31),
                Month("M8", 8, 31),
                Month("M9", 9, 30),
                Month("M10", 10, 31),
                Month("M11", 11, 30),
                Month("M12", 12, 31)
                ]
            )

    def test_get_num_days_in_year(self):
        num_days = self.calendar_system.get_num_days_in_year()
        self.assertEqual(num_days, 365)

    def test_get_num_weeks_in_year(self):
        num_weeks = self.calendar_system.get_num_weeks_in_year()
        self.assertEqual(num_weeks, 52)

    def test_date_to_absolute_day(self):
        date = Date(self.calendar_system, 2026, 1, 1)
        absolute_day = self.calendar_system.date_to_absolute_day(date)
        self.assertEqual(absolute_day, 2026 * 365 + 1)

    def test_absolute_day_to_date(self):
        absolute_day = 2026 * 365 + 1
        date = self.calendar_system.absolute_day_to_date(absolute_day)
        self.assertEqual(date, Date(self.calendar_system, 2026, 1, 1))

    def test_get_ab_days(self):
        early_date = Date(self.calendar_system, 2025, 6, 7)
        later_date = Date(self.calendar_system, 2026, 9, 16)
        days_between = self.calendar_system.get_ab_days(early_date, later_date)
        self.assertEqual(days_between, 466)

    def test_sort_dates(self):
        date_list = [
            Date(self.calendar_system, 2026, 12, 25),
            Date(self.calendar_system, 2025, 7, 4),
            Date(self.calendar_system, 2026, 1, 17),
            Date(self.calendar_system, 2026, 12, 24),
            Date(self.calendar_system, 1973, 11, 28)
        ]
        sorted_date_list = self.calendar_system.sort_dates(date_list)
        
        self.assertEqual(sorted_date_list, [
            Date(self.calendar_system, 1973, 11, 28),
            Date(self.calendar_system, 2025, 7, 4),
            Date(self.calendar_system, 2026, 1, 17),
            Date(self.calendar_system, 2026, 12, 24),
            Date(self.calendar_system, 2026, 12, 25)
            ]
        )

    def test_get_weekday_name(self):
        date = Date(self.calendar_system, 0, 1, 3)
        weekday_name = self.calendar_system.get_weekday_name(date)
        self.assertEqual(weekday_name, "WD3")

    def test_formatted_time_between_days(self):
        early_date = Date(self.calendar_system, 2025, 6, 7)
        later_date = Date(self.calendar_system, 2026, 9, 16)
        formatted_time = self.calendar_system.formatted_time_between_days(early_date, later_date)
        self.assertEqual(formatted_time, "1 year(s), 2 month(s), 39 day(s)")

if __name__ == '__main__':
    unittest.main()