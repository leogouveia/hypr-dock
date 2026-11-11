# Orientações para trabalhar neste fork

## Contexto e estado inicial

Este repositório é o fork `leogouveia/hypr-dock`, baseado em
`lotos-linux/hypr-dock`. Na preparação inicial, apenas README.md e este arquivo
foram alterados. Não tratar os objetivos do README como funcionalidades prontas.

A etapa inicial não autoriza instalação, compilação, migração de configurações
ou aplicação das adaptações existentes em outros checkouts. Trabalhos futuros
devem seguir o escopo da solicitação vigente do usuário.

## Forma de trabalhar

- Responder em português brasileiro e explicar mudanças de comportamento.
- Usar os arquivos e o código real como fonte de verdade; não inventar APIs.
- Preferir mudanças pequenas e localizadas, em uma funcionalidade por vez.
- Preservar alterações preexistentes e não substituir arquivos inteiros sem
  necessidade. Não copiar automaticamente código de outro checkout.
- Ler `git status`, os remotes e o diff antes de editar ou integrar mudanças.
- Preservar licença, créditos e referências ao projeto original.
- Manter configurações pessoais, binários, caches, logs e capturas fora do Git.

## Mapa do projeto

- `cmd/hypr-dock/`: entrada do dock.
- `cmd/hypr-alttab/`: ferramenta separada de alternância de janelas.
- `internal/app/`: composição e atualização da lista de aplicativos.
- `internal/item/` e `internal/btnctl/`: itens, ações e interações.
- `internal/layering/` e `internal/detectZone/`: camadas e ativação pela borda.
- `internal/pvctl/`, `internal/pvwidget/` e `internal/hysc/`: prévias e captura.
- `pkg/ipc/`: comunicação com o Hyprland.
- `configs/default/`: configurações e tema distribuídos pelo projeto.

## Compatibilidade e limites

- A base usa Go e GTK3; conferir requisitos em `go.mod` e no Makefile.
- O ambiente-alvo inicial usa Hyprland 0.57.x com Lua e janelas flutuantes.
  Verificar a API da versão efetivamente instalada antes de mudar o IPC.
- Não presumir que comandos antigos de dispatch funcionam no ambiente-alvo.
- Preservar Waybar, atalhos, foco, posição das janelas e configurações externas.
- Evitar mudanças globais de teclado, mouse, GTK, temas ou energia para resolver
  um comportamento exclusivo do dock.
- Não recriar os binds `XF86AudioMicMute`, `XF86MonBrightnessUp` ou
  `XF86MonBrightnessDown`; foram removidos do ambiente por problemas de entrada.
- Não suspender o computador nem executar testes que desliguem o monitor.
- Não instalar nem iniciar o hypr-alttab como consequência de alterar o dock.

## Branches, atualização e contribuições

- `origin` deve apontar para o fork; `upstream`, para o projeto original.
- Criar uma branch por correção ou funcionalidade quando começar implementação.
- Manter commits de compatibilidade, recursos e documentação separados quando
  forem mudanças independentes.
- Não usar force-push, rebase destrutivo ou sobrescrever trabalho local alheio.
- Atualizações devem ser preparadas e verificadas antes de substituir qualquer
  instalação. Se houver conflito ou falha, preservar a versão em uso.
- Não fazer commits, pushes, PRs ou publicações adicionais fora da autorização
  da tarefa vigente. Criar um fork não autoriza enviar mensagens ao upstream.

## Validação

- Documentação: conferir links, caminhos, comandos e consistência com o estado
  real; `git diff --check` é suficiente quando não há alteração de código.
- Código Go: aplicar `gofmt` aos arquivos alterados e usar testes adequados ao
  comportamento, sem exigir sessão gráfica para testes comuns.
- Testes gráficos devem ser explícitos e usar janelas descartáveis. Verificar
  tanto o foco quanto a sobreposição visual, além da posição do cursor.
- Não fechar, minimizar, mover ou redimensionar aplicativos reais do usuário
  para testar; não capturar o desktop inteiro sem necessidade.
- Registrar separadamente o que foi testado automaticamente, observado na tela
  e ainda depende de validação do usuário.
- Atualizar o README quando uma funcionalidade de fato passar a existir; não
  apresentar intenção, compilação bem-sucedida ou resposta IPC como prova de
  funcionamento completo.
