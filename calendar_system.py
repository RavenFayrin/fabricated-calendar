from calendar_classes import Date


class CalendarSystem:
    def __init__(self, name, weekdays, months):
        self.name = name
        self.weekdays = weekdays
        self.months = months

    def get_num_days_in_year(self):
        days_in_year = 0
        for month in self.months:
            days_in_year += month.num_days
        return days_in_year
    
    def get_num_weeks_in_year(self):
        days_in_week = len(self.weekdays)
        return self.get_num_days_in_year() // days_in_week

    def date_to_absolute_day(self, date):
        days_in_year = self.get_num_days_in_year()
        absolute_day = date.year * days_in_year
        for month in self.months:
            if month.order < date.month.order:
                absolute_day += month.num_days
        absolute_day += date.day
        return absolute_day
    
    def absolute_day_to_date(self, absolute_day):
        days_in_year = self.get_num_days_in_year()
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

    def sort_dates(self, dates):
        ab_dates = []
        sorted_dates = []
        for date in dates:
            ab_dates.append(self.date_to_absolute_day(date))
        sorted_dates = sorted(ab_dates)
        for i in range(len(sorted_dates)):
            sorted_dates[i] = self.absolute_day_to_date(sorted_dates[i])
        return sorted_dates
    
    def get_ab_days(self, early_date, later_date):
        early_ab = self.date_to_absolute_day(early_date)
        later_ab = self.date_to_absolute_day(later_date)
        return later_ab - early_ab

    def formatted_time_between_days(self, early_date, later_date):
        years = 0
        months = 0 
        days = 0

        ab_days = self.get_ab_days(early_date, later_date)
        days_in_year = self.get_num_days_in_year()
        years += ab_days // days_in_year
        ab_days = ab_days % days_in_year

        days += early_date.month.num_days - early_date.day
        ab_days -= early_date.month.num_days - early_date.day
        days += later_date.day
        ab_days -= later_date.day

        starting_month = int(early_date.month.order)
        for month in self.months[starting_month:]:
            if ab_days >= month.num_days:
                ab_days -= month.num_days
                months += 1
            else:
                break
        return f"{years} year(s), {months} month(s), {days} day(s)"