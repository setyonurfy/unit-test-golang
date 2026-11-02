# Go Unit Test

Project latihan belajar **unit testing di Golang**. Isinya contoh penggunaan paket `testing` bawaan Go, library [testify](https://github.com/stretchr/testify) (`assert`, `require`, dan `mock`), serta `benchmark`.

## Struktur Project

```
go-unit-test/
├── entity/
│   └── category.go                   # Struct Category (Id, Name)
├── helper/
│   ├── hello_world.go                # Fungsi HelloWorld dan HelloTayo
│   └── helllo_world_test.go          # Berbagai contoh test dan benchmark
├── repository/
│   ├── category_repository.go        # Interface CategoryRepository
│   └── category_repository_mock.go   # Mock repository (testify/mock)
├── service/
│   ├── category_service.go           # CategoryService.Get(id)
│   └── category_service_test.go      # Unit test dengan mock
├── go.mod
└── go.sum
```

## Prasyarat

- [Go](https://go.dev/dl/) versi 1.21.5 atau lebih baru

## Instalasi

```bash
git clone <url-repository-kamu>
cd go-unit-test
go mod tidy
```

`go mod tidy` akan mengunduh dependensi seperti `testify` dan merapikan `go.mod`.

## Menjalankan Test

Semua test di seluruh package:

```bash
go test ./...
```

Dengan output rinci (verbose):

```bash
go test -v ./...
```

Satu package saja:

```bash
go test -v ./helper
go test -v ./service
```

Satu fungsi test tertentu:

```bash
go test -v -run TestHelloWorld ./helper
```

Satu subtest tertentu:

```bash
go test -v -run TestSubTest/subtest1 ./helper
```

Dengan laporan cakupan (coverage):

```bash
go test -cover ./...
```

## Menjalankan Benchmark

Semua benchmark (test biasa dilewati):

```bash
go test -v -run=NONE -bench=. ./helper
```

Satu benchmark tertentu:

```bash
go test -v -run=NONE -bench=BenchmarkHelloWorld ./helper
```

Dengan informasi alokasi memori:

```bash
go test -v -run=NONE -bench=. -benchmem ./helper
```

## Materi yang Dicakup

### Package `helper`

Contoh dasar unit test pada fungsi sederhana:

| Materi | Fungsi |
| --- | --- |
| Test dasar dengan `t.Error` | `TestHelloWorld` |
| Menghentikan test dengan `t.Fatal` | `TestHelloTayo` |
| Assertion dengan `assert` | `TestHelloWorldAssert` |
| Assertion dengan `require` (menghentikan test saat gagal) | `TestHelloWorldRequire` |
| Melewati test dengan `t.Skip` | `TestSkip` |
| Kode sebelum dan sesudah semua test dengan `TestMain` | `TestMain` |
| Subtest dengan `t.Run` | `TestSubTest` |
| Table test | `TestTableHelloWorld` |
| Benchmark, sub-benchmark, dan table benchmark | `BenchmarkHelloWorld`, `BenchmarkHelloTayo`, `BenchmarkSub`, `BenchmarkTable` |

### Package `service` dan `repository`

Contoh **mocking** dengan `testify/mock`:

- `CategoryRepository` adalah interface dengan method `FindById(id string) *entity.Category`.
- `CategoryRepositoryMock` adalah implementasi palsu dari interface tersebut, dipakai khusus untuk test.
- `CategoryService.Get(id)` memanggil repository. Jika kategori tidak ditemukan, hasilnya `nil` dan error `category not found`.

Contoh penggunaan mock pada test:

```go
var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	// Atur mock: FindById("1") mengembalikan nil
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}
```

## Catatan

- `TestSkip` sengaja dilewati selain di Windows (`runtime.GOOS != "windows"`).
- `TestMain` menampilkan teks "Before test" dan "After test" di sekitar semua test dalam package `helper`.

## Dependensi

- [github.com/stretchr/testify](https://github.com/stretchr/testify) (`assert`, `require`, `mock`)