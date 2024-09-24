# Roll Some Dice

`roll-some` is a super simple cli and a htmx web interface for rolling
several rows of different types of die.

```
row = expr | expr, row | expr, operator, row ;
expr = modifier, action | action | int ;
action = "d", int | "(", row, ")" ;
modifier = low | high | int ;
low = "l", int | int, "l", int ;
high = "h", int | int, "h", int ;
operator = "+", | "-" ;
```

## Future Plans
If I feel like it I might update this at a later point. The things I would like to improve are (in no particular order):

* Update general code structure
* Improve CLI API and add documentation
* Improve HTMX powered web interface
