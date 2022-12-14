Pasta chamada: "challenger-algorithm-go" e dentro dela vai ter outra pasta

challenger-algorithm-go
- model
  - model.go
main.go

dentro da pata challenger-algorithm-go vai ter o seu arquivo principal main.go
e na mesma pasta tu vai criar a pasta model com o arquivo chamado model.go
esse model é onde fica suas struct
não quero struct dentro do arquivo main.go
e uma pasta chamada person
com um arquivo chamado person.go
model/model.go -> onde fica suas estrutura de dados no caso as struct
person/person.go -> onde via ficar as funções que vai receber os inputs e processar tudo.
main.go -> onde vai ser executado tudo, e chamar as funçõe.
isso se chama "Arquitetura"
dentro de person quero a função que vai receber um nome e uma idade e adicionar a uma lista
no seu main dps que adicionar todos os nomes e idade quero que ordene essa lista do mais velho pro mais novo
exemplo:

João
19 anos

paulo
20 anos

witalo
25 anos

carlos
30 anos

-----------------------------


"tipo" deve ser um dos seguintes mencionados abaixo!
build: Crie alterações relacionadas (por exemplo: relacionado ao npm/adicionando dependências externas)
chore: Uma alteração de código que o usuário externo não verá (por exemplo: alterar para arquivo .gitignore ou arquivo .prettierrc)
feat: Um novo recurso
fix: Uma correção de bug
docs: Alterações relacionadas à documentação
refactor: Um código que não corrige bug nem adiciona um recurso. (por exemplo: você pode usar isso quando houver alterações semânticas, como renomear uma variável/nome de função)
perf: Um código que melhora o desempenho
style: Um código relacionado ao estilo
test: Adicionando novo teste ou fazendo alterações no teste existente