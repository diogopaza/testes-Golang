<h1>Testes em Go </h1>
<p>Go incorpora em sua filosofia a ideia de testes. Já tem bibliotecas nativas.</p>
<p>Os arquivos usados para testes em Go devem ter o sufixo <strong>_test.go</strong>, neste artigo o arquivo de testes se chama testes_test.go</p>
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



