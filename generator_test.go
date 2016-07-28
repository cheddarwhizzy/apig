package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var userModel = &Model{
	Name: "User",
	Fields: []*Field{
		&Field{
			Name:        "ID",
			JSONName:    "id",
			Type:        "uint",
			Tag:         "",
			Association: nil,
		},
		&Field{
			Name:        "Name",
			JSONName:    "name",
			Type:        "string",
			Tag:         "",
			Association: nil,
		},
		&Field{
			Name:        "CreatedAt",
			JSONName:    "created_at",
			Type:        "*time.Time",
			Tag:         "",
			Association: nil,
		},
		&Field{
			Name:        "UpdatedAt",
			JSONName:    "updated_at",
			Type:        "*time.Time",
			Tag:         "",
			Association: nil,
		},
	},
}

var detail = &Detail{
	VCS:       "github.com",
	User:      "wantedly",
	Project:   "api-server",
	Model:     userModel,
	Models:    []*Model{userModel},
	ImportDir: "github.com/wantedly/api-server",
}

func compareFiles(f1, f2 string) bool {
	c1, _ := ioutil.ReadFile(f1)
	c2, _ := ioutil.ReadFile(f2)

	return bytes.Compare(c1, c2) == 0
}

func TestGenerateApibIndex(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateApibIndex")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateApibIndex(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "docs", "index.apib")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("API Blueprint index is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "docs", "index.apib")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate API Blueprint index correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

func TestGenerateApibModel(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateApibModel")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateApibModel(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "docs", "user.apib")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("API Blueprint model is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "docs", "user.apib")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate API Blueprint model correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

func TestGenerateSkeleton(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "copyStaticFiles")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(tempDir)

	outDir := filepath.Join(tempDir, "api-server")

	if err := generateSkeleton(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	files := []string{
		"README.md",
		".gitignore",
		"main.go",
		filepath.Join("db", "db.go"),
		filepath.Join("db", "pagination.go"),
		filepath.Join("router", "router.go"),
		filepath.Join("middleware", "set_db.go"),
		filepath.Join("server", "server.go"),
		filepath.Join("helper", "field.go"),
		filepath.Join("helper", "field_test.go"),
		filepath.Join("version", "version.go"),
		filepath.Join("version", "version_test.go"),
		filepath.Join("controllers", ".gitkeep"),
		filepath.Join("models", ".gitkeep"),
	}

	for _, file := range files {
		_, err := os.Stat(filepath.Join(outDir, file))
		if err != nil {
			t.Fatalf("Static file is not copied: %s", file)
		}
	}
}

func TestGenerateController(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateController")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateController(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "controllers", "user.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("Controller file is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "controllers", "user.go")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate controller correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

func TestGenerateREADME(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateREADME")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateREADME([]*Model{userModel}, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "README.md")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("README is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "README.md")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate README correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

func TestGenerateRouter(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateRouter")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateRouter(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "router", "router.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("Router file is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "router", "router.go")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate router correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}

func TestGenerateDB(t *testing.T) {
	outDir, err := ioutil.TempDir("", "generateDB")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateDB(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	path := filepath.Join(outDir, "db", "db.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("Router file is not generated: %s", path)
	}

	fixture := filepath.Join("testdata", "db", "db.go")

	if !compareFiles(path, fixture) {
		c1, _ := ioutil.ReadFile(fixture)
		c2, _ := ioutil.ReadFile(path)
		t.Fatalf("Failed to generate db.go correctly.\nexpected:\n%s\nactual:\n%s", string(c1), string(c2))
	}
}
