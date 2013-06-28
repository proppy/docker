package docker

import (
	_ "net/rpc"
)

type Service struct {
	*Server
}

func (s *Service) Version(_ *struct{}, version *APIVersion) error {
	*version = s.DockerVersion()
	return nil
}

func (s *Service) ContainerKill(name string, _ *struct{}) error {
	return s.Server.ContainerKill(name)
}

func (s *Service) ImagesSearch(term string, res *[]APISearch) (err error) {
	*res, err = s.Server.ImagesSearch(term)
	return
}

type ImagesRequest struct {
	All bool
	Filter string
}

func (s *Service) Images(req *ImagesRequest, res *[]APIImages ) (err error) {
	*res, err = s.Server.Images(req.All, req.Filter)
	return
}

func (s *Service) DockerInfo(_ *struct{}, res *APIInfo) error {
	*res = *s.Server.DockerInfo()
	return nil
}

func (s *Service) ImageHistory(name string, res *[]APIHistory) (err error) {
	*res, err = s.Server.ImageHistory(name)
	return
}

func (s *Service) ContainerChanges(name string, res *[]Change) (err error) {
	*res, err = s.Server.ContainerChanges(name)
	return
}

type ContainerRequest struct {
	All, Size bool
	N int
	Since, Before string
}

func (s *Service) Containers(req *ContainerRequest, res *[]APIContainers) error {
	*res = s.Server.Containers(req.All, req.Size, req.N, req.Since, req.Before)
	return nil
}

type ContainerCommitRequest struct {
	Name, Repo, Tag, Author, Comment string
	Config Config
}

func (s *Service) ContainerCommit(req *ContainerCommitRequest, res *string) (err error) {
	*res, err = s.Server.ContainerCommit(req.Name, req.Repo, req.Tag, req.Author, req.Comment, &req.Config)
	return
}

type ContainerTagRequest struct {
	Name, Repo, Tag string
	Force bool
}

func (s *Service) ContainerTag(req *ContainerTagRequest, _ *struct{}) error {
	return s.Server.ContainerTag(req.Name, req.Repo, req.Tag, req.Force)
}

func (s *Service) ContainerCreate(req *Config, res *string) (err error) {
	*res, err = s.Server.ContainerCreate(req)
	return
}

type ContainerRestartRequest struct {
	Name string
	T int
}

func (s *Service) ContainerRestart(req *ContainerRestartRequest) error {
	return s.Server.ContainerRestart(req.Name, req.T)
}

type ContainerDestroyRequest struct {
	Name string
	RemoveVolume bool
}

func (s *Service) ContainerDestroy(req *ContainerDestroyRequest, _ *struct{}) error {
	return s.Server.ContainerDestroy(req.Name, req.RemoveVolume)
}

type ImageDeleteRequest struct {
	Name string
	AutoPrune bool
}

func (s *Service) ImageDelete(req *ImageDeleteRequest, res *[]APIRmi) (err error) {
	*res, err = s.Server.ImageDelete(req.Name, req.AutoPrune)
	return
}

type ContainerStartRequest struct {
	Name string
	HostConfig HostConfig
}

func (s *Service) ContainerStart(req *ContainerStartRequest, _ *struct{}) error {
	return s.Server.ContainerStart(req.Name, &req.HostConfig)
}

type ContainerStopRequest struct {
	Name string
	T int
}

func (s *Service) ContainerStop(req *ContainerStopRequest, _ *struct{}) error {
	return s.Server.ContainerStop(req.Name, req.T)
}

func (s *Service) ContainerWait(name string, res *int) (err error) {
	*res, err = s.Server.ContainerWait(name)
	return err
}

type ContainerResizeRequest struct {
	Name string
	H int
	W int
}

func (s *Service) ContainerResize(req *ContainerResizeRequest, _ *struct{}) error {
	return s.Server.ContainerResize(req.Name, req.H, req.W)
}
