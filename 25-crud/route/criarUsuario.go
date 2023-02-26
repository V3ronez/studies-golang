package route

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lib/pq"
)

type usuarioType struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("falha ao receber o json"))
	}
	var usuario usuarioType
	if err = json.Unmarshal(body, &usuario); err != nil {
		w.Write([]byte("Erro ao converter json em struct"))
		return
	}

	db, err := banco.Connectar()
	if err != nil {
		w.Write([]byte("Erro ao connectar."))
		return
	}
	defer db.Close()
	beginTrans, err := db.Begin()
	if err != nil {
		w.Write([]byte("Erro ao iniciar transicao"))
		// "Erro ao criar o statement."
		return
	}

	statement, err := beginTrans.Prepare(pq.CopyIn("usuarios", "id", "nome", "email"))
	if err != nil {
		w.Write([]byte(err.Error()))
		// "Erro ao criar o statement."
		return
	}
	beginTrans.Commit()
	defer statement.Close()

	// inserir, err := statement.Exec(usuario.Nome, usuario.Email)
	// if err != nil {
	// 	w.Write([]byte("Erro ao inserir o usuario."))
	// 	return
	// }

	w.WriteHeader(http.StatusCreated) //201
	w.Write([]byte(fmt.Sprintln("Usuario cadastrado")))
}
