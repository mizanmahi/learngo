# Tasks for module - 4


## Task 1 — Animal Sound Maker 🐾

**Concepts:** struct, interface, receiver function, slice, range, if/else

Create a `Animal` struct with fields `Name` and `Sound`.
Define a `Speaker` interface with a `Speak()` method.
Implement `Speak()` as a receiver function that prints something like:

```
Dog says: Woof!
```

Make a slice of different animals, loop through them with `range`, and call `Speak()` on each.

**Bonus:** Add an `if/else` to print a special message if the animal's name is `"Cat"`.

---

## Task 2 — Simple Shopping Cart 🛒

**Concepts:** struct, receiver function, slice, range, functions, variables

Create a `Product` struct with `Name`, `Price`, and `Quantity`.
Write a receiver function `Total()` that returns `Price * Quantity`.

Create a `Cart` struct that holds a slice of `Product`.
Write a receiver function `GrandTotal()` that loops through all products and sums up their totals.

Print a receipt that looks something like:

```
Apple      x2   $3.00
Bread      x1   $2.50
---------------------
Total:          $5.50
```

---

## Task 3 — Shape Area Calculator 📐

**Concepts:** struct, interface, receiver function, pointers, slice, range

Define a `Shape` interface with an `Area() float64` method.

Create two structs:
- `Circle` with field `Radius`
- `Rectangle` with fields `Width` and `Height`

Implement `Area()` for both. Store them in a `[]Shape` slice and print each shape's area using `range`.

**Bonus:** Use a pointer receiver `*Circle` and see how it differs from a value receiver.

---

## Task 4 — Bank Account 🏦

**Concepts:** struct, pointer, receiver function, if/else, variables

Create a `BankAccount` struct with `Owner` and `Balance`.
Write three **pointer receiver** functions:

- `Deposit(amount float64)`
- `Withdraw(amount float64)`
- `Display()`

In `Withdraw`, use `if/else` to prevent the balance from going below zero and print a warning message.

> 💡 **Why this matters:** Try it first *without* pointer receivers and notice that the balance doesn't actually change. That's the "aha" moment for understanding pointers in Go!

---

## Task 5 — Student Report Card 📋

**Concepts:** struct, functions, slice, range, if/else, variables

Create a `Student` struct with `Name` and a `Grades []int` slice.

Write a receiver function `Average() float64` that calculates the average grade using `range`.

Write another receiver function `LetterGrade() string` that returns a grade letter based on the average:

| Average     | Grade |
|-------------|-------|
| 90 and above | A    |
| 80 – 89      | B    |
| 70 – 79      | C    |
| Below 70     | F    |

Print a report card for a few students.

---

## Suggested Order 💡

| Order | Task | Why |
|-------|------|-----|
| 1st | Task 4 — Bank Account | Best intro to why pointers matter |
| 2nd | Task 3 — Shape Calculator | Great for understanding interfaces |
| 3rd | Task 1 — Animal Sound Maker | Reinforces interfaces with fun context |
| 4th | Task 5 — Report Card | Practices slice + receiver functions |
| 5th | Task 2 — Shopping Cart | Ties everything together |