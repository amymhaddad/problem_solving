class Matrix:
    def __init__(self, matrix_string):
        self.matrix = [num for num in matrix_string.split("\n")]

    def row(self, index):
        return [int(num) for num in self.matrix[index - 1].split() if num.isdigit()]

    def column(self, index):
        total = [numbers.split(" ") for numbers in self.matrix]
        return [int(number[index - 1]) for number in total]
