Contributing to FullStory Semgrep Rules
=========================

The information below will help you set up a local development environment,
as well as performing common development tasks.

## Requirements

The only development environment requirement *should* be Python 3.7
or newer. Development and testing is actively performed on macOS and Linux,
but Windows and other supported platforms that are supported by Python
should also work.

## Development steps

Then [install semgrep CLI](https://semgrep.dev/docs/getting-started/), and you are good to start development.

### Testing

You can run tests locally with:

```bash
semgrep --test --test-ignore-todo --metrics=off
```

To test a specific file:

```bash
semgrep --test --test-ignore-todo --metrics=off --config ./go/iterate-over-empty-map.yaml ./go/iterate-over-empty-map.go
```

- Test your rule against well known, large codebases. For instance, a good set of projects to test your rules against include:
  - FSTA (this repo, of course)
  - [moby/moby](https://github.com/moby/moby) (Docker)
  - [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes)
  - [etcd-it/ectd](https://github.com/etcd-io/etcd)
  - [gin-gonic](https://github.com/gin-gonic/gin)
- Your rules shuld be reasonably performant and the false positives should be as minimal as possible. 
- Any noise rules should remain in the experimental stage or should be marked as a "linter" rather than a security rule.

### Development practices

Before publishing a new rule, or updating an existing one, make sure to review the checklist below:

- [ ] Add metadata. Semgrep [defines which metadata fields are required](https://semgrep.dev/docs/contributing/contributing-to-semgrep-rules-repository/#writing-a-rule-for-semgrep-registry)
    - [ ] Add a non-standard `metadata.description` field. It will be used as a description in the `semgrep-rules` README table.
    - For `metadata.references` provide a link to official documentation, GitHub issues, or some reputable website. Avoid linking to websites that may disappear in the future.

- [ ] Validate metadata against the official schema
    - Download python validation script `wget https://raw.githubusercontent.com/returntocorp/semgrep-rules/develop/.github/scripts/validate-metadata.py`
    - Download rules schema `wget https://raw.githubusercontent.com/returntocorp/semgrep-rules/develop/metadata-schema.yaml.schm`
    - Run `python ./validate-metadata.py -s ./metadata-schema.yaml.schm -f .`

- [ ] Add tests
    - [ ] At least one true positive (`ruleid: ` comment)
    - [ ] At least one true negative (`ok: ` comment)
    - Tests are allowed to crash when running them directly or to be meaningless
    - However, try writing tests that can be compiled or parsed by the language interpreter
    - The first few test cases should be easy to understand, the later should be more complex or check for edge-cases
    - [ ] Make sure all tests pass, run `semgrep --test --test-ignore-todo --metrics=off`

- [ ] Run official semgrep lints with `semgrep --validate --metrics=off --config ./<new-rule>.yaml`

- [ ] Review style of the rules
    - [ ] Use 2 spaces for indentation
    - [ ] Use `>-` for multiline messages
    - [ ] Use backticks in messages e.g., `$VAR`, `$FUNC`, `some.method()`
    - The `languages` field in `[go, rust]` format are preferable (not `- go \n -java`)

- [ ] Check amount of false-positives on some large public repositories

- [ ] Check performance - take a look at [r2c methodology](https://github.com/returntocorp/semgrep-rules/blob/main/tests/performance/test_public_repos.py)

- [ ] Add the new rules to the README
    - Run `python ./rules_table_generator.py` to re-generate the table
    - Manually check if the table was correctly generated

### Documentation

All information that you need to understand a rule is inside it. Semgrep documentation can be found [here](https://semgrep.dev/docs/).
e, r2c team still works on resolving it.