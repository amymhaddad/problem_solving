from string import ascii_uppercase as letters

letters_to_numbers = {letter: i for i, letter in enumerate(letters)}


def diamond_indexes(given):
    length = letters_to_numbers[given] + letters_to_numbers[given] + 1
    each_row = [0] * length

    letter_index_decreasing = letters_to_numbers[given]
    letter_index_increasing = 1

    for i in range(len(each_row)):
        if letter_index_decreasing >= 0:
            each_row[i] = letter_index_decreasing
            letter_index_decreasing -= 1
        else:
            each_row[i] = letter_index_increasing
            letter_index_increasing += 1
    return each_row


def top_half_diamond(each_row, given):
    letter_index_top = 0
    result = []

    while letter_index_top <= letters_to_numbers[given]:
        row_total = []
        for index in each_row:
            if index == letter_index_top:
                row_total.append(letters[letter_index_top])
            else:
                row_total.append(" ")
        result.append(row_total)
        letter_index_top += 1
    return result


def bottom_half_diamond(results_top_half, given, each_row):

    result = results_top_half
    letter_index_bottom = letters_to_numbers[given] - 1

    while letter_index_bottom >= 0:
        row_total = []
        for index in each_row:
            if index == letter_index_bottom:
                row_total.append(letters[letter_index_bottom])
            else:
                row_total.append(" ")

        result.append(row_total)
        letter_index_bottom -= 1
    return result


def generate_diamond(diamond):
    return ["".join(rows) for rows in diamond]


def rows(given):
    if letters[0] == given:
        return [letters[0]]

    each_row = diamond_indexes(given)
    results_top_half = top_half_diamond(each_row, given)
    diamond = bottom_half_diamond(results_top_half, given, each_row)
    return generate_diamond(diamond)
