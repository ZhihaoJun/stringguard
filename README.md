# stringguard
string content format validation  
can be used for form or function input validation

## usage
``` golang
guard := stringguard.New()
name := guard.MaxLen("name", c.FormValue("name"), 10)
offsetInt := guard.Int("offset", c.FormValue("offset"))
err := guard.Err()
if err != nil {
  handleError(err)
}
```

## how to extend
