Dim Arg
Arg = WScript.Arguments.Item(0)
Set objShell = CreateObject("Shell.Application")
objShell.ShellExecute "master.exe", Arg, "E:\SourceCode\gpc\bin", "runas", 0