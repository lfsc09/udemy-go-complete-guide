# Course execises

#### Tips

- For monorepos it is ok to name modules with "invalid urls", where the url will not point exacly to where the `go.mod` will be.

### Section 5

Read two string values (One multi-words) from STDIN, show on the STDOUT and save it in a file.

#### Tips

- Trail `\n` and `\r` since Windows may have its line break with `\n\r`.
- `json.Marshal` requires that struct fields are public.
