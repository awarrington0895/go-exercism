package strand

func complement(dnaNucleotide rune) rune {
	switch dnaNucleotide {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		return 'A'
	}
}

func ToRNA(dna string) string {
	var rna []rune

	for _, nucleotide := range dna {
		rna = append(rna, complement(nucleotide))
	}

	return string(rna)
}
