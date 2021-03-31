Dim Arg
Arg = WScript.Arguments.Item(0)
Set objShell = CreateObject("Shell.Application")
objShell.ShellExecute "master.exe", Arg, "C:\Users\rmstj\Documents\GitHub\gpc\cmd", "runas", 0