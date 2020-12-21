package core_store

type ListMoviesCommand struct {
	Count int64 `json:"count,omitempty"`
}

func (cmd *ListMoviesCommand) Exec(service CoreService) (interface{}, error) {
	return service.ListMovies(cmd)
}

type GetMovieByNameCommand struct {
	Name string `json:"name"`
}

func (cmd *GetMovieByNameCommand) Exec(service CoreService) (interface{}, error) {
	return service.GetMovieByName(cmd)
}

type GetMovieByIdCommand struct {
	Id int64 `json:"id"`
}

func (cmd *GetMovieByIdCommand) Exec(service CoreService) (interface{}, error) {
	return service.GetMovieById(cmd)
}

type MovieRecommend struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type MovieRecommendResponse struct {
	Movies []*MovieRecommend `json:"result"`
}

type CreateUserCommand struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func (cmd *CreateUserCommand) Exec(service CoreService) (interface{}, error) {
	return service.SignUpUsingEmail(cmd)
}

type LoginUserCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cmd *LoginUserCommand) Exec(service CoreService) (interface{}, error) {
	return service.Login(cmd)
}

type GetUserByUsername struct {
	Username string `json:"username"`
}

type LoginResponse struct {
	AccessKey string `json:"access_key"`
}
