# win-process-tree

A mini windows go program that creates process trees of different shapes. It's primary
use is for testing process management tools (like the [buildkite-agent](https://github.com/buildkite/agent/)).

## Usage

Start by compiling. On windows:

    go build -o win-process-tree.exe .

Or on other operating systems:
    
    GOOS=windows go build -o win-process-tree.exe .

Then, run the program in a cmd.exe prompt on windows.

    win-process-tree 3tree

## Modes

There are currently two process tree shapes available.

### Three-deep connected tree

A three-level tree of processes that will slowly shutdown over 30 seconds, starting from the leaf processes.

Command:

    win-process-tree 3tree

Tree shape:

    5460 - win-process-tree.exe
      1584 - win-process-tree.exe
        6012 - win-process-tree.exe
          7008 - win-process-tree.exe
          5580 - win-process-tree.exe
          4884 - win-process-tree.exe
        6012 - win-process-tree.exe
          7008 - win-process-tree.exe
          5580 - win-process-tree.exe
          4884 - win-process-tree.exe
        6012 - win-process-tree.exe
          7008 - win-process-tree.exe
          5580 - win-process-tree.exe
          4884 - win-process-tree.exe
      1584 - win-process-tree.exe
        6012 - win-process-tree.exe
          7008 - win-process-tree.exe
          5580 - win-process-tree.exe
          4884 - win-process-tree.exe
        6012 - win-process-tree.exe
          7008 - win-process-tree.exe
          5580 - win-process-tree.exe
          4884 - win-process-tree.exe
        6012 - win-process-tree.exe
          7008 - win-process-tree.exe
          5580 - win-process-tree.exe
          4884 - win-process-tree.exe
      1584 - win-process-tree.exe
        6012 - win-process-tree.exe
          7008 - win-process-tree.exe
          5580 - win-process-tree.exe
          4884 - win-process-tree.exe
        6012 - win-process-tree.exe
          7008 - win-process-tree.exe
          5580 - win-process-tree.exe
          4884 - win-process-tree.exe
        6012 - win-process-tree.exe
          7008 - win-process-tree.exe
          5580 - win-process-tree.exe
          4884 - win-process-tree.exe

### Orphaned Processes

Nearly 30 orphaned processes that will run for 5 minutes.

Command:

    win-process-tree orphans

Tree shape:

    5460 - win-process-tree.exe # original process with 3 children
      1584 - win-process-tree.exe
      4884 - win-process-tree.exe
      3596 - win-process-tree.exe
    3916 - win-process-tree.exe
    6720 - win-process-tree.exe
    2108 - win-process-tree.exe
    3152 - win-process-tree.exe
    4908 - win-process-tree.exe
    888 - win-process-tree.exe
    6440 - win-process-tree.exe
    1284 - win-process-tree.exe
    6696 - win-process-tree.exe
    3848 - win-process-tree.exe
    4892 - win-process-tree.exe
    6348 - win-process-tree.exe
    5336 - win-process-tree.exe
    5012 - win-process-tree.exe
    6012 - win-process-tree.exe
    7008 - win-process-tree.exe
    5580 - win-process-tree.exe
    3504 - win-process-tree.exe
    6020 - win-process-tree.exe
    2920 - win-process-tree.exe
    5608 - win-process-tree.exe
    6704 - win-process-tree.exe
    4512 - win-process-tree.exe
    1476 - win-process-tree.exe
    5156 - win-process-tree.exe
    1880 - win-process-tree.exe
    2096 - win-process-tree.exe
