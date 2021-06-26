from collections import Counter
import re


def count_words(words):

    parsed_words = re.findall(r"[A-Za-z]+'[a-z]|[A-Za-z]+|[A-Za-z\d]", words)
    return Counter(word.lower() for word in parsed_words)
