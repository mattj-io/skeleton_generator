package main

import (
    "os"
    "log"
    "flag"
    "text/template"
)

func main()  {

    type Metadata struct {
        Operator    string
        Email       string
        Name        string
        Appver      string
        Kubever     string
        URL         string
    }

    operatorPtr := flag.String("operator", "myfirstoperator", "name of the operator we are creating")
    emailPtr := flag.String("email", "matt@example.com", "maintainer email address")
    namePtr := flag.String("name", "Matt", "maintainer name")
    appverPtr := flag.String("appver", "0.1.0", "version of app")
    kubeverPtr := flag.String("kubever", "1.15.0", "version of Kubernetes")
    urlPtr := flag.String("url", "https://kudo.dev", "application URL")
    
    flag.Parse()

    if err := ensureDir(*operatorPtr); err != nil {
        log.Fatalln("Directory tree creation failed with error: " + err.Error())
    }

    tpl, err := template.New("operator_yaml").Parse(operator_yaml)
    if err != nil {
        log.Fatalln("Could not create Template: " + err.Error())
    }

    metadata := Metadata{
        Operator: *operatorPtr,
        Email: *emailPtr,
        Name: *namePtr,
        Appver: *appverPtr,
        Kubever: *kubeverPtr,
        URL: *urlPtr,
    }

    cwd, err := os.Getwd()
    if err != nil {
        log.Fatalln("Could not get current directory: " + err.Error())
    }
    
    opyaml_path := cwd + "/" + *operatorPtr + "/operator.yaml"
    f := createFile(opyaml_path)

    if err = tpl.Execute(f, metadata); err != nil {
        log.Fatalln("Could not execute template" + err.Error())
    }
    err = f.Close()
    if err != nil {
        log.Fatalln("Cloud not close operator.yaml: " + err.Error())
    } 

    pyaml_path := cwd + "/" + *operatorPtr + "/params.yaml"
    f = createFile(pyaml_path)

    _, err = f.WriteString(params_yaml)
    if err != nil {
        f.Close()
        log.Fatalln("Could not write params.yaml: " + err.Error())
    }
    err = f.Close()
    if err != nil {
        log.Fatalln("Cloud not close params.yaml: " + err.Error())
    }   
      
}

func createFile(path string) (f *os.File) {
    f, err := os.Create(path)
    if err != nil {
        log.Fatalln(err)
    }
    return f
}

func ensureDir(dirName string) error {

    cwd, err := os.Getwd()
    if err != nil {
        return err
    }
    path := cwd + "/" + dirName + "/templates"
    
    err = os.MkdirAll(path, 0777)

    if err == nil || os.IsExist(err) {
        return nil
    } else {
        return err
    }
}
