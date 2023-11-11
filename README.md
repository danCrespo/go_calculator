# go-calculator

## What is go-calculator?

go-calculator is a _console application_ that aims to make using mathematics faster and easier within a _terminal/shell_.

Although it has always been possible to solve mathematical operations in a terminal, not all shell environments support the use of floating point numbers themselves, for example the _bash shell_. To use this type of numbers you need to use other commands such as bc, sed, awk, among others (which you may need to install), or use pipes to pass the result to another shell such as _zsh_ or _ksh_.

Although this does not imply being a disadvantage, there may be some situation in which you need to solve operations from other areas of mathematics such as algebra, geometry, trigonometry, etc. For this you would have a more complex situation to solve _-obviously if you need to do it using a terminal, for example if writing a console program or shell script-_. For those types of situations or common situations, that is what this program was intended for.

go-calculator can solve **arithmetic, algebraic, geometric, trigonometric, statistical, date-time operations and measurement unit conversions**.

## How it works? ⚙️

Because go-calculator is capable of solving operations in some areas of mathematics, there are several options to specify which formula or function you want to use:

### Arithmetic operations

This is the default use if no flags are provided.

```bash
# Very basic operations:
go-calculator '5*5'  # Output: Result: 25

# More complex operations, like grouped operations:
go-calculator '5 * 5 + (125.85 * 2 / 6^3 + (325.255 * 3)) / 2'  # Output: Result: 513.4651388888889

# ...Or you can control the number of result decimals:
go-calculator -p 2 '5 * 5 + (125.85 * 2 / 6^3 + (325.255 * 3)) / 2'  # Output: Result: 513.46
```

---

### Geometric operations:

#### **_Area_**

To get the area of a figure, the -a or --area flag must be passed and its value must be the figure name or alias, the parameters would be the perimeters:

```bash
# To get the area of a Square:
go-calculator -a square '5'  # Output: Result: 20

# Or you can pass an operation as perimeter:
go-calculator -a square '5*5'  # Output: Result: 100

# To get the area of a Triangle:
go-calculator --area tri. '5*5' 5  # Output: Result: 62.5

# Or you can use more specific pattern:
go-calculator -a square 'b=5*5' h=5  # Output: Result: 62.5
```

---
