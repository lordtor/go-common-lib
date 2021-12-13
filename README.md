# Go common lib

[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=lordtor_go-common-lib&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=lordtor_go-common-lib) [![Bugs](https://sonarcloud.io/api/project_badges/measure?project=lordtor_go-common-lib&metric=bugs)](https://sonarcloud.io/summary/new_code?id=lordtor_go-common-lib) [![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=lordtor_go-common-lib&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=lordtor_go-common-lib) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=lordtor_go-common-lib&metric=coverage)](https://sonarcloud.io/summary/new_code?id=lordtor_go-common-lib)

Tis module used as smaple module for most popular func.

## ENV parameters

* N/a

## File parameters `application.yml`

* N/a

## Lint, Test & Coverage

  1. Install linter

      ``` bash
      go install github.com/golangci/golangci-lint/cmd/golangci-lint
      ```

      run check

      ``` bash
      golangci-lint run  > checkstyle.xml
      ```

  2. Run test

      ``` bash
      go test -v -coverpkg=./... -coverprofile=profile.cov ./... -json > test_report.json
      ```

  3. Coverage

      ``` bash
      go tool cover -func profile.cov
      ```

## Methods

| Method | Description | Params | Result |
|--------|-------------|--------|--------|
| SliceContain | A function to check slice contain item string | `slice []string, item string` |  `bool` |
| UpdateList | A function for update exist slice not contain item string | `slice []string, item string` | `[]string` |
| StringFromList | A function convert slice to sting | `list []string` | `st string` |
| UpdateStructList | A function for update exist slice values from anothe slice | `main []string, in []string` | `[]string` |
| CalcHash | A function for calculate string hash sum by sha256 | `aStr string` | `string` |
| SortedMap | A function for sort slice map's | `m []map[string]string` | `out []map[string]string` |
| ListConvert | A function for convert interface to slice strings | `in interface{}` | `[]string, error` |
| UpdateMapList |  A function for update slince map by not exist map | `sliceMaps []map[string]string, itemMap map[string]string` | `[]map[string]string` |
| StandardizeSpaces | A function for trim first & last space | `s string` | `string` |
| UpdateListBySplit | A function for update slice string by splited values use slice spliters string | `slice []string, item string, spliters []string` | `[]string` |
| MapListConvert | A function for convert interface to slice map | `in interface{}` | `[]map[string]string, error` |
