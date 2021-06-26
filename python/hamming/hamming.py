def distance(strand_a, strand_b):
    return invalid_strand(strand_a, strand_b) or calculate_distance(strand_a, strand_b)


def invalid_strand(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("The strings need to be the same length.")


def calculate_distance(strand_a, strand_b):
    return sum(1 for s1, s2 in zip(strand_a, strand_b) if s1 != s2)
