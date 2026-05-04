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
    def __init__(self, id, name, order, start_year, end_year):
        self.id = id
        self.name = name
        self.order = order
        self.start_year = start_year
        self.end_year = end_year
        self.length = end_year - start_year + 1