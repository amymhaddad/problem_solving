def slices(series, length):
    if length > len(series) or length <= 0:
        raise ValueError("Invalid length")

    substrings = []
    i = 0
    while length < len(series):
        substrings.append(series[i:length])
        i, length = i + 1, length + 1
    substrings.append(series[i:length])
    return substrings
