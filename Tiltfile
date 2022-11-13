# -*- mode: Python -*-
load('ext://namespace', 'namespace_create')
load('ext://helm_remote', 'helm_remote')
load('ext://helm_resource', 'helm_resource')
namespace_create('argocd')
namespace_create('production')
namespace_create('staging')
namespace_create('dev')

helm_remote('argo-cd',
    release_name='argo-cd',
    repo_name='argo-cd',
    repo_url='https://argoproj.github.io/argo-helm',
    values=['./infrastructure/argo-cd/values.yaml'],
    namespace='argocd',
)

k8s_resource(
  'argo-cd-argocd-server',
  port_forwards=['8080:8080'],
)

k8s_yaml("{}/{}.yaml".format('infrastructure/app', "greeter"))
k8s_resource(
  objects=['greeter:applicationset'],
  new_name="greeter-application-set"
)

# #Dev Deployment
# docker_build('greeter',
#   context='.',
#   dockerfile='./app/docker/Dockerfile',
#   entrypoint='/main',
#   only=[
#       './app/cmd/main.go',
#       './app/go.mod',
#       './app/go.sum'
#   ],
#   live_update=[
#     sync('./app/go.mod', '/service/go.mod'),
#     sync('./app/go.sum', '/service/go.sum'),
#     sync('./app/cmd', '/service/cmd'),
#   ]
# )

# helm_resource(
#   name="greeter",
#   namespace='dev',
#   chart="./chart",
#   image_deps= ["greeter"],
#   image_keys= [('deployment.image', 'deployment.tag')],
#   port_forwards=["5555:8000"]
# )