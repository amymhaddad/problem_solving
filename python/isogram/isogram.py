import re


def is_isogram(string):
    letters = re.sub(r"[- ]+", "", string.lower())
    return len(letters) == len(set(letters))
