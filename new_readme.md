## Como instalar o Committer

O Committer é uma ferramenta desenvolvida para facilitar e padronizar o processo de criação de commits em projetos que utilizam o Git.

Com uma configuração simples e rápida, você poderá garantir que suas mensagens de commit sigam um padrão definido, tornando o histórico do seu projeto mais organizado e compreensível.

---

1. Realize a instalação do GO Lang

Para instalar o Go Lang, execute o seguinte comando no terminal:

```bash
sudo apt update -y
sudo apt install golang-go -y
```

2. Realize o clone do repositório: `git clone git@github.com:blmarquess/committer.git`
3. Acesse o diretório: `commiter/`
4. No diretório execute o comando: `go mod tidy`
5. Execute: `go build -o /bin/cz ./cmd/committer`
6. Após o build, deve-se copiar o `config_template_default.json` que encontra-se no caminho `internal/config/` para o diretório `~/.committer/`
7. Por fim, em `~/.committer/`, renomeie o `config_template_default.json` para `config.json`
