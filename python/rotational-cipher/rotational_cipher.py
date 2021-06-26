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
