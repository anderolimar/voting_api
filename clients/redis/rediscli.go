package rediscli

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func test() {
	// Conecta ao Redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Erro ao conectar ao Redis:", err)
		return
	}
	defer conn.Close()

	// Define uma chave e um valor
	key := "test_key"
	value := "test_value"

	// Armazena o valor na chave no Redis
	_, err = conn.Do("SET", key, value)
	if err != nil {
		fmt.Println("Erro ao definir a chave no Redis:", err)
		return
	}
	fmt.Printf("Chave '%s' definida com valor '%s'\n", key, value)

	// Recupera o valor da chave no Redis
	result, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("Erro ao obter a chave do Redis:", err)
		return
	}
	fmt.Printf("Chave '%s' tem valor: %s\n", key, result)

	// Inicia uma transação
	conn.Send("MULTI")
	conn.Send("INCR", "counter")
	conn.Send("INCR", "counter")
	conn.Send("GET", "counter")

	// Executa a transação
	res, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		fmt.Println("Erro ao executar a transação:", err)
		return
	}

	// Decodifica os resultados da transação
	var counter int
	_, err = redis.Scan(res, &counter)
	if err != nil {
		fmt.Println("Erro ao escanear o resultado da transação:", err)
		return
	}
	fmt.Printf("Valor do contador após a transação: %d\n", counter)
}
