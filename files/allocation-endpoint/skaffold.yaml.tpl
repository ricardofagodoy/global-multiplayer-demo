apiVersion: skaffold/v2beta29
kind: Config
deploy:
  kustomize:
    paths:
      - "./${cluster_name}"
    buildArgs: ["--enable-helm"]
    flags:
      apply: ['--server-side'] # Avoid the "Too long: must have at most 262144 bytes" problem