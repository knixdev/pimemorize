# Pi Memorize

-----

Pi Memorize is a simple CLI tool to check how well you know Pi. 
It will by default prompt you once for all decimal places starting from the first.

- **-limit:** (int)
Limit the digit upper bound for random digits (default 100001)
- **-loop:** (bool)
Keep trying forever until a control signal
- **-random:** (bool)
Pick a random digit to test your PI memory
- **-start** (int)
Starting digit for prompt (default 1)
- **-tipme**: (bool)
Display a tip of starting digits
  
Examples

Run as normal from starting digit:
```bash
knixdev@divmod:~/projects/pimemorize$ go build -o pimemorize && ./pimemorize
PI Memory Checker version 1.0
Enter your digits: 14159
Correct! - 14159
```

Get a tip from your starting digit:
```bash
knixdev@divmod:~/projects/pimemorize$  ./pimemorize -tipme
PI Memory Checker version 1.0
Digits 1 - 5 is: 14159
Enter your digits:
```

Adjust starting digit:
```bash
knixdev@divmod:~/projects/pimemorize$ ./pimemorize -tipme -start=51
PI Memory Checker version 1.0
Digits 51 - 55 is: 58209
Enter your digits:
```

Loop forever:
```bash
knixdev@divmod:~/projects/pimemorize$ ./pimemorize -tipme -start=51 -loop
PI Memory Checker version 1.0
Digits 51 - 55 is: 58209
Enter your digits: 58210
Oops! It's actually: 58209, you entered: 58210
Enter your digits:
```

Pick a random digit of Pi from 0 - full length:
```bash
knixdev@divmod:~/projects/pimemorize$ ./pimemorize -random
PI Memory Checker version 1.0
Enter digit number [5482]: 7
Oops its actually: 0
```

Pick a random digit of Pi to guess up to a limit of decimal places:
```bash
knixdev@divmod:~/projects/pimemorize$ ./pimemorize -random -limit=100
PI Memory Checker version 1.0
Enter digit number [92]: 6
Oops its actually: 4
```