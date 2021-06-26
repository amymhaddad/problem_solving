def largest(min_factor, max_factor):
    palindrome_products = get_palindrome_products(min_factor, max_factor)
    if min_factor == max_factor:
        return (None, [])
    largest_palindrome_product = max(palindrome_products)
    return (
        largest_palindrome_product,
        factors(min_factor, max_factor, largest_palindrome_product),
    )


def smallest(min_factor, max_factor):
    if min_factor > max_factor:
        raise ValueError("The min factor can't be larger than the max factor")

    palindrome_products = get_palindrome_products(min_factor, max_factor)
    if not palindrome_products:
        return (None, [])
    smallest_palindrome_product = min(palindrome_products)
    return (
        smallest_palindrome_product,
        factors(min_factor, max_factor, smallest_palindrome_product),
    )


def get_palindrome_products(min_factor, max_factor):
    start = min_factor
    all_multiples = set()

    while start <= max_factor:
        for num in range(start, max_factor + 1):
            multiple = start * num
            all_multiples.add(multiple)
        start += 1

    return [num for num in list(all_multiples) if is_palin(str(num))]


def factors(min_factor, max_factor, num):
    given = num
    total_factors = []
    for i in range(min_factor, max_factor + 1):
        for k in range(1, i + 1):
            if i * k == given:
                factor_pair = [i, k]
                total_factors.append(factor_pair)
    return total_factors


def is_palin(number_as_string):
    if len(number_as_string) <= 1:
        return True
    return number_as_string[0] == number_as_string[-1] and is_palin(
        number_as_string[1:-1]
    )
