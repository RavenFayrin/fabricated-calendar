 class WeekDay:
    def __init__(self, id, name, order):
        self.id = id
        self.name = name
        self.order = order

class Month:
    def __init__(self, id, name, order, num_days):
        self.id = id
        self.name = name
        self.order = order
        self.num_days = num_days

class Era:
    def __init__(self, id, name, order, start_year):
        self.id = id
        self.name = name
        self.order = order
        self.start_year = start_year

class CalendarSystem:
    def __init__(self, name, weekdays, months, eras):
        self.name = name
        self.weekdays = weekdays
        self.months = months
        self.eras = eras

    #def date_to_num(self):

    #def num_to_date(self):

class Date:
    def __init__(self, calendar_system, era, year, month, day):
        self.calendar_system = calendar_system
        self.era = era
        self.year = year
        self.month = month
        self.day = day
