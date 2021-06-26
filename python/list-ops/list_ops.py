import functools


def append(list1, list2):
    return list1 + list2


def concat(lists):
    total = []
    for list in lists:
        total += list
    return total


def filter(function, list):
    return [num for num in list if function(num)]


def length(list):
    return len(list)


def map(function, list):
    return [function(num) for num in list]


def foldl(function, list, initial):
    if not list:
        return initial
    return functools.reduce(function, list, initial)


def foldr(function, list, initial):
    if not list:
        return initial
    acc = initial
    for num in reversed(list):
        acc = function(num, acc)
    return acc


def reverse(list):
    if not list:
        return []

    ptr1 = 0
    ptr2 = len(list) - 1

    while ptr2 - ptr1 >= 1:
        list[ptr2], list[ptr1] = list[ptr1], list[ptr2]
        ptr1 += 1
        ptr2 -= 1
    return list
