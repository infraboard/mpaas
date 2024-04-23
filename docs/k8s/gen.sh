#!/bin/bash
#创建一个k8s用户，并赋予defaults空间部分资源的只读服务
UserName=mcloud
ApiServerEndpoints=https://134.175.99.215:6443
ClusterName=k3s
NS=default
mkdir -p /etc/kubernetes/pki/client/${UserName}
cd /etc/kubernetes/pki/client/${UserName}
#创建用户证书
openssl genrsa -out ${UserName}.key 2048
openssl req -new -key ${UserName}.key -out ${UserName}.csr -subj "/CN=${UserName}"
openssl x509 -req -in ${UserName}.csr -CA /var/lib/rancher/k3s/server/tls/server-ca.crt \
-CAkey /var/lib/rancher/k3s/server/tls/server-ca.key -CAcreateserial -out ${UserName}.crt -days 3650
#查看证书有效期限
#openssl x509 -noout -text -in ${UserName}.crt

#创建user 访问Kubernetes config file
#设置一个集群名称并倒入证书
/usr/local/bin/k3s kubectl config set-cluster ${ClusterName} \
  --server=${ApiServerEndpoints} \
  --certificate-authority=/etc/kubernetes/pki/ca.crt \
  --embed-certs=true \
  --kubeconfig=./${UserName}.config

# 将客户的证书导入配置文件
/usr/local/bin/k3s kubectl config set-credentials ${UserName} \
  --client-certificate=${UserName}.crt \
  --client-key=${UserName}.key \
  --embed-certs=true \
  --kubeconfig=./${UserName}.config

#nsmaspace 设置用户默认访问的ns
#设置上下文，把集群和用户导入到一起
/usr/local/bin/k3s kubectl config set-context ${UserName}@${ClusterName} \
  --cluster ${ClusterName} \
  --user=${UserName} \
  --namespace=${NS} \
  --kubeconfig=./${UserName}.config
#将用户绑定到上下文上
/usr/local/bin/k3s kubectl config use-context ${UserName}@${ClusterName} \
  --kubeconfig=./${UserName}.config

#创建role角色设置权限
/usr/local/bin/k3s kubectl apply -f - <<EOF
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: ${UserName}-role-bind
  namespace: ${NS}
subjects:
- kind: User
  name: ${UserName} 
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role        
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
EOF

#创建用户
useradd -m ${UserName}
echo "${UserName}:Zs1gmm!" | chpasswd
mkdir -p /home/${UserName}/.kube/
cp $PWD/${UserName}.config /home/${UserName}/.kube/config
chown ${UserName}.${UserName} /home/${UserName}/.kube/config
chmod 600 /home/${UserName}/.kube/config

echo "kubernetes The configuration file location is $PWD/${UserName}.config"
echo "test command: KUBECONFIG=$PWD/${UserName}.config /usr/local/bin/k3s kubectl get pods"