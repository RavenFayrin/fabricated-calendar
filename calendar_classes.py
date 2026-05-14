class WeekDay:
    def __init__(self, name, order):
        self.name = name
        self.order = order

class Month:
    def __init__(self, name, order, num_days):
        self.name = name
        self.order = order
        self.num_days = num_days

class Date:
    def __init__(self, calendar_system, year, month, day):
        self.calendar_system = calendar_system
        self.year = year
        self.month = month
        self.day = day

    def __eq__(self, other):
        if isinstance(other, Date):
            return self.calendar_system == other.calendar_system \
                and self.day == other.day \
                and self.month == other.month \
                and self.year == other.year 
        return False
