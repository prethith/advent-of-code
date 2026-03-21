# My Advent of Code (AoC) Attempts

This repo contains the code (in various languages) for my attempts at the [Advent of Code](https://adventofcode.com) competition.

Repo structure:
```
.
├── 2025
│   ├── go
│   ├── inputs (empty in this repo)
│   └── rust
├── 2024 (not added yet)
│   ├── go
│   ├── inputs (empty in this repo)
│   └── rust
└── README.md
```

Adhering to the competition rules that forbid sharing the input, the `inputs/` folder is empty on this GitHub repo. You can access it directly for each question on the official website. 

The repo is organized by year, and then by language. 

## How to run this locally?

1. Clone this repo using:
```
git clone https://github.com/prethith/advent-of-code.git
```
2. Store the input for each day's puzzle in `inputs/dayxx.txt`, where `xx` ranges from `01` to whenever that year's competition ended. For example, in 2025, the day ranges are from `01` to `12`.
3. For each day, you can store the test input to each question in `inputs/dayxx_test.txt`.
4. Language-specific instructions:
    - Go:
        From the top `go/` folder, run:

        ```bash
        go run main.go 1 # or any other day like 2,3,4,.....
        go run main.go 1 --test # to use the test input
        ```
