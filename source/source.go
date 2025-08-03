package source

import (
    "sync"
    "os"
    "context"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Source interface {
	Get(name string) (string, error)
}

type EnvSource struct {}

func (s EnvSource) Get(name string) (string, error) {
    return os.Getenv(name), nil
}

var Sources = make(map[string]Source)

// There is no reason RegisterSource should be called from multiple
// routines but it felt wrong to leave it un-safe.
var mu sync.Mutex
func RegisterSource(n string, s Source) {
    mu.Lock()
    Sources[n] = s
    mu.Unlock()
}

type AwsSecretSource struct {
    Region string
}

func (s AwsSecretSource) Get(name string) (string, error) {
    config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(s.Region))
    if err != nil {
        return "" , err
    }

    svc := secretsmanager.NewFromConfig(config)
    input := &secretsmanager.GetSecretValueInput{
        SecretId:     aws.String(name),
        VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
    }

    result, err := svc.GetSecretValue(context.TODO(), input)
    if err != nil {
    // For a list of exceptions thrown, see
    // https://<<{{DocsDomain}}>>/secretsmanager/latest/apireference/API_GetSecretValue.html
        return "", err
    }

    // Decrypts secret using the associated KMS key.
    var secretString string = *result.SecretString

    return secretString, nil
}


