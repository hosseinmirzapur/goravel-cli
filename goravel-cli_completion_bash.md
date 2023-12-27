## goravel-cli completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(goravel-cli completion bash)

To load completions for every new session, execute once:

#### Linux:

	goravel-cli completion bash > /etc/bash_completion.d/goravel-cli

#### macOS:

	goravel-cli completion bash > $(brew --prefix)/etc/bash_completion.d/goravel-cli

You will need to start a new shell for this setup to take effect.


```
goravel-cli completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.goravel-cli.yaml)
```

### SEE ALSO

* [goravel-cli completion](goravel-cli_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 27-Dec-2023