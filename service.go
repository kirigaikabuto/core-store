package core_store

import (
	"bytes"
	"encoding/json"
	"fmt"
	movie_store "github.com/kirigaikabuto/movie-store"
	users_store "github.com/kirigaikabuto/users-store"
	"io/ioutil"
	"net/http"
)

type CoreService interface {
	ListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie, error)
	GetMovieByName(cmd *GetMovieByNameCommand) (*movie_store.Movie, error)
	GetMovieById(cmd *GetMovieByIdCommand) ([]movie_store.Movie, error)
	SignUpUsingEmail(cmd *CreateUserCommand) (*users_store.User, error)
}

type coreService struct {
	amqpRequests AmqpRequests
}

func NewCoreService(amqpReq AmqpRequests) CoreService {
	return &coreService{
		amqpReq,
	}
}

func (svc *coreService) ListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie, error) {
	movies, err := svc.amqpRequests.GetListMovies(cmd)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (svc *coreService) GetMovieByName(cmd *GetMovieByNameCommand) (*movie_store.Movie, error) {
	movie, err := svc.amqpRequests.GetMovieByName(cmd)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (svc *coreService) GetMovieById(cmd *GetMovieByIdCommand) ([]movie_store.Movie, error) {
	movie, err := svc.amqpRequests.GetMovieById(cmd)
	if err != nil {
		return nil, err
	}
	requestBody, err := json.Marshal(map[string]string{
		"title": movie.Name,
	})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post("http://127.0.0.1:5000/api/v1/recommend/", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	output := &MovieRecommendResponse{}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, err
	}
	var movies []movie_store.Movie
	for _, v := range output.Movies {
		newMovie, err := svc.amqpRequests.GetMovieByName(&GetMovieByNameCommand{
			v.Name,
		})
		if err != nil {
			return nil, err
		}
		newMovie.Score = v.Score
		movies = append(movies, *newMovie)
	}
	return movies, nil
}

func (svc *coreService) SignUpUsingEmail(cmd *CreateUserCommand) (*users_store.User, error) {
	newUser, err := svc.amqpRequests.CreateUser(cmd)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
