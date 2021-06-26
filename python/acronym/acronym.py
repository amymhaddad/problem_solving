import re


def abbreviate(words):
    split_words = re.split("[\s_-]", words)

    return "".join([word[0].capitalize() for word in split_words if word])
