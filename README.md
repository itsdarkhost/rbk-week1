# Text Converter CLI

CLI утилита для обработки текста с поддержкой команд `(cap)`, `(up)`, `(low)`, `(hex)`, `(bin)` и нормализации пунктуации.

---

## Способы запуска

### 1. Запуск через `go run`

Напрямую из исходников:

```bash
go run ./cmd/cli test.txt output.txt
```


### 2. Сборка и запуск бинарника

Собери CLI как бинарник:

```bash
go build -o textconv ./cmd/cli
```

Запуск самого бинарника:

```bash
./textconv test.txt output.txt
```

---

## Входные данные

* `input` — путь к входному файлу
* `output` — путь к выходному файлу

Пример:

```bash
go run ./cmd/app test.txt result.txt
```