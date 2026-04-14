date = "04.13.2026"

def parse_date(date):
    if "/" in date:
        parsed = date.split("/")
    elif "." in date:
        parsed = date.split(".")
    elif "-" in date:
        parsed = date.split("-")
    else:
        raise Exception(f"ERROR: {date} is an invailed format")
    return parsed
    

def date_to_num(date):
    parsed = parse_date(date)
    year, month, day = parsed[2], parsed[0], parsed[1]
    return int(year + month + day)

print(date_to_num(date))

#def num_to_date(date):


#def time_between_dates(date1, date2):
