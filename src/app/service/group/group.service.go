package group

import (
	"context"
	"github.com/isd-sgcu/rnkm65-gateway/src/app/dto"
	"github.com/isd-sgcu/rnkm65-gateway/src/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type Service struct {
	client proto.GroupServiceClient
}

func NewService(client proto.GroupServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) FindByToken(token string) (result *proto.Group, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, errRes := s.client.FindByToken(ctx, &proto.FindByTokenGroupRequest{Token: token})
	if errRes != nil {
		st, ok := status.FromError(errRes)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return nil, &dto.ResponseErr{
					StatusCode: http.StatusNotFound,
					Message:    st.Message(),
					Data:       nil,
				}
			default:

				log.Error().
					Err(errRes).
					Str("service", "group").
					Str("module", "findByToken").
					Msg("Error while connecting to service")

				return nil, &dto.ResponseErr{
					StatusCode: http.StatusServiceUnavailable,
					Message:    "Service is down",
					Data:       nil,
				}
			}
		}
		log.Error().
			Err(errRes).
			Str("service", "group").
			Str("module", "findByToken").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Group, nil
}

func (s *Service) Create(in *dto.GroupDto) (result *proto.Group, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpDto := &proto.Group{
		LeaderID: in.LeaderID,
		Token:    in.Token,
	}

	res, errRes := s.client.Create(ctx, &proto.CreateGroupRequest{Group: grpDto})
	if errRes != nil {

		log.Error().
			Err(errRes).
			Str("service", "group").
			Str("module", "create").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Group, nil
}

func (s *Service) Update(id string, in *dto.GroupDto) (result *proto.Group, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	grpDto := &proto.Group{
		Id:       id,
		LeaderID: in.LeaderID,
		Token:    in.Token,
	}

	res, errRes := s.client.Update(ctx, &proto.UpdateGroupRequest{Group: grpDto})
	if errRes != nil {
		st, ok := status.FromError(errRes)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return nil, &dto.ResponseErr{
					StatusCode: http.StatusNotFound,
					Message:    st.Message(),
					Data:       nil,
				}
			default:

				log.Error().
					Err(errRes).
					Str("service", "group").
					Str("module", "update").
					Msg("Error while connecting to service")

				return nil, &dto.ResponseErr{
					StatusCode: http.StatusServiceUnavailable,
					Message:    "Service is down",
					Data:       nil,
				}
			}
		}

		log.Error().
			Err(errRes).
			Str("service", "group").
			Str("module", "update").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Group, nil

}

func (s *Service) Delete(id string) (result bool, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, errRes := s.client.Delete(ctx, &proto.DeleteGroupRequest{Id: id})
	if errRes != nil {
		st, ok := status.FromError(errRes)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return false, &dto.ResponseErr{
					StatusCode: http.StatusNotFound,
					Message:    st.Message(),
					Data:       nil,
				}
			default:

				log.Error().
					Err(errRes).
					Str("service", "group").
					Str("module", "delete").
					Msg("Error while connecting to service")

				return false, &dto.ResponseErr{
					StatusCode: http.StatusServiceUnavailable,
					Message:    "Service is down",
					Data:       nil,
				}
			}
		}

		log.Error().
			Err(errRes).
			Str("service", "group").
			Str("module", "delete").
			Msg("Error while connecting to service")

		return false, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Success, nil

}
