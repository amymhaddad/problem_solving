letters_value = {
    "A,E,I,O,U,L,N,R,S,T": 1,
    "D,G": 2,
    "B,C,M,P": 3,
    "F,H,V,W,Y": 4,
    "K": 5,
    "J,X": 8,
    "Q,Z": 10,
}


def score(word):
    total = 0
    for letter in word:
        for letters, value in letters_value.items():
            if letter.upper() in letters:
                total += value
    return total
