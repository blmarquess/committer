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

---

# commiter

## Config file

```json
{
  "name": "default",
  "template": "$commit_type ($scope): $commit_message \n\n $commit_long_message \n\n #$issue",
  "steps": [
    {
      "name": "commit_type",
      "type": "select",
      "options": [
        { "message": "feat", "value": "feat" },
        { "message": "feat", "value": "feat" },
        { "message": "feat", "value": "feat" },
        { "message": "feat", "value": "feat" },
        { "message": "feat", "value": "feat" }
      ]
    },
    {
      "name": "commit_message",
      "type": "input",
      "label": "Short Description:",
      "optional": false
    },
    {
      "name": "scope",
      "type": "input",
      "format": "($)",
      "optional": true,
      "question": "Add scope for this?",
      "label": "Scope:"
    },
    {
      "name": "commit_long_message",
      "type": "input",
      "optional": true,
      "question": "Add long message?",
      "label": "Long message:"
    },
    {
      "name": "issue",
      "type": "input",
      "format": "issue: #$",
      "optional": true,
      "question": "This commit resolves a issue?",
      "label": "issue(only number)"
    }
  ]
}
```

// por padrão todos os imputes são obrigatórios se precisar deixar opcional tem que adicionar valor que é opcional

// format: pode adicionar um padrão para aplicar ao input usando o input como variável pelo carácter $

## TODO Next steps

- [ ] gen cli questions by config file
- [ ] add view with select files to add a staged area
- [ ] add handler to git errors
- [ ] fix readme.md
- [ ] fix bug config directory
