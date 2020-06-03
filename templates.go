package main

var operator_yaml = `apiVersion: kudo.dev/v1beta1
name: "{{.Operator}}"
operatorVersion: "0.1.0"
appVersion: "{{.Appver}}
kubernetesVersion: "{{.Kubever}}"
maintainers:
  - name: {{.Name}}
    email: {{.Email}}
url: {{.URL}}
tasks:
  - name: app
    kind: Apply
    spec:
      resources:
        - deployment.yaml
plans:
  deploy:
    strategy: serial
    phases:
      - name: main
        strategy: parallel
        steps:
          - name: everything
            tasks:
              - app
`

var params_yaml = `apiVersion: kudo.dev/v1beta1
parameters:
  - name: OPTIONAL_PARAM
    description: "This parameter is not required."
    required: False
  - name: REQUIRED_PARAM
    description: "This parameter is required but does not have a default value."
    required: True
  - name: ARRAY_PARAM
    description: "This parameter describes an array of values."
    default:
      - user1
      - user2
    type: array
  - name: MAP_PARAM
    description: "This parameter describes a map of values."
    default:
      label1: foo
      label2: bar
    type: map
`


