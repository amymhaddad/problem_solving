def factors(value):

    divisor = 2

    total_factors = []

    while value != 1:

        if value % divisor == 0:
            total_factors.append(divisor)
            result = value // divisor
            value = result
        else:

            divisor += 1

    return total_factors
