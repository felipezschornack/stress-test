# Sistema para Teste de Stress

## Objetivo
Criar um sistema CLI em Go para realizar testes de carga em um serviço web. 
O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.
O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

## Entrada de Parâmetros via CLI
```
--url: URL do serviço a ser testado
--requests: Número total de requests
--concurrency: Número de chamadas simultâneas
```

## Execução do Teste

- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.

## Geração de Relatório

Apresentar um relatório ao final dos testes contendo:

- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Executando a aplicação diretamente com Go
```
go run . stress --url=https://google.com --requests=50 --concurrency=10
```

## Executando a aplicação utilizando Docker
Com o Docker instalado em sua estação de trabalho (https://www.docker.com/), execute o comando:
```
docker compose run stress --url=https://google.com --requests=50 --concurrency=10
```



