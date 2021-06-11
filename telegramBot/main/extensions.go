package main

func GetExtension(Language string) string {
	switch Language {
	case "GNU C++17", "GNU C++14", "GNU C++11", "GNU C++", "MS C++ 2017", "MS C++", "GNU C++17 (64)":
		{
			return ".cpp"
		}
	case "GNU C11":
		{
			return ".c"
		}
	case "Node.js", "JavaScript":
		{
			return ".js"
		}
	case "Scala":
		{
			return ".scala"
		}
	case "Rust":
		{
			return ".rs"
		}
	case "Ruby 3":
		{
			return ".rb"
		}
	case "PyPy 3", "PyPy 2", "Python 3 + libs", "Python 3", "Python 2":
		{
			return ".py"
		}
	case "PHP":
		{
			return ".php"
		}
	case "Perl":
		{
			return ".pl"
		}
	case "PascalABC.NET", "FPC", "Delphi":
		{
			return ".pas"
		}
	case "Ocaml":
		{
			return ".ml"
		}
	case "Kotlin":
		{
			return ".kt"
		}
	case "Java 8", "Java 11":
		{
			return ".java"
		}
	case "Haskell":
		{
			return ".hs"
		}
	case "Go":
		{
			return ".go"
		}
	case "D":
		{
			return ".d"
		}
	case "Mono C#", ".NET Core C#":
		{
			return ".cs"
		}

	default:
		{
			return ".txt"
		}
	}
}
