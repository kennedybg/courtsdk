## Como contribuir

Para contribuir, basta seguir os seguintes passos:

Seguir as instruções do tópico [Instalação](INSTALLATION.md)

Acessar o diretório do projeto
```
cd courtsdk
```

Crie uma branch
```
git branch nome-da-sua-branch-ou-feature
```

Mude para sua branch
```
git checkout nome-da-sua-branch-ou-feature
```

Após fazer as alterações que precisar, adicione as alterações para staging:

Para adicionar todas:
```
git add .
```

Para adicionar um arquivo específico.
```
git add seu_arquivo.extensao
```

Com os arquivos em staging, faça o commit e push:

```
git commit -m 'Sua mensagem de commit, ex: TASK-01: Feature X'
git push origin nome-da-sua-branch-ou-feature
```

Abra o Merge Request em
```
https://gitlab.com/equipe-ninja/BaaS/courtsdk/merge_requests/new
```

ou pelo link gerado após o comando git push acima.

* Descreva suas alterações na página de Merge Request.
* Atribua a opção Assignee a você.
* Defina a label para Code Review.
* Envie seu Merge Request.
* Copie o link e compartilhe no canal code-review do Teams.

\#VQV
