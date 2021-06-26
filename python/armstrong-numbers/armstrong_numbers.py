def is_armstrong_number(number):

    total, given_number = 0, number

    while given_number > 0:
        total += (given_number % 10) ** len(str(number))
        given_number //= 10

    return number == total
