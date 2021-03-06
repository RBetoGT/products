package parsers

import (
	"testing"
)

func Test_script_Success(t *testing.T) {
	grammarFile := "./grammar/grammar.json"
	pars := createParserForTests("script", grammarFile)

	file := "./tests/codes/script/all.rod"
	ins, err := pars.ExecuteFile(file)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	script := ins.(Script)
	name := script.Name()
	if name != "my_name" {
		t.Errorf("the name was expected to be %s, %s returned", "my_name", script.Name())
		return
	}

	version := script.Version()
	if version != "2019.03.23" {
		t.Errorf("the version was expected to be %s, %s returned", "2019.03.23", script.Version())
		return
	}

	scriptPath := script.Script()
	if scriptPath.String() != "./my/script.rod" {
		t.Errorf("the script was expected to be %s, %s returned", "./my/script.rod", script.Script().String())
		return
	}

	langPath := script.Language()
	if langPath.String() != "./my/lang.rod" {
		t.Errorf("the language was expected to be %s, %s returned", "./my/lang.rod", script.Language().String())
		return
	}
}
