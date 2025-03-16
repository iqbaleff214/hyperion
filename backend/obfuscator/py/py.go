package py

type Obfuscation struct{}

func (o Obfuscation) Minify(code string) (string, error) {
	return code, nil
}

func (o Obfuscation) RefactorLoopStatement(code string) (string, error) {
	return code, nil
}

func (o Obfuscation) RefactorIfStatement(code string) (string, error) {
	return code, nil
}

func (o Obfuscation) RenameVariable(code string) (string, error) {
	return code, nil
}

func (o Obfuscation) RenameConstant(code string) (string, error) {
	return code, nil
}

func (o Obfuscation) RenameFunction(code string) (string, error) {
	return code, nil
}

func (o Obfuscation) RemoveComment(code string) (string, error) {
	return code, nil
}
