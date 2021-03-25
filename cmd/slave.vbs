Dim Arg
Arg = WScript.Arguments.Item(0) + " " + WScript.Arguments.Item(1) + " " + WScript.Arguments.Item(2) + " " + WScript.Arguments.Item(3) + " " + WScript.Arguments.Item(4)
Set objShell = CreateObject("Shell.Application")
objShell.ShellExecute "slave.exe", Arg, "E:\SourceCode\gpc\bin", "runas", 0
