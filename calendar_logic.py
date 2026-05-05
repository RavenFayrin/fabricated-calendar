def parse_date(date_str):
    if "/" in date_str:
        parsed = date_str.split("/")
    elif "." in date_str:
        parsed = date_str.split(".")
    elif "-" in date_str:
        parsed = date_str.split("-")
    else:
        raise Exception(f"ERROR: {date_str} is an invailed format")
    return parsed
    

def date_to_num(date_str):
    parsed = parse_date(date_str)
    year, month, day = parsed[2], parsed[0], parsed[1]
    return int(year + month + day)

def num_to_date(date_int, delimiter="/"):
    date_str = str(date_int)
    year, month, day = date_str[0:4], date_str[4:6], date_str[6:8]
    return f"{month}{delimiter}{day}{delimiter}{year}"


#def time_between_dates(older_date, newer_date):
