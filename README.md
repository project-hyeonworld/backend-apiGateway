# Gongate

## 🧑‍💻: Intro
> 실시간 프록시 관리 및 서비스 통합을 위한 Golang과 Nginx 기반 동적 API-Gateway

❓ Problem : 실시간 프록시 추가 / 제거 기능 부재 😮

❗ Idea : HTTP 엔드포인트로 Nginx 제어 🤔

💯 Solution : 동적 프록시 관리 API 😁

</br>

## 🖥️: API Documentation
https://hyeonwoody.com/swagger/?urls.primaryName=gongate

</br>

## 🧱: Structure
```
cmd
│   ├── api
│   │   ├── nginx
│   │   │   └── cmd
│   │   │       ├── api
│   │   │       │   ├── nginxController.go
│   │   │       │   ├── nginxControllerInterface.go
│   │   │       │   ├── nginxHandler.go
│   │   │       │   ├── nginxHandlerInterface.go
│   │   │       │   ├── nginxService.go
│   │   │       │   ├── nginxServiceInterface.go
│   │   │       │   └── OsBusiness.go
│   │   │       ├── configuration
│   │   │       │   ├── dependencyInjection
│   │   │       │   │   └── di.go
│   │   │       │   └── secret
│   │   │       │       └── value.go
│   │   │       ├── go.mod
│   │   │       ├── go.sum
│   │   │       ├── lib
│   │   │       │   └── nginxLib.go
│   │   │       └── main.go
│   │   ├── shared
│   │   │   ├── common
│   │   │   │   ├── go.mod
│   │   │   │   ├── go.sum
│   │   │   │   ├── IController.go
│   │   │   │   ├── IHandler.go
│   │   │   │   ├── IRepository.go
│   │   │   │   ├── IService.go
│   │   │   │   └── model
│   │   │   │       ├── NginxConfig.go
│   │   │   │       └── proxyServer.go
│   │   │   └── proto
│   │   │       └── nginx
│   │   │           ├── nginx_grpc.pb.go
│   │   │           ├── nginx.pb.go
│   │   │           └── nginx.proto
│   │   └── site
│   │       └── cmd
│   │           ├── api
│   │           │   ├── nginxConfigBusiness.go
│   │           │   ├── siteController.go
│   │           │   ├── siteControllerInterface.go
│   │           │   ├── siteHandler.go
│   │           │   ├── siteHandlerInterface.go
│   │           │   ├── siteRepository.go
│   │           │   ├── siteRepositoryInterface.go
│   │           │   ├── siteService.go
│   │           │   └── siteServiceInterface.go
│   │           ├── configuration
│   │           │   ├── dependencyInjection
│   │           │   │   └── di.go
│   │           │   └── secret
│   │           │       └── value.go
│   │           ├── go.mod
│   │           ├── go.sum
│   │           ├── lib
│   │           │   └── siteLib.go
│   │           └── main.go
│   ├── configuration
│   │   └── secret
│   │       └── value.go
│   ├── go.mod
│   ├── go.sum
│   ├── go.work
│   ├── go.work.sum
│   ├── hyeonworld_session-service
│   └── main.go
├── Dockerfile
└── nginx
    ├── nginx.conf
    ├── sites-available
    │   ├── config files for aplication
    └── sites-enabled
        ├── symbolic links to sites-available
```
<br>

## ✅: Implementation
1. nginx 프록시 conf 파일 인코딩 후 프록시 서버 추가
```golang
func (s *Service) Add(proxyServer *model.ProxyServer) error {
	content, _ := s.biz.ReadFile(&proxyServer.ApplicationName)
	if content == "" {
		_, err := s.biz.CreateFile(proxyServer)
		if err != nil {
			return err
		}
		s.biz.CreateSymlink(&proxyServer.ApplicationName)
		return nil
	}
	config, err := s.biz.ParseNginxConfig(&content)
	if err != nil {
		return err
	}
	s.biz.AddProxyServer(&config, proxyServer)

	patchContent := config.ToString()

	if err := s.biz.PatchFile(&proxyServer.ApplicationName, &patchContent); err != nil {
		return fmt.Errorf("failed to patch file: %v\n", err)
	}

	return nil
}
```

## 📞: Contact
- 이메일: hyeonwoody@gmail.com
- 블로그: https://velog.io/@hyeonwoody
- 깃헙: https://github.com/hyeonwoody

</br>

## 🛠️: Technologies Used
> Go 1.23.5

> Nginx 

</br>
