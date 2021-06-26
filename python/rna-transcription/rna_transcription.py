dna_rna = {"G": "C", "C": "G", "T": "A", "A": "U"}


def to_rna(dna_strand):

    return "".join([dna_rna[ele] for ele in dna_strand])
