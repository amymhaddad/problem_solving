import re


def is_valid(isbn):
    if not valid_numbers(isbn):
        return False

    numbers = valid_numbers(isbn)
    isbn_length = len(numbers)

    total = 0
    for i, num in enumerate(numbers, 1):
        if num == "X":
            num = 10
        total += int(num) * isbn_length
        isbn_length -= 1

    return total % 11 == 0


def valid_numbers(isbn):
    numbers = [num for num in isbn[:-1] if num.isdigit()]

    valid_check_char = re.search(r"[X\d]$", isbn)
    if valid_check_char:
        numbers.append(isbn[-1])

    if not valid_check_char or len(numbers) != 10:
        return False
    return numbers
