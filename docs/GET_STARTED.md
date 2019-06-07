# Começando

Para começar, importe este SDK para sua aplicação:

```Go
import "gitlab.com/equipe-ninja/BaaS/courtsdk"
```

# Control

A estrutura [Control](../structs.go) é responsável guardar todas as Engines registradas. Possui também [métodos](../control.go) para executar cada Engine.

Após importar o SDK para sua aplicação, é necessário criar um estrutura de controle:

```Go
package main

import (
	"gitlab.com/equipe-ninja/BaaS/courtsdk"
)

func main() {
	control := courtsdk.NewControl()
}
```

# Engine

A estrutura [Engine](../structs.go) representa um programa do tipo Crawler/Scraper. Implementa [métodos](../engine.go) para inicialização, configuração e comunicação com o Elasticsearch (detalhes na página Engine).

Após criarmos um Control, agora precisamos criar pelo menos uma Engine:

```Go
package main

import (
	"gitlab.com/equipe-ninja/BaaS/courtsdk"
)

func main() {
	control := courtsdk.NewControl()
    minhaEngine := courtsdk.NewEngine(
        courtsdk.Court("TRIBUNAL_EXEMPLO"),
		courtsdk.Base("BASE_DOCUMENTOS_EXEMPLO"),
		courtsdk.EntryPoint(MINHA_FUNCAO_ENGINE),
    )
}
```

Uma Engine possui várias configurações, porém apenas as três acima são obrigatórias:


* **COURT** = Nome do tribunal, será usado para compor o ID do documento no Elasticsearch.

* **BASE** = Nome da base de documentos, será usado para gerar o ID do documento e o DocumentType no Elasticsearch.

* **ENTRYPOINT** = Função de entrada para determinada Engine, é o que será executado quando o controle iniciar a Engine.


**Exemplos:**

```Go
courtsdk.Court("TST"),
courtsdk.Base("basePrecedentes"),
courtsdk.EntryPoint(precedentesNormativosEngine),
```

Resultará em documentos com padrão de ID:


* **TST-basePrecedentes-1**

* **TST-basePrecedentes-2**

* **TST-basePrecedentes-3**


```Go
courtsdk.Court("STF"),
courtsdk.Base("baseSumulas"),
courtsdk.EntryPoint(funcaoSumulas),
```
Resultará em documentos com padrão de ID:

* **STF-baseSumulas-1**

* **STF-baseSumulas-2**

* **STF-baseSumulas-3**

```Go
courtsdk.Court("STJ"),
courtsdk.Base("baseAcordaos"),
courtsdk.EntryPoint(acordaosCrawler),
```

Resultará em documentos com padrão de ID:

* **STJ-baseAcordaos-1**

* **STJ-baseAcordaos-2**

* **STJ-baseAcordaos-3**

Em todos os exemplos acima, a função definida como EntryPoint da engine, será a responsável por coletar as informações, o SDK apenas gerencia as execuções e recuperações.

**PS:** Os ID's numéricos são apenas exemplos, cabe a Engine definir como classificar os ID's.

# Registro

Após criarmos nossa Engine, devemos registra-lá na estrutura de controle para que ela possa ser executada:

```Go
package main

import (
	"gitlab.com/equipe-ninja/BaaS/courtsdk"
)

func main() {
	control := courtsdk.NewControl()
    minhaEngine := courtsdk.NewEngine(
        courtsdk.Court("TRIBUNAL_EXEMPLO"),
		courtsdk.Base("BASE_DOCUMENTOS_EXEMPLO"),
		courtsdk.EntryPoint(MINHA_FUNCAO_ENGINE),
    )
    control.Register(minhaEngine)
}
```

Dessa forma, a Engine "minhaEngine" agora está no array de Engines da estrutura Control.

# Iniciar

Após definir todas as Engines necessárias, basta iniciar o controle:

```Go
package main

import (
	"gitlab.com/equipe-ninja/BaaS/courtsdk"
)

func main() {
	control := courtsdk.NewControl()
    minhaEngine := courtsdk.NewEngine(
        courtsdk.Court("TRIBUNAL_EXEMPLO"),
		courtsdk.Base("BASE_DOCUMENTOS_EXEMPLO"),
		courtsdk.EntryPoint(MINHA_FUNCAO_ENGINE),
    )
    control.Register(minhaEngine)
    control.Start()
}
```


Pronto, sua aplicação será executada até que a Engine termine ou falhe em todos os níveis de recuperação da estrutura de controle.

Informações detalhadas na seção índice de conteúdo.

**Atualizado: 07/06/2019**
