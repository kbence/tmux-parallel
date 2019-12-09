# tmux-parallel

`tmux-parallel` aims to be a GNU parallel compatible alternative to build and run commands in a batch while also observing the live output using [tmux](https://github.com/tmux/tmux).

## Installation

### Compiling from source

You'll need to have at least Go 1.13. Execute the following command:

    go get -u github.com/kbence/tmux-parallel

This should put the binary into your `$GOPATH/bin` directory. Alternatively, you can clone the repository or download it as a Zip and extract. In that case, issue the following command in the project's root directory:

    go get .

## Usage

The basic format of a command looks something like this:

    tmux-parallel [options] command template [::: arguments | :::+ arguments] ...

If no parameter references are present in the command, the expanded argument list will get appended to the base command, for example:

    tmux-parallel echo args ::: 1 2 ::: 3 4

will execute the following commands in parallel:

    echo 1 3
    echo 1 4
    echo 2 3
    echo 2 4

If references are present, `tmux-parallel` will only replace markers with their matching arguments in order, columns without a reference will not be expanded at all.

    tmux-parallel echo {2} ::: 1 ::: 2

will result in calling

    echo 2

`tmux-parallel` does not support reading from stdin at the moment.

## Argument blocks

Currently, the following argument expansions are supported:

### Cartesian join (:::)

This type of join generates all the possible permutations of the items on the left side and the block following until the next operator. Example:

    tmux-parallel echo ::: a b ::: c d

will result in argument pairs `a b`, `a d`, `b c` and `b d`.

### 1:1 join (:::+)

This operator marks arguments that needs to be joined 1:1. Since the number of arguments can differ, the smallest amount will be used for expansion. Example:

    tmux-parallel echo ::: 1 2 3 :::+ 4 5

will only use argument pairs `1 4` and `2 5`, `3` will be ignored completely. When a chain of `:::+` connected arguments get expanded with a Cartesian operator (`:::`), the already expanded list will be used. Example:

    tmux-parallel echo ::: a b ::: 1 2 3 :::+ 4 5

will result in arguments `a 1 4`, `a 2 5`, `b 1 4`, `b 2 5`.
