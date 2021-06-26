import re


def response(hey_bob):

    letters = "".join(re.findall(r"[a-zA-Z]|[a-zA-Z\?]", hey_bob))
    question = "".join(letters).endswith("?")

    if letters.isupper() and question:
        return "Calm down, I know what I'm doing!"

    if letters and question or question:
        return "Sure."

    if hey_bob.isspace() or not hey_bob:
        return "Fine. Be that way!"

    if letters.isupper():
        return "Whoa, chill out!"

    return "Whatever."
