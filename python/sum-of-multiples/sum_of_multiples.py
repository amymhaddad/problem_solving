def sum_of_multiples(limit, multiples):

    zero_multiples = len(multiples) == 1 and multiples[0] == 0 or multiples == []

    if zero_multiples:
        return 0

    each_multiple = 0
    unique_multiples = set()
    while each_multiple < len(multiples):

        if multiples[each_multiple] == 0:
            each_multiple += 1
            continue

        for i in range(multiples[each_multiple], limit, multiples[each_multiple]):
            unique_multiples.add(i)
        each_multiple += 1

    return sum(unique_multiples)
