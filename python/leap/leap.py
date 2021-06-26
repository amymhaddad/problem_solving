def leap_year(year):
    if year % 4 != 0:
        return False

    if year % 100 and year % 400 or year % 400 == 0:
        return True
    else:
        return False
    return True
