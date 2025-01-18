package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Parsing dos argumentos da linha de comando
	url := flag.String("url", "", "URL do serviço a ser testado.")
	totalRequests := flag.Int("requests", 0, "Número total de requests.")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas.")
	flag.Parse()

	// Validação dos parâmetros
	if *url == "" {
		fmt.Println("Erro: --url é obrigatório.")
		flag.Usage()
		os.Exit(1)
	}
	if *totalRequests <= 0 {
		fmt.Println("Erro: --requests deve ser maior que 0.")
		flag.Usage()
		os.Exit(1)
	}
	if *concurrency <= 0 {
		fmt.Println("Erro: --concurrency deve ser maior que 0.")
		flag.Usage()
		os.Exit(1)
	}
	if *concurrency > *totalRequests {
		*concurrency = *totalRequests
	}

	// Iniciando o temporizador
	startTime := time.Now()

	// Canais para gerenciar as requisições
	requests := make(chan int, *totalRequests)

	// Preenchendo o canal de requisições
	for i := 0; i < *totalRequests; i++ {
		requests <- i
	}
	close(requests)

	// WaitGroup para aguardar todas as goroutines
	var wg sync.WaitGroup

	// Contadores atômicos
	var successCount int64
	statusCodes := make(map[int]int)
	var mu sync.Mutex // Mutex para proteger o map

	// Função worker
	worker := func() {
		defer wg.Done()
		for range requests {
			resp, err := http.Get(*url)
			if err != nil {
				// Tratamento de erro, pode-se considerar status code 0 ou outro valor
				mu.Lock()
				statusCodes[0]++
				mu.Unlock()
				continue
			}
			// Lendo o status code
			status := resp.StatusCode
			resp.Body.Close()

			// Atualizando contadores
			if status == 200 {
				atomic.AddInt64(&successCount, 1)
			}
			mu.Lock()
			statusCodes[status]++
			mu.Unlock()
		}
	}

	// Iniciando as goroutines
	wg.Add(*concurrency)
	for i := 0; i < *concurrency; i++ {
		go worker()
	}

	// Aguardando todas as goroutines finalizarem
	wg.Wait()

	// Calculando o tempo total
	totalTime := time.Since(startTime)

	// Gerando o relatório
	fmt.Println("===== Relatório de Teste de Carga =====")
	fmt.Printf("URL Testada: %s\n", *url)
	fmt.Printf("Total de Requests: %d\n", *totalRequests)
	fmt.Printf("Concorrência: %d\n", *concurrency)
	fmt.Printf("Tempo Total: %v\n", totalTime)
	fmt.Printf("Requests com Status 200: %d\n", successCount)
	fmt.Println("Distribuição de Status Codes:")
	for code, count := range statusCodes {
		if code == 0 {
			fmt.Printf("  Erros de Conexão: %d\n", count)
		} else {
			fmt.Printf("  %d: %d\n", code, count)
		}
	}
}
