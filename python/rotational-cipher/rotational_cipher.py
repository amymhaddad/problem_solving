
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
