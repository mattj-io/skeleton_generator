## Skeleton generator for KUDO operator development

Usage:

  -appver string

    	version of app (default "0.1.0")

  -email string
    
    	maintainer email address (default "matt@example.com")

  -kubever string

    	version of Kubernetes (default "1.15.0")

  -name string

    	maintainer name (default "Matt")

  -operator string

    	name of the operator we are creating (default "myfirstoperator")

  -url string

    	application URL (default "https://kudo.dev")

Currently creates the appropriate directory structure, along with a partially populated operator.yaml and an example params.yaml

```
        mattbook:test matt$ tree
        .
        ├── operator.yaml
        ├── params.yaml
        └── templates

        1 directory, 2 files
```

