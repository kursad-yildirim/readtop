# tuff-toolkit
## read-top
This utility reads dumped "top command output[1]" data and displays it as a regular top command.

User can adjust refresh rates and also a token to search for. Normal speed will be used as the default refresh rate. When the provided token (E.g. a process name to track) is captured the output will switch to the slow speed.

Default values are:

|**OPTION**           |**VALUE**   |
|:---                 |:---:       |
|**Normal Speed**     |0.1 seconds |
|**Slow Speed**       |1 seconds   |
|**Token to Look**    |""          |
|**Lines to Display** |50          |


Usage:

```shell
$ ./read-top -h


read-top:
        Displays linux 'top' command output, which is dumped to a text file, as a self refreshing regular top command.

Usage: read-top [OPTION] FILENAME

Options
        -r/--refresh     Normal Speed in seconds (Default 0.1)
        -s/--slow        Slow Speed in seconds (Default 1)
        -t/--token       Token to Look (If the token is captured refresh rate will switch to 'slow')
        -l/--lines       Number of lines to display (Default 50)
        -h/--help        Display help


Examples:
        read-top /tmp/top.out
        read-top -s 1 -l 25 /tmp/top.out
        read-top -s 0.2 -S 1 -t my_process /tmp/top.out
```

[1] Top command output can be generated by using the command `top -d 3 -n 100 -b > /tmp/top.out &` for getting 100 command runs separated by 3 seconds.
