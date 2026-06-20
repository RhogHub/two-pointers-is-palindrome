# Two Pointers em Go: Validação de Palíndromos

Este projeto implementa a técnica **Two Pointers** para verificar se uma
palavra é um palíndromo.

## Algoritmo

A estratégia consiste em utilizar dois índices:

-   `left`: começa no início da string.
-   `right`: começa no final da string.

``` go
func isPalindrome(s string) bool {
    left := 0
    right := len(s) - 1

    for left < right {
        if s[left] != s[right] {
            return false
        }

        left++
        right--
    }

    return true
}
```

## Exemplo

Para a palavra `arara`:

  left   right   Comparação
  ------ ------- ------------
  a      a       ✓
  r      r       ✓
  a      a       fim

## Complexidade Temporal

Se a string possui tamanho `n`, o número aproximado de comparações é:

    n / 2

Logo, a complexidade temporal é:

    O(n)

## Complexidade Espacial

São utilizados apenas dois índices inteiros:

-   left
-   right

Portanto, a complexidade espacial é:

    O(1)

## Benchmark

Executando a função diversas vezes:

``` go
const Iterations = 100_000_000
```

O tempo médio é calculado por:

    tempo médio = tempo total / número de iterações

Exemplo:

    Tempo total: 4.3s
    Iterações: 100.000.000

    ≈ 43 ns/op

## Conclusão

A abordagem Two Pointers apresenta:

-   Complexidade temporal O(n)
-   Complexidade espacial O(1)
-   Apenas uma passagem pela string
-   Não cria estruturas auxiliares
-   Excelente desempenho para este tipo de problema
