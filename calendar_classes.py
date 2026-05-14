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

    def date_to_absolute_day(self, date):
        days_in_year = 0
        for month in self.months:
            days_in_year += month.num_days
        absolute_day = date.year * days_in_year
        for month in self.months:
            if month.order < date.month:
                absolute_day += month.num_days
        absolute_day += date.day
        return absolute_day
    
    def absolute_day_to_date(self, absolute_day):
        days_in_year = 0
        for month in self.months:
            days_in_year += month.num_days
        year = absolute_day // days_in_year
        absolute_day = absolute_day % days_in_year
        month = 1
        for m in self.months:
            if absolute_day >= m.num_days:
                absolute_day -= m.num_days
                month += 1
            else:
                break
        day = absolute_day
        return Date(self, year, month, day)

    def len_between_dates(self, early_date, later_date):
        early_ab = date_to_absolute_day(early_date)
        later_ab = date_to_absolute_day(later_date)
        len_between = later_ab - early_ab
        return absolute_day_to_date(len_between)

class Date:
    def __init__(self, calendar_system, year, month, day):
        self.calendar_system = calendar_system
        self.year = year
        self.month = month
        self.day = day
