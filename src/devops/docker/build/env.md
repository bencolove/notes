# Environment Variables
* ARG (visible only in `build`)
* ENV (visible both in `build` and `afterwards`)
# `.env`
# combine them

## Capture Envar
Both `VAR` or `:VAR` will be passed on `as-is`
`"$VAR"` in `Dockerfile` will capture the value in current envrionment(build process is essentially a running container which will then be freezed)

Refer to the var by `${Name:-value_if_absent}` or
`${Name:+value_if_present_otherwise_empty}`

## Refer
```Dockerfile
ARG DyArg
ENV Name=${DyArg:-DEFAULT_VALUE}

RUN echo "value is ${Name}"
```

`docker build --progress=plain --build-arg Name=value .`