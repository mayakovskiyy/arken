package ggo

import (
	"log"

	"github.com/jcalabro/atmos/lexgen"
	"github.com/jcalabro/atmos/lexicon"
)

func main() {
	schemas, err := lexicon.ParseDir(".")
	if err != nil {
		log.Fatalf("failed  to parse lexicon dir: %v", err)
	}

	cat := lexicon.NewCatalog()
	if err := cat.AddAll(schemas); err != nil {
		log.Fatalf("failed to add schemas to catalog: %v", err)
	}
	if err := cat.Resolve(); err != nil {
		log.Fatalf("failed to resolve catalog: %v", err)
	}

	cfg := &lexgen.Config{
		SharedTypesDir:    "api/lextypes",
		SharedTypesPkg:    "lextypes",
		SharedTypesImport: "arken/api/lextypes",
		Packages: []lexgen.PackageConfig{
			{
				Prefix:  "fun.arken",
				Package: "arken",
				OutDir:  "api/arken",
				Import:  "arken/api/arken",
			},
		},
	}

	files, err := lexgen.Generate(cfg, cat)
	if err != nil {
		log.Fatalf("failed to generate code: %v", err)
	}

	if err := lexgen.WriteFiles(files); err != nil {
		log.Fatalf("failed to write generated files: %v", err)
	}
}
