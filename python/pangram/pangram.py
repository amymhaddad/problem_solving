# Initial Solution - Use a string
from string import ascii_lowercase as lowercase_letters


def is_pangram(sentence):
    letters = ""

    for char in sentence.lower():
        if char.isalpha() and char not in letters:
            letters += char
    return len(letters) == len(lowercase_letters)


# Approach 1
from string import ascii_lowercase as lowercase_letters


def is_pangram(sentence):
    total_letters = len(lowercase_letters) * [0]

    for char in sentence.lower():
        if char.isalpha():
            letter_index = ord(char) - ord("a")
            total_letters[letter_index] = 1

    return sum(total_letters) == len(lowercase_letters)


# Approach 2
from string import ascii_lowercase as lowercase_letters


def is_pangram(sentence):
    for letter in lowercase_letters:
        if letter not in sentence.lower():
            return False
    return True


# Approach 3 - Sol 1
from string import ascii_lowercase as lowercase_letters


def is_pangram(sentence):
    sentence_letters = set()

    for letter in sentence.lower():
        if letter.isalpha():
            sentence_letters.add(letter)
    return sentence_letters == set(lowercase_letters)


# Approach 3 - Sol 2
from string import ascii_lowercase as lowercase_letters


def is_pangram(sentence):
    total_letters = 0

    for i, char in enumerate(sentence):
        if char.isalpha():
            letter_index = ord(char.lower()) - ord("a")
            letter_position = 1 << letter_index
            total_letters = total_letters | letter_position

    return 2 ** len(lowercase_letters) - 1 == total_letters


# Approach 4 - Sol 1
import re


def is_pangram(sentence):
    return len(set(letter for letter in re.findall(r"[a-z]", sentence.lower()))) == 26


# Approach 4 - Sol 2
import re
from string import ascii_lowercase as lowercase_letters


def is_pangram(sentence):
    letters = set()
    for letter in re.findall(r"[a-z]", sentence.lower()):
        letters.add(letter)
    return len(letters) == len(lowercase_letters)
