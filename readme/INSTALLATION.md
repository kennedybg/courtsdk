## Instalação local

### 1) Instalar Go lang:

Mac OSX:
```
brew install go
```
### 2) Clonar o repositorio:

```
git clone git@gitlab.com:equipe-ninja/BaaS/courtsdk.git
cd courtsdk
```

## Instalar como dependência de uma aplicação Go lang:

**Atenção:** Seu terminal precisa estar autenticado.

### 1) Adicionar (SDK GitHub):

```
Adicione a URL do repositório (sem HTTPs:// e sem o .git) no arquivo go.mod da sua aplicação
Caso a aplicação não possua, execute:
go mod init <NOME_DA_SUA_APLICACAO>
Agora adicione no arquivo go.mod recém criado.
```

### 2) Adicionar (SDK GiLab):

```
TODO
```

## Definir as variaveis de ambiente:

**Atenção:** As alterações de variáveis devem ser feitas na aplicação que implementa este SDK.

Ativar Go modules, caso contrário não será possível baixar as dependências nas versões corretas:
```
GO111MODULE=on
```

PS: Esses valores são padrões, exporte somente as variáveis que precisar alterar.

Variaveis de config do Elasticsearch.
Os valores definidos são padrões, configure de acordo com a necessidade.
```
ELASTIC_URL=http://localhost
ELASTIC_PORT=9200
ELASTIC_INDEX=jurisprudences_dev
ELASTIC_RETRY_CONNECTION=10
ELASTIC_RETRY_PING=5
```

Variaveis de comportamento padrão das Engines:
Os valores definidos são padrões, configure de acordo com a necessidade.
```
ENGINE_IS_ASYNC=TRUE
ENGINE_MAX_FAILURES=10
ENGINE_REQUESTS_PER_INTERVAL=10
ENGINE_REQUEST_DELAY=1000
ENGINE_REQUEST_TIMEOUT=25
ENGINE_GOROUTINE_RANGE=200
ENGINE_MAX_RECOVERIES=5
```

Variaveis de comportamento padrão do controle de Engines:
Os valores definidos são padrões, configure de acordo com a necessidade.

```
CONTROL_IS_CONCURRENT=FALSE
CONTROL_MAX_CONCURRENT_ENGINES=2
CONTROL_ACTION_DELAY=25
```

Caso seja necessário persistir essas variáveis, configure de acordo com o ambiente na aplicação que implementa o SDK, Ex:

* Dockerfile
* .bashrc

## Testar:

```
go test -v
```
