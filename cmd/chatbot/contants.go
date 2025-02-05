package chatbot

const (
	TO_BINARY = iota
	TO_TEXT
	TO_DECIMAL
	GENERATE_PASSWORD
	QUIT
)

const (
	SYSTEM int = iota
	CONVERTER
	PASSWORD
)

var options = []Option{
	{CONVERTER, TO_BINARY, "Convertir texte en binaire", "Entrez le texte : ", "✅ Texte en binaire :", &ParseText{}},
	{CONVERTER, TO_TEXT, "Convertir binaire en texte", "Entrez le binaire : ", "✅ Binaire en texte :", &ParseBinary{}},
	{CONVERTER, TO_DECIMAL, "Convertir binaire en décimal", "Entrez le binaire : ", "✅ Binaire en décimal :", &ParseBinary{}},
	{CONVERTER, TO_BINARY, "Convertir décimal en binaire", "Entrez le nombre décimal : ", "✅ Décimal en binaire :", &ParseInt{}},
	{PASSWORD, GENERATE_PASSWORD, "Générer un mot de passe", "Entrez la longueur du mot de passe : ", "✅ Mot de passe généré :", &ParseUint{}},
	{SYSTEM, QUIT, "Quitter", "", "", nil},
}
