if [[ -z $1 ]]; then
        echo "USAGE: sh $0 values.yaml"
else
        helm install -f $1 --name mysql-poc-5  .
fi
