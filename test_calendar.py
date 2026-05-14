import unittest

from calendar_classes import CalendarSystem, Date, WeekDay, Month

class TestCalendarSystem(unittest.TestCase):
    def setUp(self):
        self.calendar_system = CalendarSystem(
            "Test Calendar", 
            [
                WeekDay("Setacyde", 1),
                WeekDay("M’férst", 2),
                WeekDay("Teztor", 3),
                WeekDay("Wysellor", 4),
                WeekDay("TanéRhéim", 5),
                WeekDay("Fedrist", 6),
                WeekDay("Scynthay", 7)
                ],
            [
                Month("First", 1, 30),
                Month("Second", 2, 30),
                Month("Third", 3, 30),
                Month("Fourth", 4, 30),
                Month("Fifth", 5, 30),
                Month("Sixth", 6, 30),
                Month("Seventh", 7, 30),
                Month("Eigth", 8, 30),
                Month("Ninth", 9, 30),
                Month("Tenth", 10, 30),
                Month("Eleventh", 11, 30),
                Month("Twelfth", 12, 30)
                ]
            )

    def test_date_to_absolute_day(self):
        date = Date(self.calendar_system, 2506, 6, 7)
        absolute_day = self.calendar_system.date_to_absolute_day(date)
        self.assertEqual(absolute_day, 2506 * 360 + 150 + 7)

    def test_absolute_day_to_date(self):
        absolute_day = 2506 * 360 + 150 + 7
        date = self.calendar_system.absolute_day_to_date(absolute_day)
        self.assertEqual(date.year, 2506)
        self.assertEqual(date.month, 6)
        self.assertEqual(date.day, 7)

    def test_len_between_dates(self):
        early_date = Date(self.calendar_system, 2506, 6, 7)
        later_date = Date(self.calendar_system, 2507, 9, 16)
        time_between = self.calendar_system.len_between_dates(early_date, later_date)
        self.assertEqual(time_between.year, 1)
        self.assertEqual(time_between.month, 3)
        self.assertEqual(time_between.day, 9)

    def test_sort_dates(self):
        date_list = [
            Date(self.calendar_system, 2507, 8, 3),
            Date(self.calendar_system, 2506, 6, 7),
            Date(self.calendar_system, 2507, 9, 16),
            Date(self.calendar_system, 2498, 3, 25),
            Date(self.calendar_system, 2507, 9, 27)
        ]
        sorted_date_list = self.calendar_system.sort_dates(date_list)
        self.assertEqual(sorted_date_list[0].year, 2498)
        self.assertEqual(sorted_date_list[0].month, 3)
        self.assertEqual(sorted_date_list[0].day, 25)
        self.assertEqual(sorted_date_list[1].year, 2506)
        self.assertEqual(sorted_date_list[1].month, 6)
        self.assertEqual(sorted_date_list[1].day, 7)
        self.assertEqual(sorted_date_list[2].year, 2507)
        self.assertEqual(sorted_date_list[2].month, 8)
        self.assertEqual(sorted_date_list[2].day, 3)
        self.assertEqual(sorted_date_list[3].year, 2507)
        self.assertEqual(sorted_date_list[3].month, 9)
        self.assertEqual(sorted_date_list[3].day, 16)
        self.assertEqual(sorted_date_list[4].year, 2507)
        self.assertEqual(sorted_date_list[4].month, 9)
        self.assertEqual(sorted_date_list[4].day, 27)

if __name__ == '__main__':
    unittest.main()
