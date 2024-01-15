# Passerby

Passerby is a package for hide command-line arguments for Golang program.

## Features

- Hide command-line arguments from an environment variable
- Delete a file after a specified duration

## Usage

### Hooking command-line arguments

To hook command-line arguments, simply import package in your code:

```golang
import (
    _ "github.com/lylemi/passerby"
)
```

The package will then read the ARG environment variable, split it by spaces, and append it to os.Args.

### Deleting a file after a duration

To delete a config file after a specified duration, call DeleteFileAfterDuration(filePath string):

```golang
import "github.com/lylemi/passerby"

func main() {
    passerby.DeleteFileAfterDuration("path/to/your/file.txt")
    // Your code here
}
```

By default, the file will be deleted after 60 seconds. To change the duration, set the DELETE_DURATION environment variable to a valid Go duration string (e.g., "30s" for 30 seconds, "1m" for 1 minute, etc.).

## Example

see example dir

```bash
export ARG='-name real_passed_name'
./cmdline --fake-arg fake_passed
```

```bash
export ARG='-file config.json'
export DELETE_DURATION=0
./configfile --fake-arg fake_passed
```

## Contribution

If you have any ideas or suggestions, please feel free to submit a pull request. We appreciate any contributions.

## Contact

If you have any questions or suggestions, please feel free to contact us.
