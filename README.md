# fs-semgrep-rules
At Fullstory, we leverage Semgrep as a core tool in our security engineering efforts to detect potential issues in our codebase. This involves not only optimizing existing rules but also developing new ones to identify code patterns that could lead to security vulnerabilities.

While many of the rules we create are tailored to our internal codebase, we also develop rules that are broadly applicable to a wide range of projects. The rules shared in this repository are designed to address common code patterns and potential vulnerabilities that are relevant to many codebases.

We are continually refining these rules and adding new ones to improve their effectiveness in finding code bugs that could result in security flaws.

_Note:_ The setup of this repository was in part inspired by other semgrep repos which we have contributed in the past, including [Semgrep's own repo of rules](https://github.com/semgrep/semgrep-rules) as well as [Trail of Bits' Semgrep repo](https://github.com/trailofbits/semgrep-rules).

### Testing

You can run tests locally with:

```bash
semgrep --test --test-ignore-todo --metrics=off
```

To test a specific file:

```bash
semgrep --test --test-ignore-todo --metrics=off --config ./go/iterate-over-empty-map.yaml ./go/iterate-over-empty-map.go
```

## Rules

### go

| ID | Impact | Confidence | Description |
| -- | :----: | :--------: | ----------- |
| [creds-from-jwtconfig](go/creds-from-jwtconfig.yaml) | 🟧 | 🌘 | Using JWT configuration from JSON rather than using service accounts could lead to exposed credentials in code and other insecure key management practices |
| [defer-in-loop](go/defer-in-loop.yaml) | 🟩 | 🌗 | Resource leak due improper use of `defer` |
| [gcs-path-traversal](go/gcs-path-traversal.yaml) | 🟧 | 🌗 | An HTTP redirect was found to be crafted from user-input leading to an open redirect vulnerability |
| [insecure-dir-creation](go/insecure-dir-creation.yaml) | 🟧 | 🌘 | Insecure handling of file and directory writes |
| [missing-close-on-file](go/missing-close-on-file.yaml) | 🟩 | 🌗 | Handling of open file descriptors |
| [missing-defer-http](go/missing-defer-http.yaml) | 🟩 | 🌗 | Handling of HTTP response bodies |


### optimizations

| ID | Impact | Confidence | Description |
| -- | :----: | :--------: | ----------- |
| [math-random-used](optimizations/math-random-used.yaml) | 🟧 | 🌗 | Finds likely cases where `math/rand` may be used insecurely. For the optimization, we exclude functions like `Shuffle` which are rarely used cryptographically |
