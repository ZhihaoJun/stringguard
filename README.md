# stringguard
string content format validation  
can be used for form or function input validation

## usage
* clone repo into `$GOPATH/zhihaojun.com/stringguard`

``` golang
import "zhihaojun.com/stringguard"

guard := stringguard.NewGuard()
name := guard.MaxLen("name", c.FormValue("name"), 10)
offsetInt := guard.Int("offset", c.FormValue("offset"))
err := guard.Err()
if err != nil {
  handleError(err)
}
```

## how to extend
``` golang
func (g *stringguard.Guard) YourRulesName() {
  // first check err
  if g.err != nil {
    return
  }
  // rule check and set g.err properly
  // you can check out for stringguard.go to see some implementation
}
```
