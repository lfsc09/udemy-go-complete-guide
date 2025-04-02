# Course execises

#### Tips

- For monorepos it is ok to name modules with "invalid urls", where the url will not point exacly to where the `go.mod` will be.

</br>

### Section 5

This section focus on the use of `struct`, how to declare and use it.

##### Problem

Read two string values (One multi-words) from STDIN, show on the STDOUT and save it in a file.

#### Tips

- Trail `\n` and `\r` since Windows may have its line break with `\n\r`.
- `json.Marshal` requires that struct fields are public.

</br>

### Section 6

This section focus on the use of `interface`, showing how to use an interface as abstract parameter in a function, embedding interfaces one inside the other.

##### Problem

Now we have Notes and Todos which are similar structs. The program must handle them with abstraction to reduce code.
