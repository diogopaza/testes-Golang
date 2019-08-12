<h1>Testes em Go </h1>
<p>Go incorpora em sua filosofia a ideia de testes. Já tem bibliotecas nativas.</p>
<p>Os arquivos usados para testes em Go devem ter o sufixo <strong>_test.go</strong>, neste artigo o arquivo de testes se chama testes_test.go</p>
<p>A função deve começar com a palavra Test seguida do restante do nome para o Go interpretar que é um teste. Por exemplo TestSoma.</p>
<p>O comando para executar teste é o <strong>go test</strong></p>
<h2>Cover</h2>
<p>Essa ferramenta serve para medir o quanto do codigo esta sendo testado. </p>
<h3>GOPATH</h3>
<p>Go precisa de um caminho absoluto para o ambiente do projeto</p>
<p><strong>export GOPATH={PWD}</strong> com esse comando dentro um diretorio como por exemplo meu-projeto/src/testes para definir a variavel de ambiente.</p>
<p>go get golang.org/x/tools/cmd/cover comando para instalar a ferramenta cover</p>
<p>go test funcoes -cover == retorna a porcentagem do codigo qu foi testada</p>
<p>go test -coverprofile=c.out funcoes == cria arquivo com a exibicao da parte do codigo que nao foi testada</p>
<p>go tool cover -html=c.out == mostra no navegador um arquivo com o copdigo nao testado</p>
<h1>Benchmark</h1>
<p>O Go tambem traz uma ferramenta de benchmark para teste de desempenho, que retorna o tempo e também pode retornar a quantidade de memória usada por cada função testada.</p>
<p>A função deve começar com a palavra Benchmark seguida do restante do nome para o Go interpretar que é um teste de benchmark. Por exemplo BenchmarkSoma.</p>
<p>O comando go test -bench=. faz o benchmark retornando as funcoes que foram testadas com o tempo de execução de cada uma delas.</p>
<p>O comando go test -bench=. -benchmem retorna a quantidade de memória utilizada em cada teste.</p>



