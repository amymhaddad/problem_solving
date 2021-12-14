
#Sol 1 
from string import ascii_lowercase as lower, ascii_uppercase as upper

ALPHA_LENGTH = 26


def rotate(text, key):

    new_string = ""
    for char in text:
        if not char.isalpha():
            new_string += char
            continue

        if char.isupper():
            upper_to_index = {letter: i for i, letter in enumerate(upper)}
            curr_index = upper_to_index[char]
            letters = upper

        else:
            lower_to_index = {letter: i for i, letter in enumerate(lower)}
            curr_index = lower_to_index[char]
            letters = lower

        rotated_index = (curr_index + key) % ALPHA_LENGTH
        new_string += list(letters)[rotated_index]

    return new_string

#Sol 2
ALPHA_LENGTH = 26
START = 97

def rotate(text, key):
    rotated_chars = []

    is_upper = False
    for ch in text:
        if not ch.isalpha():
            rotated_chars.append(ch)
            continue
        if ch.isupper():
            is_upper = True
            ch = ch.lower()
        shift = ord(ch) - START
        shift_val = (shift + key) % ALPHA_LENGTH
        new_val = shift_val + START

        if is_upper:
            new_val = chr(new_val).upper()
            rotated_chars.append(new_val)
            is_upper = False
        else:
            rotated_chars.append(chr(new_val))
    return "".join(rotated_chars)
