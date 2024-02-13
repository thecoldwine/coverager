# Coverager

No dependencies coverage-to-markdown-table converter. Useful for CI pipelines and some documents.

```console
go test -v -coverprofile=coverage.out ./...
go tool cover -func=coverage.out | coverager --type=gofunc
```
And this will yield something like:

```markdown
| Filename | Function | Coverage |
|----------|----------|----------|
| github.com/thecoldwine/coverager/gofunc.go:9: | ParseGoFunc | ✅ 0.0% |
| github.com/thecoldwine/coverager/main.go:21: | main | ✅ 0.0% |
| github.com/thecoldwine/coverager/reporter.go:28: | RenderReport | ✅ 88.9% |
| github.com/thecoldwine/coverager/util.go:8: | ParsePercentage | ✅ 0.0% |
| Total | | ✅ 13.8% |
```

Common usage then will be to put this app into your CI pipeline, i.e. GHA:

```yaml
    - name: Test and coverage
      run: |
        go test -v -coverprofile=coverage.out ./...
        go tool cover -func=coverage.out | coverager --type=gofunc >> $GITHUB_STEP_SUMMARY
```

Supported coverage report formats:
* golang func

Planned coverage report formats:
* cobertura
* jacoco