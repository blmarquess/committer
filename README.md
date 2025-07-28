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
