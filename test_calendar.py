import unittest

from calendar_classes import CalendarSystem, Date, WeekDay, Month

class TestCalendarSystem(unittest.TestCase):
    def setUp(self):
        self.calendar_system = CalendarSystem(
            "Test Calendar", 
            [WeekDay("Setacyde", 1), WeekDay("M’férst", 2), WeekDay("Teztor", 3), WeekDay("Wysellor", 4), WeekDay("TanéRhéim", 5), WeekDay("Fedrist", 6), WeekDay("Scynthay", 7)],
            [Month("First", 1, 30), Month("Second", 2, 30), Month("Third", 3, 30), Month("Fourth", 4, 30), Month("Fifth", 5, 30), Month("Sixth", 6, 30),
             Month("Seventh", 7, 30), Month("Eigth", 8, 30), Month("Ninth", 9, 30), Month("Tenth", 10, 30), Month("Eleventh", 11, 30), Month("Twelfth", 12, 30)])

    def test_date_to_num(self):
        date = Date(self.calendar_system, 2506, 6, 7)
        num = self.calendar_system.date_to_num(date)
        self.assertEqual(num, 2506 * 360 + 150 + 7)

if __name__ == '__main__':
    unittest.main()