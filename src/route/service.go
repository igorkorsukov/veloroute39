package route

import (
	"infra"
)

type RouteService interface {
	Route(uid string, r *Route) error
	Routes(args string, rs *[]Route) error
	Create(r Route, uid *string) error
	Update(r Route, args *string) error
	Remove(uid string, args *string) error
}

type routeService struct {
	Repo RouteRepository
}

func NewRouteService(repo RouteRepository) RouteService {
	if repo == nil {
		repo = NewRepository(nil)
	}
	s := &routeService{Repo: repo}
	return s
}

func (s *routeService) Route(uid string, r *Route) error {
	rt, err := s.Repo.Find(uid)
	*r = rt
	return err
}

func (s *routeService) Routes(args string, rs *[]Route) error {
	l, err := s.Repo.List()
	*rs = l
	return err
}

func (s *routeService) Create(r Route, uid *string) error {

	r.UID = infra.NewUUID()
	if r.Color == "" {
		r.Color = s.defaultColor()
	}
	err := s.Repo.Add(r)
	if err != nil {
		return err
	}
	*uid = r.UID
	return nil
}

func (s *routeService) Update(r Route, args *string) error {
	err := s.Repo.Delete(r.UID)
	if err != nil {
		return err
	}

	err = s.Repo.Add(r)
	if err != nil {
		return err
	}
	return nil
}

func (s *routeService) Remove(uid string, args *string) error {
	err := s.Repo.Delete(uid)
	if err != nil {
		return err
	}
	return nil
}

func (s *routeService) defaultColor() string {
	return "#cc0000"
}