# GoLogging

## Como usar?

### Instalação

Para que você possa instalar o conector basta executar um **go get** no pacote:

```go
 go get bitbucket.org/stonepayments/gologging
```

### Utilização

**Exemplo:**
```go
package main
import bitbucket.org/stonepayments/gologging

func init() {
    gologging.SetOutput("path/to/log/file.log")
}

func main() {
    logger := gologging.New("Main")
    logger.Info("info message", "extra", "parameters")
    logger.Debug("debug message", "extra", "parameters")
    logger.Warn("warn message", "extra", "parameters")
    logger.Error("error message", "extra", "parameters")
    logger.Fatal("fatal message", "extra", "parameters")
}

```

## Responsáveis imediatos

    Mateus Rodrigues - mralves@stone.com.br

    Rafael Portugal - rportugal@stone.com.br


## Como contribuir?

Todo o processo necessário para você contribuir encontra-se em nosso [CONTRIBUTING](https://stonepayments.atlassian.net/wiki/display/RC/Contributing+para+projetos+Go).
