package core_store

import (
	"encoding/json"
	"errors"
	"github.com/djumanoff/amqp"
	movie_store "github.com/kirigaikabuto/movie-store"
	users_store "github.com/kirigaikabuto/users-store"
)

type AmqpRequests struct {
	clt amqp.Client
}

type ErrorSt struct {
	Text string `json:"text"`
}

func NewAmqpRequests(clt amqp.Client) *AmqpRequests {
	return &AmqpRequests{
		clt: clt,
	}
}

func (r *AmqpRequests) GetListMovies(cmd *ListMoviesCommand) ([]movie_store.Movie, error) {
	response, err := r.call("movie.list", cmd)
	if err != nil {
		return nil, err
	}
	var movies []movie_store.Movie
	err = json.Unmarshal(response.Body, &movies)
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *AmqpRequests) GetMovieByName(cmd *GetMovieByNameCommand) (*movie_store.Movie, error) {
	response, err := r.call("movie.getByName", cmd)
	if err != nil {
		return nil, err
	}
	var movie *movie_store.Movie
	err = json.Unmarshal(response.Body, &movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *AmqpRequests) GetMovieById(cmd *GetMovieByIdCommand) (*movie_store.Movie, error) {
	response, err := r.call("movie.get", cmd)
	if err != nil {
		return nil, err
	}
	var movie *movie_store.Movie
	err = json.Unmarshal(response.Body, &movie)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (r *AmqpRequests) CreateUser(cmd *CreateUserCommand) (*users_store.User, error) {
	response, err := r.call("users.create", cmd)
	if err != nil {
		return nil, err
	}

	var user *users_store.User
	err = json.Unmarshal(response.Body, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AmqpRequests) call(path string, data interface{}) (*amqp.Message, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	response, err := r.clt.Call(path, amqp.Message{
		Body: jsonData,
	})
	amqpError := &ErrorSt{}
	err = json.Unmarshal(response.Body, &amqpError)
	if err != nil {
		return nil, err
	}
	if amqpError.Text != "" {
		return nil, errors.New(amqpError.Text)
	}
	return response, nil
}
