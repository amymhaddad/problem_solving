import itertools as it


def cipher_text(plain_text):
    letters = [char.lower() for char in plain_text if char.isalnum()]

    cols, rows = _create_rectangle(letters)

    spaces_needed = cols * rows - len(letters)
    if spaces_needed:
        letters.append(" ")

    rectangle = []
    start = 0
    stop = cols
    for i in range(rows):
        each_row = letters[start:stop]
        rectangle.append(each_row)
        start = stop
        stop += cols

    cols = [list(col) for col in it.zip_longest(*rectangle, fillvalue=" ")]
    return " ".join(["".join(col) for col in cols])


def _create_rectangle(chars):
    col = 1

    while True:
        row = col - 1
        for num in range(2):
            if (row * col >= len(chars)) and col >= row and col - row <= 1:
                return col, row
            else:
                row += 1
        col += 1
    return col, row
