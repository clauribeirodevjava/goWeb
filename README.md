# goWeb

# Criar um container mysql para poder usar um exemplo de banco de dados:

# Criar e Iniciar um Container MySQL

docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 mysql:latest
# Acessar o Container MySQL
docker exec -it mysql-container mysql -uroot -p

# Digitar a senha quando solicitado
# Iniciar a sessão MySQL e criar um banco de dados e uma tabela

CREATE DATABASE go_course;

CREATE USER 'root'@'%' IDENTIFIED BY 'root';
GRANT ALL PRIVILEGES ON go_course.* TO 'root'@'%';
FLUSH PRIVILEGES;

USE go_course;

CREATE TABLE exemplo (
  id INT NOT NULL AUTO_INCREMENT,
  nome VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

# ##################################################Criar um programa Go para conectar e testar o container mysql###################################

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Configurar a string de conexão
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_course?charset=utf8")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer db.Close()

	// Preparar a consulta SELECT
	query := "SELECT id, nome FROM exemplo"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Erro ao preparar a consulta SELECT:", err)
		return
	}
	defer rows.Close()

	// Iterar sobre os resultados
	for rows.Next() {
		var id int
		var nome string
		if err := rows.Scan(&id, &nome); err != nil {
			fmt.Println("Erro ao ler os resultados:", err)
			return
		}
		fmt.Printf("ID: %d, Nome: %s\n", id, nome)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Erro nos resultados:", err)
	}
}


