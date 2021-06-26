from datetime import datetime, timedelta

GIGASECOND = 10 ** 9


def add(moment):
    return moment + timedelta(seconds=GIGASECOND)
