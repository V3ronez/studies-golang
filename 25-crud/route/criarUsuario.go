package route

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	statement, err := db.Prepare("INSERT INTO usuarios(nome, email) VALUES ($1, $2) RETURNING id")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement."))
		return
	}
	defer statement.Close()

	inserir, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao inserir o usuario."))
		return
	}

	w.WriteHeader(http.StatusCreated) //201
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso: %d", inserir)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	var query string = "SELECT * FROM usuarios"

	db, err := banco.Connectar()

	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}
	defer db.Close()

	linhas, err := db.Query(query)
	if err != nil {
		w.Write([]byte("Erro ao buscar ao banco"))
		return
	}
	defer linhas.Close()

	var usuarios []usuarioType
	for linhas.Next() {
		var usuario usuarioType
		if err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao buscar ao banco"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao fazer encode do dados do banco"))
		return
	}
}

func BuscarUsuarioID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	stringQuery := "SELECT * FROM usuarios where id = $1"
	ID, err := strconv.ParseUint(parametros["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao parse parametros"))
		return
	}
	db, err := banco.Connectar()

	if err != nil {
		w.Write([]byte("Erro ao connectar ao banco"))
		return
	}

	linha, err := db.Query(stringQuery, ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar a linha"))
		return
	}

	var usuario usuarioType
	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao scannear linha"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Erro ao encode json usuario"))
		return
	}
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Erro ao pegar o parametro"))
		return
	}
	corpoRequisicao, err := io.ReadAll(r.Body)
	fmt.Appendln(corpoRequisicao)
	if err != nil {
		w.Write([]byte("Erro ao pegar o corpo da requisicao"))
		return
	}
	var usuario usuarioType
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		w.Write([]byte("Erro ao fazer o parse do json"))
		return
	}
	db, err := banco.Connectar()

	if err != nil {
		w.Write([]byte("Erro ao se connectar ao banco"))
		return
	}
	defer db.Close()

	var query string = "UPDATE usuarios SET nome = $1, email = $2 WHERE id = $3"
	stmt, err := db.Prepare(query)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(&usuario.Nome, &usuario.Email, ID); err != nil {
		w.Write([]byte("Erro ao atualizar usuario"))
	}
	w.WriteHeader(http.StatusNoContent)

}
