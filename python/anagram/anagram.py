from collections import Counter


def find_anagrams(word, candidates):
    sorted_given_word = "".join([letter for letter in sorted(word.lower())])
    anagrams = {}
    for candidate in candidates:
        if candidate.lower() == word.lower():
            continue
        anagram = is_anagram(candidate, sorted_given_word)
        if anagram and candidate.lower() not in anagrams:
            anagrams[candidate.lower()] = candidate
    return [word for word in anagrams.values()]


def is_anagram(candidate, sorted_given_word):

    sorted_candidate = "".join([letter for letter in sorted(candidate.lower())])

    return Counter(sorted_given_word) == Counter(sorted_candidate)
