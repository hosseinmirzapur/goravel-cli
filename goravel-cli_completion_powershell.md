## goravel-cli completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	goravel-cli completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
goravel-cli completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.goravel-cli.yaml)
```

### SEE ALSO

* [goravel-cli completion](goravel-cli_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 27-Dec-2023