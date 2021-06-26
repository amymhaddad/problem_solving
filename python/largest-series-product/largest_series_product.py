from functools import reduce


def largest_product(series, size):
    return invalid_input(series, size) or get_product(series, size)


def invalid_input(series, size):
    if size > len(series):
        raise ValueError("The size exceeds the series length")

    if series.isalpha():
        raise ValueError("The series can only contain digits.")

    if size < 0:
        raise ValueError("The size must be a positive number.")


def get_product(series, size):
    if not series or not size:
        return 1

    nums_as_list = [int(num) for num in series]

    pt1, pt2 = 0, size

    values = []
    while pt2 <= len(nums_as_list):
        values.append(nums_as_list[pt1:pt2])
        pt1 += 1
        pt2 += 1

    return max([reduce(lambda x, y: x * y, number) for number in values])
