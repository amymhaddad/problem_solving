def convert(number):
    factors = get_factors(number)
    raindrop_sounds = get_raindrop_sounds(factors)

    return raindrop_sounds or str(number)


def get_factors(number):
    return [i for i in range(1, number + 1) if number % i == 0]


def get_raindrop_sounds(factors):
    sounds = {3: "Pling", 5: "Plang", 7: "Plong"}
    return "".join([sounds[factor] for factor in factors if factor in sounds])
