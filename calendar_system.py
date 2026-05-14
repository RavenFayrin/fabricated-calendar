from calendar_classes import Date


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
    
    def get_weekday_name(self, date):
        ab_date = self.date_to_absolute_day(date)
        remainder = (ab_date % len(self.weekdays)) - 1
        return self.weekdays[remainder].name

    def len_between_dates(self, early_date, later_date):
        early_ab = self.date_to_absolute_day(early_date)
        later_ab = self.date_to_absolute_day(later_date)
        len_between = later_ab - early_ab
        len_between = self.absolute_day_to_date(len_between)
        len_between.month -= 1
        return len_between

    def sort_dates(self, dates):
        ab_dates = []
        sorted_dates = []
        for date in dates:
            ab_dates.append(self.date_to_absolute_day(date))
        sorted_dates = sorted(ab_dates)
        for i in range(len(sorted_dates)):
            sorted_dates[i] = self.absolute_day_to_date(sorted_dates[i])
        return sorted_dates