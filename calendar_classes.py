class WeekDay:
    def __init__(self, name, order):
        self.name = name
        self.order = order

class Month:
    def __init__(self, name, order, num_days):
        self.name = name
        self.order = order
        self.num_days = num_days

class Era:
    def __init__(self, name, order, start_year, end_year):
        self.name = name
        self.order = order
        self.start_year = start_year
        self.end_year = end_year