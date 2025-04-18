# Sonar-Cloud 🚀

## 📋 Resumo
Integração contínua (CI) é uma prática de DevOps que automatiza a construção, teste e integração de código sempre que há uma alteração no repositório.  
No processo de CI, cada commit no código-fonte é automaticamente verificado por uma pipeline de testes e build, garantindo que o código seja integrado frequentemente e de forma segura.

Essa prática ajuda a:
- Detectar erros mais rapidamente;
- Melhorar a qualidade do software;
- Acelerar o ciclo de desenvolvimento;
- Permitir uma entrega contínua e sem interrupções.

---

## 🛠️ Tecnologias Utilizadas

- **Go 1.23**
- **GitHub Actions**
- **SonarCloud**

---

## 🔧 Estrutura do Projeto

```bash
├── .github/workflows/ci.yaml     # Pipeline de CI com integração ao SonarCloud
├── sum.go                        # Função principal de soma
├── sum_test.go                   # Testes unitários da função
├── sonar-project.properties      # Configurações do SonarCloud
├── go.mod                        # Arquivo de dependências do Go
└── .gitignore                    # Arquivo de configurações do Git
```

---

## ⚙️ Pipeline de CI (GitHub Actions)

```yaml
name: ci-sonarcloud
on:
  pull_request:
    branches:
    - develop

jobs:
  check-application:
    runs-on: ubuntu-latest

    strategy:
      matrix:
        go: ['1.23']

    steps:

      - name: Check out repository
        uses: actions/checkout@v3

      - uses: actions/setup-go@v2
        with:
          go-version: ${{ matrix.go }}
          
      - run: go test -coverprofile=coverage.out

      - name: SonarCloud Scan
        uses: SonarSource/sonarcloud-github-action@master
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          SONAR_TOKEN: ${{ secrets.SONAR_TOKEN }}
```

---

## ✅ Testes

Arquivo `sum_test.go` com testes simples para a função `sum`:

```go
func TestSum(t *testing.T) {
	result := sum(2, 3)
	if result != 5 {
		t.Error("O resultado é diferente de 5")
	}
}
```

---

## 📌 Observações

- A branch monitorada pela pipeline é a `develop`.
- O token do SonarCloud (`SONAR_TOKEN`) precisa ser adicionado como segredo no repositório do GitHub.

---

## 📤 Como Executar

```bash
go run sum.go
```

## ✅ Como Testar

```bash
go test
```

