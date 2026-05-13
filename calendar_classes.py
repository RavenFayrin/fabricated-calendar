class WeekDay:
    def __init__(self, name, order):
        self.name = name
        self.order = order

class Month:
    def __init__(self, name, order, num_days):
        self.name = name
        self.order = order
        self.num_days = num_days

class CalendarSystem:
    def __init__(self, name, weekdays, months):
        self.name = name
        self.weekdays = weekdays
        self.months = months

    def date_to_num(self, date):
        days_in_year = 0
        for month in self.months:
            days_in_year += month.num_days
        num = date.year * days_in_year
        for month in self.months:
            if month.order < date.month:
                num += month.num_days
        num += date.day
        return num
    
    #def num_to_date(self):

class Date:
    def __init__(self, calendar_system, year, month, day):
        self.calendar_system = calendar_system
        self.year = year
        self.month = month
        self.day = day
