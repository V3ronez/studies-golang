package enderecos

import "strings"

func TipoEndereco(texto string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	textoFinal := strings.ToLower(texto)
	primeiraPalavra := strings.Split(textoFinal, " ")[0]

	enderecoValido := false
	for _, palavra := range tiposValidos {
		if palavra == primeiraPalavra {
			enderecoValido = true
		}
	}
	if enderecoValido {
		return strings.Title(textoFinal)
	}
	return "tipo de endereco invalido"
}
