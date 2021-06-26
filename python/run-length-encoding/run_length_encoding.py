import re


def decode(string):

    if _is_empty(string):
        return ""

    num = ""
    decode_string = ""
    for char in string:
        letter_or_space = re.search(r"[a-zA-Z ]+", char)
        single_character = letter_or_space and not num
        multiple_characters = letter_or_space and num

        if char.isdigit():
            num += char

        elif single_character:
            decode_string += char

        elif multiple_characters:
            decode_string += str(int(num) * char)
            num = ""

    return decode_string


def encode(string):
    if _is_empty(string):
        return ""

    total_encoded_chars = []
    letters = ""

    for i, char in enumerate(string):
        if letters == "" or char in letters:
            letters += char
        else:
            encoded_string = encode_letters(letters)
            total_encoded_chars.append(encoded_string)
            letters = char
    ec = encode_letters(letters)
    total_encoded_chars.append(ec)

    return "".join(total_encoded_chars)


def _is_empty(string):
    return len(string) == 0


def encode_letters(letters):

    compressed_chars = ""
    if len(letters) == 1:
        compressed_chars += letters[0]
    else:
        compressed_chars += str(len(letters)) + letters[0]
    return compressed_chars
