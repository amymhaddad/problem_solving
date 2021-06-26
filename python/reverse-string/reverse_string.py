def reverse(text):

    if not text:
        return ""

    text_as_list = [letter for letter in text]
    start, stop = 0, len(text_as_list) - 1

    while stop >= start:
        text_as_list[start], text_as_list[stop] = (
            text_as_list[stop],
            text_as_list[start],
        )
        stop, start = stop - 1, start + 1

    return "".join(text_as_list)
