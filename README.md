# hypr-dock — fork de Leonardo Gouveia

Fork de [lotos-linux/hypr-dock](https://github.com/lotos-linux/hypr-dock), um dock
em Go e GTK3 para Hyprland. O objetivo é desenvolver melhorias incrementais de
organização e interação, mantendo uma relação clara com o projeto original.

## Estado deste fork

**Adaptação Lua disponível na branch `feature/hyprland-lua`.** Ela foi migrada do
checkout experimental, cuja instalação foi retirada. O fork não é instalado
automaticamente. A base original é `468590b84fbce53911a736a266f516861aabf171`.

Ambiente validado: Hyprland 0.57.2 com configuração Lua, janelas flutuantes e
Waybar existente. A adaptação é opt-in: `HYPR_DOCK_LUA=1` traduz os comandos
legados de foco e fechamento para dispatchers Lua. Sem essa variável, esses
comandos permanecem no formato original.

Ao ativar pelo dock, a janela recebe foco e é elevada, preservando a posição do
cursor. A supressão de warps vale somente durante a ação Lua síncrona; a
preferência anterior é restaurada mesmo em caso de erro. A comunicação no
socket tem prazo de dois segundos e fecha a conexão também em caso de falha.
Isso não adiciona suporte Lua genérico a todos os comandos ou ferramentas do
repositório; foco e fechamento são os comandos traduzidos nesta etapa.

## Validação da adaptação

Teste de pacote sem sessão gráfica (o teste gráfico é ignorado por padrão):

```sh
go test ./pkg/ipc
```

Teste gráfico explícito, executado dentro do Hyprland:

```sh
python3 tests/integration/hypr-dock-focus-test.py
```

O teste cria e fecha duas janelas descartáveis sobrepostas, usa a entrada IPC
real do aplicativo e verifica foco, sobreposição por amostra de pixels, posição
do cursor e restauração de `cursor.no_warps`. O cursor é levado temporariamente
à borda inferior para reproduzir a ativação pelo dock e restaurado ao final.
Requer Go, Python com PyGObject/GTK3, `grim`, `hyprctl` e monitor sem rotação com
escala 1; não é executado automaticamente em CI. `GO_BINARY` permite selecionar
outro executável Go. A confirmação de funcionamento não implica cobertura de
todos os comportamentos de mouse, configurações de foco ou múltiplos monitores.

A prévia estática foi observada funcionando na instalação experimental anterior.
O teste automatizado atual cobre ativação; não automatiza cliques GTK reais nem
valida o fechamento de aplicativos. Autostart, minimização e atualização com
reversão ainda não foram implementados neste fork.

## Melhorias planejadas — ainda não implementadas

- Separação visual entre favoritos e aplicativos abertos não fixados.
- Atualização da organização ao fixar ou desafixar um aplicativo.
- Indicação do aplicativo ativo, distinta da quantidade de janelas abertas.
- Reordenação de favoritos por arrasto, com persistência da ordem.
- Processo simples de atualização, validação e reversão do executável.

Cada funcionalidade deverá ser tratada separadamente. A inclusão nesta lista
não significa que ela já esteja disponível ou tenha sido testada.

## Organização e manutenção

- `origin`: [leogouveia/hypr-dock](https://github.com/leogouveia/hypr-dock).
- `upstream`: [lotos-linux/hypr-dock](https://github.com/lotos-linux/hypr-dock).
- Preservar a autoria, a licença e a documentação original.
- Desenvolver mudanças em branches específicas, com commits pequenos e testes
  proporcionais; evitar misturar aparência pessoal com correções do aplicativo.
- Manter as configurações pessoais fora do código-fonte.
- Avaliar atualizações do upstream antes de incorporá-las; não substituir o
  executável em uso sem uma etapa explícita de instalação e reversão.
- Preparar correções gerais para contribuição ao upstream quando solicitado.

### Gitflow do fork

| Branch | Finalidade |
| --- | --- |
| `main` | Versões estáveis; sem as adaptações ainda em desenvolvimento. |
| `develop` | Base e integração das funcionalidades do fork. |
| `feature/*` | Uma funcionalidade por branch, criada a partir de `develop`. |
| `release/*` | Preparação de versões para integração em `main` e `develop`. |
| `hotfix/*` | Correções urgentes a partir de `main`, integradas também em `develop`. |

As adaptações atuais estão em `feature/hyprland-lua` e ainda não foram
integradas em `develop` ou `main`. `develop` foi criada na base do fork anterior
às adaptações. Branches herdadas do projeto original não são usadas para nosso
desenvolvimento, e o remoto `upstream` não recebe nossos pushes.

Fluxo de uma próxima funcionalidade, com o diretório de trabalho limpo:

```sh
git switch develop
git switch -c feature/nome-da-funcionalidade
```

Revisão, validação e integração em `develop` são etapas posteriores. Não é
necessário instalar uma extensão Gitflow para seguir essa organização.

O clone de desenvolvimento fica em `~/Developer/hypr-dock-fork`. O antigo
checkout experimental `~/Developer/hypr-dock` e sua instalação de teste foram
retirados após a migração. Preferências pessoais foram preservadas em backup
local e não fazem parte do Git.

As orientações para trabalhar no repositório estão em [AGENTS.md](AGENTS.md).

## Documentação original do upstream

O conteúdo abaixo foi preservado como referência do projeto original. Os
comandos de instalação, os exemplos na sintaxe tradicional do Hyprland e o
pacote AUR referem-se ao upstream; não representam uma distribuição própria
deste fork. Para as ações Lua documentadas acima, é necessário optar por
`HYPR_DOCK_LUA=1`. Nenhuma instalação é executada pelos testes descritos aqui.

---

# hypr-dock
### Interactive dock panel for Hyprland

Translations: [`Русский`](README_RU.md)

<img width="1360" height="768" alt="250725_16h02m52s_screenshot" src="https://github.com/user-attachments/assets/041d2cf6-13ba-4c89-a960-1903073ff2d4" />
<img width="1360" height="768" alt="250725_16h03m09s_screenshot" src="https://github.com/user-attachments/assets/0c1ad8ca-37c1-4fd6-a48d-46f74c2d2609" />

[![YouTube](https://img.shields.io/badge/YouTube-Video-FF0000?logo=youtube)](https://youtu.be/HHUZWHfNAl0?si=ZrRv2ggnPBEBS5oY)
[![AUR](https://img.shields.io/badge/AUR-Package-1793D1?logo=arch-linux)](https://aur.archlinux.org/packages/hypr-dock)

## Installation

### Dependencies

- `go` (make)
- `gtk3`
- `gtk-layer-shell`

### Installation
! The first build may take a very long time due to gtk3 bindings !
```bash
git clone https://github.com/lotos-linux/hypr-dock.git
cd hypr-dock
make get
make build
make install
```

### Uninstallation
```bash
make uninstall
```

### Local run (dev mode)
```bash
make exec
```

## Running

### Launch parameters:

```text
  -config string
    	config file (default "~/.config/hypr-dock")
  -dev
    	enable developer mode
  -log-level string
    	log level (default "info")
  -theme string
    	theme dir
```
#### All parameters are optional.

Default configuration and themes are installed in `/etc/hypr-dock`
On first run, they are copied to `~/.config/hypr-dock`
### Add to `hyprland.conf`:

```text
exec-once = hypr-dock
bind = Super, D, exec, hypr-dock
```

### And configure blur if needed
```text
layerrule = blur true,match:namespace hypr-dock
layerrule = ignore_alpha 0,match:namespace hypr-dock
layerrule = blur true,match:namespace dock-popup
layerrule = ignore_alpha 0,match:namespace dock-popup
```

#### The dock only supports one running instance, so running it again will close the previous one.

## Configuration

#### !Important!
If you already had `hypr-dock` installed or you are updating, please note that starting from version 1.2.0, an `ini`-like configuration format is now used. If you have old configs remaining in your user folder, be sure to delete or move them. Then simply adapt your settings to the new format.

### Available parameters in `hypr-dock.conf`

```ini
[General]
CurrentTheme = lotos

# Icon size (px) (default 23)
IconSize = 23

# Window overlay layer height (background, bottom, top, overlay) (default top)
Layer = top

# Exclusive Zone (true, false) (default true)
Exclusive = true

# SmartView (true, false) (default false)
SmartView = false

# Window position on screen (top, bottom, left, right) (default bottom)
Position = bottom

# Delay before hiding the dock (ms) (default 400)
AutoHideDelay = 400   # Only for SmartView

# Use system gap (true, false) (default true)
SystemGapUsed = true

# Indent from the edge of the screen (px) (default 8)
Margin = 8

# Distance of the context menu from the window (px) (default 5)
ContextPos = 5

[General.preview]
# Window thumbnail mode selection (none, live, static) (default none)
Mode = none
# "none"   - disabled (text menus)
# "static" - last window frame
# "live"   - window streaming
      
# !WARNING! 
# BY SETTING "Mode" TO "live" OR "static", YOU AGREE TO THE CAPTURE 
# OF WINDOW CONTENTS.
# THE "HYPR-DOCK" PROGRAM DOES NOT COLLECT, STORE, OR TRANSMIT ANY DATA.
# WINDOW CAPTURE OCCURS ONLY FOR THE DURATION OF THE THUMBNAIL DISPLAY!
#   
# Source code: https://github.com/lotos-linux/hypr-dock

# Live preview fps (0 - ∞) (default 30)
FPS = 30

# Live preview bufferSize (1 - 20) (default 5)
BufferSize = 5

# Popup show/hide/move delays (ms)
ShowDelay = 500  # (default 500)
HideDelay = 350  # (default 350)
MoveDelay = 100  # (default 100)
```
#### If a parameter is not specified, the default value will be used

## Understanding non-obvious parameters

### SmartView
Similar to auto-hide: if `true`, the dock stays beneath all windows, but moves the cursor to the screen edge - the dock rises above them

### Exclusive
Activates special layer behavior where tiling windows do not overlap the dock

### SystemGapUsed
- When `SystemGapUsed = true`, the dock sets its margin from the screen edge using values from the `hyprland` configuration, specifically `general:gaps_out` values, and dynamically updates when the `hyprland` configuration changes
- When `SystemGapUsed = false`, the margin from the screen edge is set by the `Margin` parameter

### General.preview
- `ShowDelay`, `HideDelay`, `MoveDelay` - preview popup action delays in milliseconds
- `FPS`, `BufferSize` - only used when `Mode = live`

#### Preview appearance settings are configured through theme files

### Pinned applications are stored in `~/.local/share/hypr-dock/pinned`
To pin/unpin, open the application's context menu in the dock and click `pin`/`unpin`
#### Example
```text
firefox
code-oss
kitty
org.telegram.desktop
nemo
org.kde.ark
sublime_text
qt6ct
one.ablaze.floorp
```
You can edit it manually. But why? ¯\_(ツ)_/¯

## Themes

#### Themes are located in `~/.config/hypr-dock/themes/`

### A theme consists of
- `theme.conf`
- `style.css`
- A folder with `svg` files for indicating the number of running applications (see [themes_EN.md](https://github.com/lotos-linux/hypr-dock/blob/main/docs/customize/themes_EN.md))

### Theme config
```ini
[Theme]
# Distance between elements (px) (default 9)
Spacing = 5


[Theme.preview]
# Size (px) (default 120)
Size = 120

# Image/Stream border-radius (px) (default 0)
BorderRadius = 0

# Popup padding (px) (default 10)
Padding = 10
```
#### Customize `style.css` as you wish. Detailed styling documentation will be provided later.
