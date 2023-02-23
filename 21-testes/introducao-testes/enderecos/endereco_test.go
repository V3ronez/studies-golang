package enderecos

//TESTE DE UNIDADE
import (
	"testing"
)

func TestTipoEndereco(t *testing.T) {
	enderecoEnviado := "Rua dos Bobos"
	respostaEsperada := "Rua dos Bobos"
	tipoDeEndereco := TipoEndereco(enderecoEnviado)

	if respostaEsperada != tipoDeEndereco {
		t.Error("o tipo recebido é invalido")
	}
	return 
}
