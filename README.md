# gh-pages

This branch is the Helm chart repository for
[signalilo](https://github.com/saremox/signalilo), published via GitHub
Pages and maintained automatically by
[helm/chart-releaser-action](https://github.com/helm/chart-releaser-action)
(see `.github/workflows/helm.yml` on `main`).

Don't edit this branch by hand — packaged chart releases (`.tgz`) and the
generated `index.yaml` are pushed here by that workflow whenever a tag is
pushed. The chart source lives on `main` under `charts/signalilo`.

Once GitHub Pages is enabled for this branch (Settings -> Pages -> Source:
Deploy from a branch -> `gh-pages` / `/(root)`), the chart repository will
be reachable at:

```
helm repo add signalilo https://saremox.github.io/signalilo/
```
